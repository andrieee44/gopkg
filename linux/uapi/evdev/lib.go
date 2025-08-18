package evdev

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
func GetBitmask[T InputCode](
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
		return nil, err
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

// BitmaskRequest returns a closure that calls [EVIOCGBIT] for the given
// event code.
func BitmaskReq(event InputEventCode) func(uint32) (uint32, error) {
	return func(length uint32) (uint32, error) {
		return EVIOCGBIT(event, length)
	}
}

// IsMultiTouch reports whether the given absolute axis code represents a
// multi-touch event, such as contact dimensions, position, tool type, or
// tracking information. Returns true for ABS_MT_* codes, false otherwise.
func IsMultiTouch(abs InputAbsoluteCode) bool {
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
