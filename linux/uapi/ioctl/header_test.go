//go:build linux

package ioctl_test

import (
	"testing"

	"github.com/andrieee44/gopkg/linux/uapi/internal/uapicgo"
	"github.com/andrieee44/gopkg/linux/uapi/ioctl"
)

func TestConstants(t *testing.T) {
	t.Parallel()
	uapicgo.IoctlSetup()
	uapicgo.TestConstants(t, uapicgo.IoctlConstants())
}

func TestIOC(t *testing.T) {
	type table struct {
		dir, typ, nr, size uint32
	}

	var (
		tests       []table
		test        table
		goReq, cReq uint32
		err         error
	)

	t.Parallel()
	uapicgo.IoctlSetup()

	tests = []table{
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

		cReq = uapicgo.IOC(test.dir, test.typ, test.nr, test.size)
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
