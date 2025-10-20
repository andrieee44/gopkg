//go:build linux

package ioctl_test

import (
	"testing"

	"github.com/andrieee44/gopkg/linux/uapi/ioctl"
	ioctlcgo "github.com/andrieee44/gopkg/linux/uapi/ioctl/internal/ioctl_test_cgo"
)

type iocTable struct {
	dir, typ, nr, size uint32
}

func init() {
	ioctl.SetIOC_SIZEBITS(ioctlcgo.IOC_SIZEBITS)
	ioctl.SetIOC_DIRBITS(ioctlcgo.IOC_DIRBITS)
	ioctl.SetIOC_NONE(ioctlcgo.IOC_NONE)
	ioctl.SetIOC_WRITE(ioctlcgo.IOC_WRITE)
	ioctl.SetIOC_READ(ioctlcgo.IOC_READ)
}

func TestVariables(t *testing.T) {
	type table struct {
		name        string
		goVal, cVal uint32
	}

	var (
		tests []table
		test  table
	)

	t.Parallel()

	tests = []table{
		{"IOC_NRBITS", ioctl.IOC_NRBITS, ioctlcgo.IOC_NRBITS},
		{"IOC_TYPEBITS", ioctl.IOC_TYPEBITS, ioctlcgo.IOC_TYPEBITS},
		{"IOC_NRMASK", ioctl.IOC_NRMASK, ioctlcgo.IOC_NRMASK},
		{"IOC_TYPEMASK", ioctl.IOC_TYPEMASK, ioctlcgo.IOC_TYPEMASK},
		{"IOC_NRSHIFT", ioctl.IOC_NRSHIFT, ioctlcgo.IOC_NRSHIFT},
		{"IOC_TYPESHIFT", ioctl.IOC_TYPESHIFT, ioctlcgo.IOC_TYPESHIFT},
		{"IOC_SIZESHIFT", ioctl.IOC_SIZESHIFT, ioctlcgo.IOC_SIZESHIFT},
		{"IOC_SIZEBITS", ioctl.IOC_SIZEBITS, ioctlcgo.IOC_SIZEBITS},
		{"IOC_DIRBITS", ioctl.IOC_DIRBITS, ioctlcgo.IOC_DIRBITS},
		{"IOC_NONE", ioctl.IOC_NONE, ioctlcgo.IOC_NONE},
		{"IOC_WRITE", ioctl.IOC_WRITE, ioctlcgo.IOC_WRITE},
		{"IOC_READ", ioctl.IOC_READ, ioctlcgo.IOC_READ},
		{"IOC_SIZEMASK", ioctl.IOC_SIZEMASK(), ioctlcgo.IOC_SIZEMASK},
		{"IOC_DIRMASK", ioctl.IOC_DIRMASK(), ioctlcgo.IOC_DIRMASK},
		{"IOC_DIRSHIFT", ioctl.IOC_DIRSHIFT(), ioctlcgo.IOC_DIRSHIFT},
		{"IOC_IN", ioctl.IOC_IN(), ioctlcgo.IOC_IN},
		{"IOC_OUT", ioctl.IOC_OUT(), ioctlcgo.IOC_OUT},
		{"IOC_INOUT", ioctl.IOC_INOUT(), ioctlcgo.IOC_INOUT},
		{"IOCSIZE_MASK", ioctl.IOCSIZE_MASK(), ioctlcgo.IOCSIZE_MASK},
		{"IOCSIZE_SHIFT", ioctl.IOCSIZE_SHIFT(), ioctlcgo.IOCSIZE_SHIFT},
	}

	for _, test = range tests {
		if test.cVal == test.goVal {
			continue
		}

		t.Errorf("got: %d, exp: %d: %s", test.goVal, test.cVal, test.name)
	}
}

func TestIOC(t *testing.T) {
	var (
		tests       []iocTable
		test        iocTable
		goReq, cReq uint32
		err         error
	)

	t.Parallel()

	tests = []iocTable{
		{dir: 0, typ: 0, nr: 0, size: 0},
		{dir: 1, typ: 0x12, nr: 0x34, size: 128},
		{dir: 2, typ: 0xAB, nr: 0xCD, size: 4096},
		{dir: 3, typ: 0xFF, nr: 0xFF, size: 16383},
		{dir: 0, typ: 0x01, nr: 0x01, size: 1},
		{dir: 1, typ: 0x10, nr: 0x20, size: 64},
		{dir: 2, typ: 0x80, nr: 0x40, size: 1024},
		{dir: 3, typ: 0x55, nr: 0xAA, size: 8192},
		{dir: 1, typ: 0xFE, nr: 0xEF, size: 256},
		{dir: 2, typ: 0x33, nr: 0x77, size: 2048},
	}

	for _, test = range tests {
		goReq, err = ioctl.IOC(test.dir, test.typ, test.nr, test.size)
		if err != nil {
			t.Error(err)
		}

		cReq = ioctlcgo.IOC(test.dir, test.typ, test.nr, test.size)
		if cReq == goReq {
			continue
		}

		t.Errorf(
			"got: %d, exp: %d: dir = %d, typ = %d, nr = %d, size = %d",
			goReq,
			cReq,
			test.dir,
			test.typ,
			test.nr,
			test.size,
		)
	}
}
