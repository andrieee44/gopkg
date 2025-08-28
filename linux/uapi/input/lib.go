package input

import (
	"fmt"
	"math"

	"github.com/andrieee44/gopkg/lib/bitops"
	"github.com/andrieee44/gopkg/linux/uapi/ioctl"
)

// GetBitmask performs an ioctl call on fd using a request code returned by
// req, reading a bitmask of up to count bits from the kernel and returning
// the codes whose bits are set. It returns a slice of T values representing
// the bits that were set. An error is returned if obtaining the request
// code, performing the ioctl call, or sizing the buffer fails.
func GetBitmask[T Code](
	fd uintptr,
	req func(length uint32) (uint32, error),
	count T,
) ([]T, error) {
	var (
		buf   []byte
		codes []T
		code  T
		err   error
	)

	buf = bitops.Bytes(count)
	if len(buf) > math.MaxUint32 {
		return nil, fmt.Errorf("buf length is %d: %w", len(buf), ioctl.ErrSizeOverflow)
	}

	_, err = ioctl.GetAny(fd, func() (uint32, error) {
		return req(uint32(len(buf)))
	}, &buf[0])
	if err != nil {
		return nil, fmt.Errorf("failed to get bitmask: %w", err)
	}

	codes = make([]T, 0, count)

	for code = range count {
		if !bitops.Test(buf, code) {
			continue
		}

		codes = append(codes, code)
	}

	return codes, nil
}

// BitmaskReq returns a closure that calls [EVIOCGBIT] for the given
// event code.
func BitmaskReq(event EventCode) func(uint32) (uint32, error) {
	return func(length uint32) (uint32, error) {
		return EVIOCGBIT(event, length)
	}
}

// IsMultiTouch reports whether the given absolute axis code represents a
// multi-touch event, such as contact dimensions, position, tool type, or
// tracking information. Returns true for ABS_MT_* codes, false otherwise.
func IsMultiTouch(abs AbsoluteCode) bool {
	switch abs {
	case ABS_MT_TOUCH_MAJOR,
		ABS_MT_TOUCH_MINOR,
		ABS_MT_WIDTH_MAJOR,
		ABS_MT_WIDTH_MINOR,
		ABS_MT_ORIENTATION,
		ABS_MT_POSITION_X,
		ABS_MT_POSITION_Y,
		ABS_MT_TOOL_TYPE,
		ABS_MT_BLOB_ID,
		ABS_MT_TRACKING_ID,
		ABS_MT_PRESSURE,
		ABS_MT_DISTANCE,
		ABS_MT_TOOL_X,
		ABS_MT_TOOL_Y:
		return true
	default:
		return false
	}
}

// AsCoders converts a slice of type T to a slice of Coder.
func AsCoders[T Code](codes []T) []Coder {
	var (
		coders []Coder
		idx    int
		value  T
	)

	coders = make([]Coder, len(codes))
	for idx, value = range codes {
		coders[idx] = Coder(value)
	}

	return coders
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
