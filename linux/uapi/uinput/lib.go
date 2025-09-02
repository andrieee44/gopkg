package uinput

import (
	"fmt"

	"github.com/andrieee44/gopkg/linux/uapi/ioctl"
)

func io(typ, nr uint32, errMsg string) (uint32, error) {
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

func ior[T any](typ, nr uint32, errMsg string) (uint32, error) {
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

func iow[T any](typ, nr uint32, errMsg string) (uint32, error) {
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

func iowr[T any](typ, nr uint32, errMsg string) (uint32, error) {
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

func ioc(dir, typ, nr, size uint32, errMsg string) (uint32, error) {
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
