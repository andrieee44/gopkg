// Package inputwrap provides wrappers around the [input] package that add
// consistent error message formatting for various input-related ioctl
// operations and helpers for working with input codes.
package inputwrap

import (
	"fmt"
	"os"

	"github.com/andrieee44/gopkg/linux/uapi/input"
)

// GetBitmask wraps [input.GetBitmask] and wraps the returned error with
// the file name and a custom message.
func GetBitmask[T input.Code](
	file *os.File,
	req func(length uint32) (uint32, error),
	count T,
	errMsg string,
) ([]T, error) {
	var (
		codes []T
		err   error
	)

	codes, err = input.GetBitmask(file.Fd(), req, count)
	if err != nil {
		return nil, fmt.Errorf("%s: %s: %w", file.Name(), errMsg, err)
	}

	return codes, nil
}

// AsInputCoders calls fn to obtain codes, returning them converted via
// [input.AsCoders] or the error from fn.
func AsInputCoders[T input.Code](fn func() ([]T, error)) ([]input.Coder, error) {
	var (
		codes []T
		err   error
	)

	codes, err = fn()
	if err != nil {
		return nil, err
	}

	return input.AsCoders(codes), nil
}
