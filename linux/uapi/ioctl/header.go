package ioctl

import (
	"errors"
	"fmt"
	"math"
	"unsafe"

	"github.com/andrieee44/gopkg/lib/bitops"
)

const (
	// IOC_NRBITS is the number of bits reserved for the request number
	// field.
	IOC_NRBITS = 8

	// IOC_TYPEBITS is the number of bits reserved for the request type
	// (subsystem) field.
	IOC_TYPEBITS = 8

	// IOC_NRMASK is a mask with IOC_NRBITS low bits set.
	// Use it to bound or extract the request number field.
	IOC_NRMASK = (1 << IOC_NRBITS) - 1

	// IOC_TYPEMASK is a mask with IOC_TYPEBITS low bits set.
	// Use it to bound or extract the request type field.
	IOC_TYPEMASK = (1 << IOC_TYPEBITS) - 1

	// IOC_NRSHIFT is the bit offset of the request number field.
	IOC_NRSHIFT = 0

	// IOC_TYPESHIFT is the bit offset of the request type field.
	IOC_TYPESHIFT = IOC_NRSHIFT + IOC_NRBITS

	// IOC_SIZESHIFT is the bit offset of the size field.
	IOC_SIZESHIFT = IOC_TYPESHIFT + IOC_TYPEBITS
)

var (
	// ErrSizeOverflow indicates that a size value exceeds the maximum
	// representable limit for a 32‑bit size field. This typically occurs
	// when attempting to encode a request with a payload larger than
	// the allowed maximum.
	ErrSizeOverflow error = errors.New("size exceeds 32-bit limit")

	// ErrBitOverflow indicates that a bit count exceeds the capacity
	// allocated for the field. This can happen if a shift or mask
	// operation attempts to use more bits than defined by the protocol.
	ErrBitOverflow error = errors.New("bit count exceeds allocated capacity")

	// ErrRequestTooBig indicates that the composed ioctl request code
	// would not fit within a 32‑bit unsigned integer.
	ErrRequestTooBig error = errors.New("request code exceeds 32-bit limit")

	// IOC_SIZEBITS is the number of bits reserved for the “size” field
	// in an ioctl command number. It determines the maximum encodable
	// payload size. Declared as a variable so it can be modified at
	// runtime — use SetIOC_SIZEBITS to adjust for platform or header
	// differences.
	IOC_SIZEBITS uint32 = 14

	// IOC_DIRBITS is the number of bits reserved for the “direction” field
	// (e.g., read, write, none) in an ioctl command number. Declared as
	// a variable so its value can be changed at runtime — use SetIOC_DIRBITS
	// for compatibility with system‑specific constants.
	IOC_DIRBITS uint32 = 2

	// IOC_NONE holds the value representing “no data transfer” in the
	// ioctl direction field. Declared as a variable so it can be modified
	// via SetIOC_NONE to match platform or header definitions.
	IOC_NONE uint32 = 0

	// IOC_WRITE holds the value representing a “write” direction in the
	// ioctl direction field (data moves from user space to kernel space).
	// Declared as a variable so its value can be changed using SetIOC_WRITE
	// for compatibility with system‑specific constants.
	IOC_WRITE uint32 = 1

	// IOC_READ holds the value representing a “read” direction in the
	// ioctl direction field (data moves from kernel space to user space).
	// Declared as a variable so its value can be updated dynamically via
	// SetIOC_READ to align with platform or header settings.
	IOC_READ uint32 = 2
)

// SetIOC_SIZEBITS updates the number of bits reserved for the “size” field
// in an ioctl command number. This allows runtime adjustment to match
// platform‑specific or header‑defined values.
func SetIOC_SIZEBITS(value uint32) {
	IOC_SIZEBITS = value
}

// SetIOC_DIRBITS updates the number of bits reserved for the “direction”
// field (e.g., read, write, none) in an ioctl command number. Use this to
// align the setting with platform‑specific or header‑defined values.
func SetIOC_DIRBITS(value uint32) {
	IOC_DIRBITS = value
}

// SetIOC_NONE changes the value representing “no data transfer” in the
// ioctl direction field. This is useful when the constant differs across
// platforms or needs to be overridden by header definitions.
func SetIOC_NONE(value uint32) {
	IOC_NONE = value
}

// SetIOC_WRITE changes the value representing a “write” direction in the
// ioctl direction field (data moves from user space to kernel space).
// Adjust this if your platform or headers define a different value.
func SetIOC_WRITE(value uint32) {
	IOC_WRITE = value
}

// SetIOC_READ changes the value representing a “read” direction in the
// ioctl direction field (data moves from kernel space to user space).
// Adjust this if your platform or headers define a different value.
func SetIOC_READ(value uint32) {
	IOC_READ = value
}

// IOC_SIZEMASK returns a bitmask with [IOC_SIZEBITS] low bits set.
// This mask is used to isolate or bound the “size” field in an
// ioctl command number. The value depends on the current
// [IOC_SIZEBITS] setting, which can be changed at runtime via
// [SetIOC_SIZEBITS].
func IOC_SIZEMASK() uint32 {
	return 1<<IOC_SIZEBITS - 1
}

// IOC_DIRMASK returns a bitmask with IOC_DIRBITS low bits set.
// This mask is used to isolate or bound the “direction” field
// in an ioctl command number. The value depends on the current
// IOC_DIRBITS setting, which can be changed at runtime via
// SetIOC_DIRBITS.
func IOC_DIRMASK() uint32 {
	return 1<<IOC_DIRBITS - 1
}

// IOC_DIRSHIFT returns the bit offset of the “direction” field
// in an ioctl command number. It is computed as [IOC_SIZESHIFT]
// plus the current [IOC_SIZEBITS] value, allowing the offset to
// adapt if [IOC_SIZEBITS] is changed at runtime via [SetIOC_SIZEBITS].
func IOC_DIRSHIFT() uint32 {
	return IOC_SIZESHIFT + IOC_SIZEBITS
}

// IOC_TYPECHECK reports the size in bytes of type T as a uint32,
// returning an error if the size cannot be represented within
// uint32 or exceeds the [IOC_SIZEBITS] field limit.
func IOC_TYPECHECK[T any]() (uint32, error) {
	var (
		dataType T
		size     uintptr
	)

	size = unsafe.Sizeof(dataType)
	if size > math.MaxUint32 {
		return 0, fmt.Errorf(
			"%T is %d bits: %w",
			dataType, size*8, ErrSizeOverflow,
		)
	}

	if bitops.OverflowsUnsigned(IOC_SIZEBITS, uint32(size)) {
		return 0, fmt.Errorf(
			"%T is %d bits, max bits is %d: %w",
			dataType, size*8, IOC_SIZEBITS, ErrBitOverflow,
		)
	}

	return uint32(size), nil
}

// IOC encodes the provided dir, typ, nr, and size values into a single
// 32‑bit ioctl request code. It first validates that the combined bit
// widths defined by [IOC_DIRBITS], [IOC_TYPEBITS], [IOC_NRBITS], and
// [IOC_SIZEBITS] do not exceed 32, and that each value fits within its
// allocated field. If the total exceeds 32 bits, it returns
// [ErrRequestTooBig]. If a value is too large for its field, it returns
// [ErrBitOverflow] with details.
//
// When validation passes, the request code is assembled by shifting
// each field into its position using [IOC_DIRSHIFT](), [IOC_TYPESHIFT],
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
		{"dir", dir, IOC_DIRBITS},
		{"typ", typ, IOC_TYPEBITS},
		{"nr", nr, IOC_NRBITS},
		{"size", size, IOC_SIZEBITS},
	}

	requestSize = IOC_DIRBITS + IOC_TYPEBITS + IOC_NRBITS + IOC_SIZEBITS
	if requestSize > 32 {
		return 0, fmt.Errorf(
			"request size is %d bits: %w",
			requestSize, ErrRequestTooBig,
		)
	}

	for _, check = range checks {
		if !bitops.OverflowsUnsigned(check.bitLimit, check.value) {
			continue
		}

		return 0, fmt.Errorf(
			"%s is %d, max value is %d: %w",
			check.name, check.value, (1<<check.bitLimit)-1, ErrBitOverflow,
		)
	}

	return dir<<IOC_DIRSHIFT() |
		typ<<IOC_TYPESHIFT |
		nr<<IOC_NRSHIFT |
		size<<IOC_SIZESHIFT, nil
}

// IO returns an ioctl request code for commands with no associated
// data transfer, using [IOC_NONE] for the direction, the provided typ
// and nr values, and a size of zero. It calls [IOC] to perform the
// encoding and validation.
func IO(typ, nr uint32) (uint32, error) {
	return IOC(IOC_NONE, typ, nr, 0)
}

// IOR returns an ioctl request code for commands that read data from
// the kernel into user space. It calls the internal [ioc] helper with
// [IOC_READ] as the direction and uses the generic type parameter T to
// determine the size of the data structure.
func IOR[T any](typ, nr uint32) (uint32, error) {
	return ioc[T](IOC_READ, typ, nr)
}

// IOW returns an ioctl request code for commands that write data from
// user space into the kernel. It calls the internal [ioc] helper with
// [IOC_WRITE] as the direction and uses the generic type parameter T to
// determine the size of the data structure.
func IOW[T any](typ, nr uint32) (uint32, error) {
	return ioc[T](IOC_WRITE, typ, nr)
}

// IOWR returns an ioctl request code for commands that both read from
// and write to the kernel. It calls the internal [ioc] helper with a
// direction of [IOC_READ] | [IOC_WRITE] and uses the generic type
// parameter T to determine the size of the data structure.
func IOWR[T any](typ, nr uint32) (uint32, error) {
	return ioc[T](IOC_READ|IOC_WRITE, typ, nr)
}

// IOC_DIR extracts the direction field from the given ioctl request
// number nr by shifting it right by [IOC_DIRSHIFT]() and masking with
// [IOC_DIRMASK]().
func IOC_DIR(nr uint32) uint32 {
	return nr >> IOC_DIRSHIFT() & IOC_DIRMASK()
}

// IOC_TYPE extracts the type field from the given error ioctl request number
// nr by shifting it right by [IOC_TYPESHIFT] and masking with
// [IOC_TYPEMASK].
func IOC_TYPE(nr uint32) uint32 {
	return nr >> IOC_TYPESHIFT & IOC_TYPEMASK
}

// IOC_NR extracts the command number field from the given ioctl request
// number nr by shifting it right by [IOC_NRSHIFT] and masking with
// [IOC_NRMASK].
func IOC_NR(nr uint32) uint32 {
	return nr >> IOC_NRSHIFT & IOC_NRMASK
}

// IOC_SIZE extracts the size field from the given ioctl request number
// nr by shifting it right by [IOC_SIZESHIFT] and masking with
// [IOC_SIZEMASK]().
func IOC_SIZE(nr uint32) uint32 {
	return nr >> IOC_SIZESHIFT & IOC_SIZEMASK()
}

// IOC_IN returns a request code mask indicating that data will be
// written from user space into the kernel. This is computed as
// [IOC_WRITE] shifted left by [IOC_DIRSHIFT]().
func IOC_IN() uint32 {
	return IOC_WRITE << IOC_DIRSHIFT()
}

// IOC_OUT returns a request code mask indicating that data will be read
// from the kernel into user space. This is computed as [IOC_READ]
// shifted left by [IOC_DIRSHIFT]().
func IOC_OUT() uint32 {
	return IOC_READ << IOC_DIRSHIFT()
}

// IOC_INOUT returns a request code mask indicating that data will be
// both read from the kernel and written to it. This is computed as the
// bitwise OR of [IOC_WRITE] << [IOC_DIRSHIFT]() and [IOC_READ] <<
// [IOC_DIRSHIFT]().
func IOC_INOUT() uint32 {
	return IOC_WRITE<<IOC_DIRSHIFT() | IOC_READ<<IOC_DIRSHIFT()
}

// IOCSIZE_MASK returns a mask covering the size field of an ioctl
// request number. The mask is computed by shifting [IOC_SIZEMASK]() to
// the left by [IOC_SIZESHIFT].
func IOCSIZE_MASK() uint32 {
	return IOC_SIZEMASK() << IOC_SIZESHIFT
}

// IOSIZE_SHIFT returns the bit offset of the size field within an ioctl
// request number, which is equal to [IOC_SIZESHIFT].
func IOSIZE_SHIFT() uint32 {
	return IOC_SIZESHIFT
}
