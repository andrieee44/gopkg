// Package ioctlwrap provides wrappers around the [ioctl] package that add
// consistent error message formatting for various ioctl request code
// generators.
package ioctlwrap

import (
	"fmt"
	"os"

	"github.com/andrieee44/gopkg/linux/uapi/ioctl"
)

// IO calls [ioctl.IO] to generate an ioctl request code without data
// transfer. It returns the computed request code or an error wrapped with
// errMsg.
func IO(typ, nr uint32, errMsg string) (uint32, error) {
	var (
		req uint32
		err error
	)

	req, err = ioctl.IO(typ, nr)
	if err != nil {
		return 0, fmt.Errorf("%s: %w", errMsg, err)
	}

	return req, nil
}

// IOR calls [ioctl.IOR] to generate an ioctl request code for reading data
// into a structure of type T. It returns the computed request code or an
// error wrapped with errMsg.
func IOR[T any](typ, nr uint32, errMsg string) (uint32, error) {
	var (
		req uint32
		err error
	)

	req, err = ioctl.IOR[T](typ, nr)
	if err != nil {
		return 0, fmt.Errorf("%s: %w", errMsg, err)
	}

	return req, nil
}

// IOW calls [ioctl.IOW] to generate an ioctl request code for writing data
// from a structure of type T. It returns the computed request code or an
// error wrapped with errMsg.
func IOW[T any](typ, nr uint32, errMsg string) (uint32, error) {
	var (
		req uint32
		err error
	)

	req, err = ioctl.IOW[T](typ, nr)
	if err != nil {
		return 0, fmt.Errorf("%s: %w", errMsg, err)
	}

	return req, nil
}

// IOWR calls [ioctl.IOWR] to generate an ioctl request code for reading and
// writing data to and from a structure of type T. It returns the computed
// request code or an error wrapped with errMsg.
func IOWR[T any](typ, nr uint32, errMsg string) (uint32, error) {
	var (
		req uint32
		err error
	)

	req, err = ioctl.IOWR[T](typ, nr)
	if err != nil {
		return 0, fmt.Errorf("%s: %w", errMsg, err)
	}

	return req, nil
}

// IOC calls [ioctl.IOC] to generate a fully specified ioctl request code with
// explicit direction, type, command number, and size. It returns the computed
// request code or an error wrapped with errMsg.
func IOC(dir, typ, nr, size uint32, errMsg string) (uint32, error) {
	var (
		req uint32
		err error
	)

	req, err = ioctl.IOC(dir, typ, nr, size)
	if err != nil {
		return 0, fmt.Errorf("%s: %w", errMsg, err)
	}

	return req, nil
}

// GetAny wraps [ioctl.GetAny] and wraps the returned error with the file
// name and a custom message.
func GetAny[T any](
	file *os.File,
	reqFn func() (uint32, error),
	arg *T,
	errMsg string,
) (T, error) {
	var (
		result T
		err    error
	)

	result, err = ioctl.GetAny(file.Fd(), reqFn, arg)
	if err != nil {
		return *new(T), fmt.Errorf("%s: %s: %w", file.Name(), errMsg, err)
	}

	return result, nil
}

// SetAny wraps GetAny but discards its return value, returning only the
// error.
func SetAny[T any](
	file *os.File,
	reqFn func() (uint32, error),
	arg *T,
	errMsg string,
) error {
	var err error

	_, err = GetAny(file, reqFn, arg, errMsg)
	if err != nil {
		return err
	}

	return nil
}

// GetStr wraps [ioctl.GetStr] and wraps the returned error with the file
// name and a custom message.
func GetStr(
	file *os.File,
	reqFn func(length uint32) (uint32, error),
	bufSize uint32,
	errMsg string,
) (string, error) {
	var (
		str string
		err error
	)

	str, err = ioctl.GetStr(file.Fd(), reqFn, bufSize)
	if err != nil {
		return "", fmt.Errorf("%s: %s: %w", file.Name(), errMsg, err)
	}

	return str, nil
}

// Empty wraps [ioctl.Empty] and prewraps any returned error with the file
// name and a custom message.
func Empty(file *os.File, reqFn func() (uint32, error), errMsg string) error {
	var err error

	err = ioctl.Empty(file.Fd(), reqFn)
	if err != nil {
		return fmt.Errorf("%s: %s: %w", file.Name(), errMsg, err)
	}

	return nil
}
