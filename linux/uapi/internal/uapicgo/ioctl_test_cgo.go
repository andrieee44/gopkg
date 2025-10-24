package uapicgo

// #include <linux/ioctl.h>
// #include "ioctl.h"
import "C"
import "github.com/andrieee44/gopkg/linux/uapi/ioctl"

func IoctlSetup() {
	ioctl.SetIOC_SIZEBITS(C._IOC_SIZEBITS)
	ioctl.SetIOC_DIRBITS(C._IOC_DIRBITS)
	ioctl.SetIOC_NONE(C._IOC_NONE)
	ioctl.SetIOC_WRITE(C._IOC_WRITE)
	ioctl.SetIOC_READ(C._IOC_READ)
}

func IOC(dir, typ, nr, size uint32) uint32 {
	return uint32(C.wrap_IOC(
		C.int(dir),
		C.int(typ),
		C.int(nr),
		C.int(size),
	))
}

func IoctlConstants() []Constant[uint32] {
	return []Constant[uint32]{
		{"ioctl.IOC_NRBITS", ioctl.IOC_NRBITS, C._IOC_NRBITS},
		{"ioctl.IOC_TYPEBITS", ioctl.IOC_TYPEBITS, C._IOC_TYPEBITS},
		{"ioctl.IOC_NRMASK", ioctl.IOC_NRMASK, C._IOC_NRMASK},
		{"ioctl.IOC_TYPEMASK", ioctl.IOC_TYPEMASK, C._IOC_TYPEMASK},
		{"ioctl.IOC_NRSHIFT", ioctl.IOC_NRSHIFT, C._IOC_NRSHIFT},
		{"ioctl.IOC_TYPESHIFT", ioctl.IOC_TYPESHIFT, C._IOC_TYPESHIFT},
		{"ioctl.IOC_SIZESHIFT", ioctl.IOC_SIZESHIFT, C._IOC_SIZESHIFT},
		{"ioctl.IOC_SIZEBITS", ioctl.IOC_SIZEBITS, C._IOC_SIZEBITS},
		{"ioctl.IOC_DIRBITS", ioctl.IOC_DIRBITS, C._IOC_DIRBITS},
		{"ioctl.IOC_NONE", ioctl.IOC_NONE, C._IOC_NONE},
		{"ioctl.IOC_WRITE", ioctl.IOC_WRITE, C._IOC_WRITE},
		{"ioctl.IOC_READ", ioctl.IOC_READ, C._IOC_READ},
		{"ioctl.IOC_SIZEMASK()", ioctl.IOC_SIZEMASK(), C._IOC_SIZEMASK},
		{"ioctl.IOC_DIRMASK()", ioctl.IOC_DIRMASK(), C._IOC_DIRMASK},
		{"ioctl.IOC_DIRSHIFT()", ioctl.IOC_DIRSHIFT(), C._IOC_DIRSHIFT},
		{"ioctl.IOC_IN()", ioctl.IOC_IN(), C.IOC_IN},
		{"ioctl.IOC_OUT()", ioctl.IOC_OUT(), C.IOC_OUT},
		{"ioctl.IOC_INOUT()", ioctl.IOC_INOUT(), C.IOC_INOUT},
		{"ioctl.IOCSIZE_MASK()", ioctl.IOCSIZE_MASK(), C.IOCSIZE_MASK},
		{"ioctl.IOCSIZE_SHIFT", ioctl.IOCSIZE_SHIFT, C.IOCSIZE_SHIFT},
	}
}
