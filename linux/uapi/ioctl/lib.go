//go:build linux

package ioctl

import (
	"fmt"
	"syscall"
	"unsafe"

	"github.com/andrieee44/gopkg/lib/xerr"
	"golang.org/x/sys/unix"
)

// GetAny performs an ioctl call on the given file descriptor using a
// request code from reqFn.
//
// It writes the result into arg and returns the dereferenced value on
// success. If obtaining the request code fails, or if the ioctl syscall
// returns an error, it returns the zero value of T and a wrapped error.
//
// Suitable for wrapping ioctl calls for any data type without repeating
// boilerplate syscall handling.
func GetAny[T any](
	fd uintptr,
	reqFn func() (uint32, error),
	arg *T,
) (T, error) {
	var (
		req   uint32
		errno syscall.Errno
		err   error
	)

	req, err = reqFn()
	if err != nil {
		return *new(T), fmt.Errorf("ioctl.GetAny: failed to get request code: %w", err)
	}

	_, _, errno = unix.Syscall(
		unix.SYS_IOCTL,
		fd,
		uintptr(req),
		uintptr(unsafe.Pointer(arg)),
	)
	if errno != 0 {
		return *new(T), fmt.Errorf("ioctl.GetAny: failed ioctl syscall: %w", errno)
	}

	return *arg, nil
}

// SetAny performs an ioctl call on the given file descriptor using a
// request code from reqFn, sending arg to the kernel.
//
// It discards any returned value and returns an error if obtaining the
// request code or performing the ioctl fails.
//
// Suitable for ioctl operations that write data without reading a result.
func SetAny[T any](fd uintptr, reqFn func() (uint32, error), arg *T) error {
	var err error

	_, err = GetAny(fd, reqFn, arg)

	return xerr.WrapIf("ioctl.SetAny", err)
}

// GetStr performs an ioctl call on the given file descriptor using a
// request code from reqFn, reading up to bufSize bytes into a string.
//
// It returns the string and a nil error on success. If obtaining the
// request code or performing the ioctl fails, it returns an empty string
// and a wrapped error.
func GetStr(
	fd uintptr,
	reqFn func(length uint32) (uint32, error),
	bufSize uint32,
) (string, error) {
	var (
		buf []byte
		err error
	)

	buf = make([]byte, bufSize)

	_, err = GetAny(fd, func() (uint32, error) {
		return reqFn(bufSize)
	}, &buf[0])
	if err != nil {
		return "", fmt.Errorf("ioctl.GetStr: %w", err)
	}

	return unix.ByteSliceToString(buf), nil
}

// Empty performs an ioctl call on the given file descriptor using a
// request code from reqFn, with no argument.
//
// It returns an error if obtaining the request code or performing the
// ioctl fails.
func Empty(fd uintptr, reqFn func() (uint32, error)) error {
	return xerr.WrapIf("ioctl.Empty", SetAny(fd, reqFn, new(struct{})))
}

func ioc[T any](dir, typ, nr uint32, fnName string) (uint32, error) {
	var (
		size, req uint32
		err       error
	)

	size, err = IOC_TYPECHECK[T]()
	if err != nil {
		return 0, fmt.Errorf("%s: %w", fnName, err)
	}

	req, err = IOC(dir, typ, nr, size)
	if err != nil {
		return 0, fmt.Errorf("%s: %w", fnName, err)
	}

	return req, nil
}
