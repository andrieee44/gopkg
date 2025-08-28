package ioctl

import (
	"fmt"
	"syscall"
	"unsafe"

	"golang.org/x/sys/unix"
)

// GetAny performs an ioctl call on fd using a request code from reqFn,
// writing the result into arg. Returns arg on success, or the zero value of T
// and an error if obtaining the request code or the ioctl itself fails.
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
		return *new(T), fmt.Errorf("failed to get request code: %w", err)
	}

	_, _, errno = unix.Syscall(
		unix.SYS_IOCTL,
		fd,
		uintptr(req),
		uintptr(unsafe.Pointer(arg)),
	)
	if errno != 0 {
		return *new(T), fmt.Errorf("failed ioctl syscall: %w", errno)
	}

	return *arg, nil
}

// SetAny performs an ioctl call on fd using a request code from reqFn,
// sending arg to the kernel and discarding any returned value.
// Returns an error if obtaining the request code or the ioctl itself fails.
// Useful for ioctl operations that write data without reading a result back.
func SetAny[T any](fd uintptr, reqFn func() (uint32, error), arg *T) error {
	var err error

	_, err = GetAny(fd, reqFn, arg)
	if err != nil {
		return err
	}

	return nil
}

// GetStr performs an ioctl call on fd using a request code from reqFn,
// reading up to bufSize bytes from the kernel into a string result.
// Returns the string and a nil error on success.
// An error is returned if obtaining the request code or the ioctl call fails.
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
		return "", err
	}

	return unix.ByteSliceToString(buf), nil
}

func ioc[T any](dir, typ, nr uint32) (uint32, error) {
	var (
		size uint32
		err  error
	)

	size, err = IOC_TYPECHECK[T]()
	if err != nil {
		return 0, err
	}

	return IOC(dir, typ, nr, size)
}
