//go:build linux

// Package ioctl_test_cgo provides C-backed ioctl constants extracted via cgo
// from <linux/ioctl.h>. It exists solely for internal test validation.
//
// These constants are used to verify that Go-side definitions match the
// kernel's encoding layout. This package is gated to Linux platforms and
// should not be imported outside of test code.
//
// All values are declared as Go constants for compatibility with static
// analysis and test assertions.
package ioctl_test_cgo

/*
#include <linux/ioctl.h>

unsigned int wrap_IOC(int dir, int type, int nr, int size) {
    return _IOC(dir, type, nr, size);
}
*/
import "C"

// These constants mirror the values defined in <linux/ioctl.h> and are
// extracted via cgo for internal test validation. They represent the
// bitfield layout and direction flags used in ioctl request encoding.
//
// This block exists solely to verify that Go-side definitions match the
// kernel's ABI. All values are declared as constants to support static
// analysis and deterministic test comparisons. Do not use these outside
// of test code.
const (
	IOC_NRBITS    = uint32(C._IOC_NRBITS)
	IOC_TYPEBITS  = uint32(C._IOC_TYPEBITS)
	IOC_NRMASK    = uint32(C._IOC_NRMASK)
	IOC_TYPEMASK  = uint32(C._IOC_TYPEMASK)
	IOC_NRSHIFT   = uint32(C._IOC_NRSHIFT)
	IOC_TYPESHIFT = uint32(C._IOC_TYPESHIFT)
	IOC_SIZESHIFT = uint32(C._IOC_SIZESHIFT)
	IOC_SIZEBITS  = uint32(C._IOC_SIZEBITS)
	IOC_DIRBITS   = uint32(C._IOC_DIRBITS)
	IOC_NONE      = uint32(C._IOC_NONE)
	IOC_WRITE     = uint32(C._IOC_WRITE)
	IOC_READ      = uint32(C._IOC_READ)
	IOC_SIZEMASK  = uint32(C._IOC_SIZEMASK)
	IOC_DIRMASK   = uint32(C._IOC_DIRMASK)
	IOC_DIRSHIFT  = uint32(C._IOC_DIRSHIFT)
	IOC_IN        = uint32(C.IOC_IN)
	IOC_OUT       = uint32(C.IOC_OUT)
	IOC_INOUT     = uint32(C.IOC_INOUT)
	IOCSIZE_MASK  = uint32(C.IOCSIZE_MASK)
	IOCSIZE_SHIFT = uint32(C.IOCSIZE_SHIFT)
)

// IOC encodes an ioctl command using four bitfield components.
//
// This wraps the C _IOC macro via cgo, combining dir, typ, nr, and size
// into a single uint32 command. The encoding layout is platform-dependent
// and may vary across kernel versions. This wrapper defers to the C macro
// for correctness.
//
// Use IOC to construct ioctl constants for syscall or device interaction.
// All inputs and outputs are uint32 to preserve bit layout and avoid
// platform-dependent types.
func IOC(dir, typ, nr, size uint32) uint32 {
	return uint32(C.wrap_IOC(
		C.int(dir),
		C.int(typ),
		C.int(nr),
		C.int(size),
	))
}
