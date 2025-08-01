package uapi

import "unsafe"

const (
	// IOC_NRBITS is the number of bits allocated for the
	// command number (nr) field.
	IOC_NRBITS = 8

	// IOC_TYPEBITS is the number of bits allocated for the type field.
	IOC_TYPEBITS = 8

	// IOC_SIZEBITS is the number of bits allocated for the size field.
	IOC_SIZEBITS = 14

	// IOC_DIRBITS is the number of bits allocated for the direction
	// (read/write) field.
	IOC_DIRBITS = 2

	// IOC_NRMASK masks out the nr field bits.
	IOC_NRMASK = 1<<IOC_NRBITS - 1

	// IOC_TYPEMASK masks out the type field bits.
	IOC_TYPEMASK = 1<<IOC_TYPEBITS - 1

	// IOC_SIZEMASK masks out the size field bits.
	IOC_SIZEMASK = 1<<IOC_SIZEBITS - 1

	// IOC_DIRMASK masks out the direction field bits.
	IOC_DIRMASK = 1<<IOC_DIRBITS - 1

	// IOC_NRSHIFT is the bit offset for the nr field within the ioctl code.
	IOC_NRSHIFT = 0

	// IOC_TYPESHIFT is the bit offset for the type field within
	// the ioctl code.
	IOC_TYPESHIFT = IOC_NRSHIFT + IOC_NRBITS

	// IOC_SIZESHIFT is the bit offset for the size field within
	// the ioctl code.
	IOC_SIZESHIFT = IOC_TYPESHIFT + IOC_TYPEBITS

	// IOC_DIRSHIFT is the bit offset for the direction field within
	// the ioctl code.
	IOC_DIRSHIFT = IOC_SIZESHIFT + IOC_SIZEBITS

	// IOC_NONE specifies no data transfer for the ioctl.
	IOC_NONE = 0

	// IOC_WRITE specifies a write (user to kernel) transfer for the ioctl.
	IOC_WRITE = 1

	// IOC_READ specifies a read (kernel to user) transfer for the ioctl.
	IOC_READ = 2
)

// IOC_TYPECHECK returns the size in bytes of the provided value’s type.
// It accepts any zero value Go type and wraps [unsafe.Sizeof], converting
// the result to uint. This is useful for getting the type size when
// constructing ioctl request codes.
func IOC_TYPECHECK[T any](typ T) uint {
	return uint(unsafe.Sizeof(typ))
}

// IOC packs the four ioctl components into a single request code.
// dir specifies the data transfer direction ([IOC_NONE], [IOC_READ],
// [IOC_WRITE]).
// typ is the magic number for the driver or subsystem.
// nr is the command sequence number within that magic range.
// size is the byte size of any data transfer.
// The resulting uint can be passed directly to
// [syscall.Syscall] or [unix.IoctlGetInt] and family.
func IOC(dir, typ, nr, size uint) uint {
	return dir<<IOC_DIRSHIFT |
		typ<<IOC_TYPESHIFT |
		nr<<IOC_NRSHIFT |
		size<<IOC_SIZESHIFT
}

// IO returns an ioctl request code that carries no data.
// It encodes the given magic type and command number into a uint,
// setting direction to [IOC_NONE] and size to zero.
// typ is the magic number for the driver or subsystem.
// nr is the command sequence number within that magic range.
// The resulting uint can be passed directly to
// [syscall.Syscall] or [unix.IoctlGetInt] and family.
func IO(typ, nr uint) uint {
	return IOC(IOC_NONE, typ, nr, 0)
}

// IOR returns an ioctl request code for reading data from the kernel.
// typ is the magic number for the driver or subsystem.
// nr is the command sequence number within that magic range.
// argtype should be a zero‐value Go type.
// The resulting uint can be passed directly to
// [syscall.Syscall] or [unix.IoctlGetInt] and family.
func IOR[T any](typ, nr uint, argtype T) uint {
	return IOC(IOC_READ, typ, nr, IOC_TYPECHECK(argtype))
}

// IOW returns an ioctl request code for writing data to the kernel.
// typ is the magic number for the driver or subsystem.
// nr is the command sequence number within that magic range.
// argtype should be a zero‐value Go type.
// The resulting uint can be passed directly to
// [syscall.Syscall] or [unix.IoctlGetInt] and family.
func IOW[T any](typ, nr uint, argtype T) uint {
	return IOC(IOC_WRITE, typ, nr, IOC_TYPECHECK(argtype))
}

// IOWR returns an ioctl request code for bidirectional data transfer.
// typ is the magic number for the driver or subsystem.
// nr is the command sequence number within that magic range.
// argtype should be a zero‐value Go type.
// The resulting uint can be passed directly to
// [syscall.Syscall] or [unix.IoctlGetInt] and family.
func IOWR[T any](typ, nr uint, argtype T) uint {
	return IOC(IOC_READ|IOC_WRITE, typ, nr, IOC_TYPECHECK(argtype))
}

// IOR_BAD returns an ioctl request code for read-only operations using the
// legacy bad size encoding.
// typ is the magic number for the driver or subsystem.
// nr is the command sequence number within that magic range.
// argtype should be a zero‐value Go type.
// The resulting uint can be passed directly to
// [syscall.Syscall] or [unix.IoctlGetInt] and family.
func IOR_BAD[T any](typ, nr uint, argtype T) uint {
	return IOC(IOC_READ, typ, nr, uint(unsafe.Sizeof(argtype)))
}

// IOW_BAD returns an ioctl request code for write-only operations using the
// legacy bad size encoding.
// typ is the magic number for the driver or subsystem.
// nr is the command sequence number within that magic range.
// argtype should be a zero‐value Go type.
// The resulting uint can be passed directly to
// [syscall.Syscall] or [unix.IoctlGetInt] and family.
func IOW_BAD[T any](typ, nr uint, argtype T) uint {
	return IOC(IOC_WRITE, typ, nr, uint(unsafe.Sizeof(argtype)))
}

// IOWR_BAD returns an ioctl request code for read–write operations using the
// legacy bad size encoding.
// typ is the magic number for the driver or subsystem.
// nr is the command sequence number within that magic range.
// argtype should be a zero‐value Go type.
// The resulting uint can be passed directly to
// [syscall.Syscall] or [unix.IoctlGetInt] and family.
func IOWR_BAD[T any](typ, nr uint, argtype T) uint {
	return IOC(IOC_READ|IOC_WRITE, typ, nr, uint(unsafe.Sizeof(argtype)))
}

// IOC_DIR extracts the direction bits from an ioctl request code.
// It returns one of [IOC_NONE], [IOC_WRITE], [IOC_READ], or their
// combination.
func IOC_DIR(nr uint) uint {
	return nr >> IOC_DIRSHIFT & IOC_DIRMASK
}

// IOC_TYPE extracts the magic/type field from an ioctl request code.
// The magic value identifies the driver or subsystem namespace.
func IOC_TYPE(nr uint) uint {
	return nr >> IOC_TYPESHIFT & IOC_TYPEMASK
}

// IOC_NR extracts the command number field from an ioctl request code.
// This is the sequential identifier within the magic/type namespace.
func IOC_NR(nr uint) uint {
	return nr >> IOC_NRSHIFT & IOC_NRMASK
}

// IOC_SIZE extracts the size field (in bytes) from an ioctl request code.
// This indicates how many bytes of data are copied to or from user space.
func IOC_SIZE(nr uint) uint {
	return nr >> IOC_SIZESHIFT & IOC_SIZEMASK
}

// IOC_IN returns the encoded flag for a write (user -> kernel) transfer.
// It shifts [IOC_WRITE] into the direction bits of an ioctl code.
func IOC_IN() uint {
	return IOC_WRITE << IOC_DIRSHIFT
}

// IOC_OUT returns the encoded flag for a read (kernel -> user) transfer.
// It shifts [IOC_READ] into the direction bits of an ioctl code.
func IOC_OUT() uint {
	return IOC_READ << IOC_DIRSHIFT
}

// IOC_INOUT returns the encoded flag for a bidirectional transfer.
// It combines [IOC_WRITE] and [IOC_READ] in the direction bits.
func IOC_INOUT() uint {
	return IOC_WRITE<<IOC_DIRSHIFT | IOC_READ<<IOC_DIRSHIFT
}

// IOCSIZE_MASK returns the full bitmask for the size field
// positioned at its shift offset within an ioctl request code.
func IOCSIZE_MASK() uint {
	return IOC_SIZEMASK << IOC_SIZESHIFT
}

// IOSIZE_SHIFT returns the bit offset of the size field
// within an ioctl request code.
func IOSIZE_SHIFT() uint {
	return IOC_SIZESHIFT
}
