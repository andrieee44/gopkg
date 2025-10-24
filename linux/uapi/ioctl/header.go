//go:build linux

package ioctl

import (
	"fmt"
	"math"
	"unsafe"

	"github.com/andrieee44/gopkg/lib/bitops"
	"github.com/andrieee44/gopkg/lib/xerr"
)

const (
	// IOC_NRBITS defines the number of bits reserved for the ioctl
	// request number field.
	//
	// This field identifies the specific command within a given ioctl
	// type.
	IOC_NRBITS = 8

	// IOC_TYPEBITS defines the number of bits reserved for the ioctl
	// request type field.
	//
	// This field identifies the device class or subsystem.
	IOC_TYPEBITS = 8

	// IOC_NRMASK is a bitmask with IOC_NRBITS low bits set.
	//
	// Use it to extract or bound the request number field from an
	// ioctl code.
	IOC_NRMASK = (1 << IOC_NRBITS) - 1

	// IOC_TYPEMASK is a bitmask with IOC_TYPEBITS low bits set.
	//
	// Use it to extract or bound the request type field from an ioctl
	// code.
	IOC_TYPEMASK = (1 << IOC_TYPEBITS) - 1

	// IOC_NRSHIFT is the bit offset of the request number field.
	//
	// It is typically zero, placing the number in the least
	// significant bits.
	IOC_NRSHIFT = 0

	// IOC_TYPESHIFT is the bit offset of the request type field.
	//
	// It follows the number field and precedes the size field.
	IOC_TYPESHIFT = IOC_NRSHIFT + IOC_NRBITS

	// IOC_SIZESHIFT is the bit offset of the size field.
	//
	// It follows the type field and precedes the direction field.
	IOC_SIZESHIFT = IOC_TYPESHIFT + IOC_TYPEBITS

	// IOCSIZE_SHIFT is an alias for IOC_SIZESHIFT.
	//
	// It is used in some ioctl macros and platform headers.
	IOCSIZE_SHIFT = IOC_SIZESHIFT
)

var (
	// IOC_SIZEBITS is the number of bits reserved for the "size" field
	// in an ioctl command number.
	//
	// It determines the maximum encodable payload size. Declared as a
	// variable so its value can be updated dynamically via [SetIOC_SIZEBITS]
	// to align with platform or header settings.
	IOC_SIZEBITS uint32 = 14

	// IOC_DIRBITS is the number of bits reserved for the "direction" field
	// (e.g., read, write, none) in an ioctl command number.
	//
	// Declared as a variable so its value can be updated dynamically via
	// [SetIOC_DIRBITS] to align with platform or header settings.
	IOC_DIRBITS uint32 = 2

	// IOC_NONE holds the value representing "no data transfer" in the
	// ioctl direction field.
	//
	// Declared as a variable so its value can be updated dynamically via
	// [SetIOC_NONE] to align with platform or header settings.
	IOC_NONE uint32 = 0

	// IOC_WRITE holds the value representing a "write" direction in the
	// ioctl direction field (data moves from user space to kernel space).
	//
	// Declared as a variable so its value can be updated dynamically via
	// [SetIOC_WRITE] to align with platform or header settings.
	IOC_WRITE uint32 = 1

	// IOC_READ holds the value representing a "read" direction in the
	// ioctl direction field (data moves from kernel space to user space).
	//
	// Declared as a variable so its value can be updated dynamically via
	// [SetIOC_READ] to align with platform or header settings.
	IOC_READ uint32 = 2
)

// SetIOC_SIZEBITS updates the number of bits reserved for the "size" field
// in an ioctl command number.
//
// Adjust this if your platform or headers define a different value.
func SetIOC_SIZEBITS(value uint32) {
	IOC_SIZEBITS = value
}

// SetIOC_DIRBITS updates the number of bits reserved for the "direction"
// field (e.g., read, write, none) in an ioctl command number.
//
// Adjust this if your platform or headers define a different value.
func SetIOC_DIRBITS(value uint32) {
	IOC_DIRBITS = value
}

// SetIOC_NONE changes the value representing "no data transfer" in the
// ioctl direction field.
//
// Adjust this if your platform or headers define a different value.
func SetIOC_NONE(value uint32) {
	IOC_NONE = value
}

// SetIOC_WRITE changes the value representing a "write" direction in the
// ioctl direction field (data moves from user space to kernel space).
//
// Adjust this if your platform or headers define a different value.
func SetIOC_WRITE(value uint32) {
	IOC_WRITE = value
}

// SetIOC_READ changes the value representing a "read" direction in the
// ioctl direction field (data moves from kernel space to user space).
//
// Adjust this if your platform or headers define a different value.
func SetIOC_READ(value uint32) {
	IOC_READ = value
}

// IOC_SIZEMASK returns a bitmask with [IOC_SIZEBITS] low bits set.
//
// This mask is used to isolate or bound the "size" field in an ioctl command
// number. The value depends on the current [IOC_SIZEBITS] setting, which can
// be changed at runtime via [SetIOC_SIZEBITS].
func IOC_SIZEMASK() uint32 {
	return 1<<IOC_SIZEBITS - 1
}

// IOC_DIRMASK returns a bitmask with [IOC_DIRBITS] low bits set.
//
// This mask is used to isolate or bound the "direction" field in an ioctl
// command number. The value depends on the current [IOC_DIRBITS] setting,
// which can be changed at runtime via [SetIOC_DIRBITS].
func IOC_DIRMASK() uint32 {
	return 1<<IOC_DIRBITS - 1
}

// IOC_DIRSHIFT returns the bit offset of the "direction" field in an ioctl
// command number.
//
// It is computed as [IOC_SIZESHIFT] plus the current [IOC_SIZEBITS] value,
// allowing the offset to adapt if [IOC_SIZEBITS] is changed at runtime via
// [SetIOC_SIZEBITS].
func IOC_DIRSHIFT() uint32 {
	return IOC_SIZESHIFT + IOC_SIZEBITS
}

// IOC_TYPECHECK reports the size in bytes of type T as a uint32.
//
// It returns [ErrSizeOverflow] if the size exceeds [math.MaxUint32],
// or [ErrBitOverflow] if the size cannot be encoded within
// [IOC_SIZEBITS] bits.
func IOC_TYPECHECK[T any]() (uint32, error) {
	var (
		dataType T
		size     uintptr
	)

	size = unsafe.Sizeof(dataType)
	if size > math.MaxUint32 {
		return 0, fmt.Errorf(
			"ioctl.IOC_TYPECHECK: %T is %d bytes, exceeds 32-bit limit",
			dataType, size,
		)
	}

	if bitops.OverflowsUnsigned(IOC_SIZEBITS, uint32(size)) {
		return 0, fmt.Errorf(
			"ioctl.IOC_TYPECHECK: %T is %d bytes, exceeds IOC_SIZEBITS (%d bits)",
			dataType, size, IOC_SIZEBITS,
		)
	}

	return uint32(size), nil
}

// IOC encodes the provided dir, typ, nr, and size values into a single
// 32-bit ioctl request code.
//
// It first validates that the combined bit widths defined by
// [IOC_DIRBITS], [IOC_TYPEBITS], [IOC_NRBITS], and [IOC_SIZEBITS] do not
// exceed 32, and that each value fits within its allocated field. If the
// total exceeds 32 bits, it returns [ErrRequestTooBig]. If a value is too
// large for its field, it returns [ErrBitOverflow] with details.
//
// When validation passes, the request code is assembled by shifting each
// field into its position using [IOC_DIRSHIFT], [IOC_TYPESHIFT],
// [IOC_NRSHIFT], and [IOC_SIZESHIFT]. The bit width settings may be
// changed at runtime via their respective setter functions such as
// [SetIOC_SIZEBITS] and [SetIOC_DIRBITS], which will affect both
// validation and encoding.
func IOC(dir, typ, nr, size uint32) (uint32, error) {
	type argCheck struct {
		name     string
		value    uint32
		bitLimit uint32
	}

	var (
		requestSize uint32
		checks      []argCheck
		check       argCheck
	)

	checks = []argCheck{
		{"IOC_DIRBITS", dir, IOC_DIRBITS},
		{"IOC_TYPEBITS", typ, IOC_TYPEBITS},
		{"IOC_NRBITS", nr, IOC_NRBITS},
		{"IOC_SIZEBITS", size, IOC_SIZEBITS},
	}

	requestSize = IOC_DIRBITS + IOC_TYPEBITS + IOC_NRBITS + IOC_SIZEBITS
	if requestSize > 32 {
		return 0, fmt.Errorf(
			"ioctl.IOC: request size is %d bits "+
				"(IOC_DIRBITS=%d + IOC_TYPEBITS=%d + "+
				"IOC_NRBITS=%d + IOC_SIZEBITS=%d), "+
				"exceeds 32-bit limit",
			requestSize,
			IOC_DIRBITS,
			IOC_TYPEBITS,
			IOC_NRBITS,
			IOC_SIZEBITS,
		)
	}

	for _, check = range checks {
		if !bitops.OverflowsUnsigned(check.bitLimit, check.value) {
			continue
		}

		return 0, fmt.Errorf(
			"ioctl.IOC: %s=%d (limit: %d for %d bits)",
			check.name,
			check.value,
			(1<<check.bitLimit)-1,
			check.bitLimit,
		)
	}

	return dir<<IOC_DIRSHIFT() |
		typ<<IOC_TYPESHIFT |
		nr<<IOC_NRSHIFT |
		size<<IOC_SIZESHIFT, nil
}

// IO returns an ioctl request code for a command with no data transfer.
//
// It uses [IOC_NONE] as the direction and sets the size to zero.
func IO(typ, nr uint32) (uint32, error) {
	return xerr.WrapIf1("ioctl.IO", func() (uint32, error) {
		return IOC(IOC_NONE, typ, nr, 0)
	})
}

// IOR returns an ioctl request code for a command that reads data from
// the kernel.
//
// The size is derived from the generic type parameter T.
func IOR[T any](typ, nr uint32) (uint32, error) {
	return xerr.WrapIf1("ioctl.IOR", func() (uint32, error) {
		return ioc[T](IOC_READ, typ, nr)
	})
}

// IOW returns an ioctl request code for a command that writes data to the
// kernel.
//
// The size is derived from the generic type parameter T.
func IOW[T any](typ, nr uint32) (uint32, error) {
	return xerr.WrapIf1("ioctl.IOW", func() (uint32, error) {
		return ioc[T](IOC_WRITE, typ, nr)
	})
}

// IOWR returns an ioctl request code for a command that both reads and
// writes data.
//
// The size is derived from the generic type parameter T.
func IOWR[T any](typ, nr uint32) (uint32, error) {
	return xerr.WrapIf1("ioctl.IOWR", func() (uint32, error) {
		return ioc[T](IOC_READ|IOC_WRITE, typ, nr)
	})
}

// IOC_DIR extracts the direction field from an ioctl request code.
//
// The caller must be aware that the direction field offset may vary across
// platforms. This function reflects the current runtime state of
// [IOC_SIZEBITS] and uses [IOC_DIRSHIFT] and [IOC_DIRMASK] to extract the
// field.
func IOC_DIR(nr uint32) uint32 {
	return nr >> IOC_DIRSHIFT() & IOC_DIRMASK()
}

// IOC_TYPE extracts the type field from an ioctl request code.
//
// This field identifies the device class or subsystem. The extraction uses
// [IOC_TYPESHIFT] and [IOC_TYPEMASK], which are fixed-width constants.
func IOC_TYPE(nr uint32) uint32 {
	return nr >> IOC_TYPESHIFT & IOC_TYPEMASK
}

// IOC_NR extracts the command number field from an ioctl request code.
//
// This field identifies the specific command within a type. It is extracted
// using [IOC_NRSHIFT] and [IOC_NRMASK], which are fixed-width constants.
func IOC_NR(nr uint32) uint32 {
	return nr >> IOC_NRSHIFT & IOC_NRMASK
}

// IOC_SIZE extracts the size field from an ioctl request code.
//
// The caller must be aware that the size field width may vary across
// platforms. This function reflects the current runtime state of
// [IOC_SIZEBITS] via [IOC_SIZEMASK].
func IOC_SIZE(nr uint32) uint32 {
	return nr >> IOC_SIZESHIFT & IOC_SIZEMASK()
}

// IOC_IN returns a mask indicating a write from user space to the kernel.
//
// The returned value reflects the current runtime state of [IOC_DIRSHIFT].
// Unlike C macros, this is computed dynamically to ensure correctness.
func IOC_IN() uint32 {
	return IOC_WRITE << IOC_DIRSHIFT()
}

// IOC_OUT returns a mask indicating a read from the kernel to user space.
//
// The returned value reflects the current runtime state of [IOC_DIRSHIFT].
// Unlike C macros, this is computed dynamically to ensure correctness.
func IOC_OUT() uint32 {
	return IOC_READ << IOC_DIRSHIFT()
}

// IOC_INOUT returns a mask indicating both read and write directions.
//
// The returned value reflects the current runtime state of [IOC_DIRSHIFT].
// Unlike C macros, this is computed dynamically to ensure correctness.
func IOC_INOUT() uint32 {
	return IOC_WRITE<<IOC_DIRSHIFT() | IOC_READ<<IOC_DIRSHIFT()
}

// IOCSIZE_MASK returns a mask covering the size field in an ioctl request
// code.
//
// The caller must be aware that the size field width may vary across
// platforms. This function reflects the current runtime state of})
// [IOC_SIZEBITS] via [IOC_SIZEMASK].
func IOCSIZE_MASK() uint32 {
	return IOC_SIZEMASK() << IOC_SIZESHIFT
}
