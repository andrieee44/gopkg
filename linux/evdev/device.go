package evdev

import (
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"unsafe"

	"github.com/andrieee44/gopkg/linux/uapi/evdev"
	"github.com/andrieee44/gopkg/linux/uapi/ioctl"
)

var (
	// ErrUnsupportedEvent is returned when the device does not support
	// the specified input event code. In such cases, callers should
	// invoke the underlying syscall directly.
	ErrUnsupportedEvent error = errors.New("unsupported event, use syscall instead")

	// ErrNotMultitouch is returned when the provided absolute event code
	// does not correspond to any multitouch axis or slot index.
	ErrNotMultitouch error = errors.New("is not a multitouch code")
)

// Device represents an evdev device.
// It wraps the opened /dev/input/eventN file.
type Device struct {
	file *os.File
}

// NewDevice opens the evdev device at the given path and returns a [Device].
// The device file is opened in read-write mode. The caller is responsible
// for closing the device when no longer needed.
func NewDevice(path string) (*Device, error) {
	var (
		device *Device
		file   *os.File
		err    error
	)

	file, err = os.OpenFile(filepath.Clean(path), os.O_RDWR, 0)
	if err != nil {
		return nil, fmt.Errorf("failed to open evdev device: %w", err)
	}

	device = &Device{
		file: file,
	}

	return device, nil
}

// Devices enumerates /dev/input/event* for event devices, opens each one,
// and returns a slice of [Device] pointers. If any device fails to open,
// an error is returned and no devices are returned.
func Devices() ([]*Device, error) {
	var (
		devices []*Device
		device  *Device
		paths   []string
		path    string
		err     error
	)

	paths, err = filepath.Glob("/dev/input/event*")
	if err != nil {
		return nil, fmt.Errorf("failed to enumerate event devices at /dev/input/event*: %w", err)
	}

	devices = make([]*Device, 0, len(paths))
	for _, path = range paths {
		device, err = NewDevice(path)
		if err != nil {
			return nil, err
		}

		devices = append(devices, device)
	}

	return devices, nil
}

// Fd returns the evdev device's underlying file descriptor.
func (dev *Device) Fd() uintptr {
	return dev.file.Fd()
}

// Version returns the evdev device's driver version.
func (dev *Device) Version() (int32, error) {
	return getAny(
		dev.Fd(),
		evdev.EVIOCGVERSION,
		new(int32),
		"failed to get event device driver version",
	)
}

// ID returns the evdev device’s identifier.
func (dev *Device) ID() (evdev.InputID, error) {
	return getAny(
		dev.Fd(),
		evdev.EVIOCGID,
		new(evdev.InputID),
		"failed to get event device ID",
	)
}

// Repeat returns the current key‐repeat parameters of the evdev device in
// milliseconds. The returned array holds two values:
// uint32[0] = Delay before key repeat starts.
// uint32[1] = Period between repeats when a key is held.
func (dev *Device) Repeat() ([2]uint32, error) {
	return getAny(
		dev.Fd(),
		evdev.EVIOCGREP,
		new([2]uint32),
		"failed to get repeat settings of evdev device",
	)
}

// SetRepeat sets the key‐repeat parameters of the evdev device in
// milliseconds. The settings array holds two values:
// uint32[0] = Delay before key repeat starts.
// uint32[1] = Period between repeats when a key is held.
func (dev *Device) SetRepeat(settings [2]uint32) error {
	return setAny(
		dev.Fd(),
		evdev.EVIOCSREP,
		&settings,
		"failed to set repeat settings of evdev device",
	)
}

// Scancode retrieves the keycode for a scancode on the evdev device.
// The codes array holds two values:
// uint32[0] = Scancode to look up.
// uint32[1] = Populated keycode for the given scancode.
func (dev *Device) Scancode(codes [2]uint32) ([2]uint32, error) {
	return getAny(
		dev.Fd(),
		evdev.EVIOCGKEYCODE,
		&codes,
		"failed to get keycode of evdev device",
	)
}

// ScancodeV2 retrieves the keycode for a given keymap entry on this
// evdev device. The keymap parameter specifies the input index to query
// in its Index field and is updated with the retrieved keycode in its
// Keycode field.
func (dev *Device) ScancodeV2(keymap evdev.InputKeymapEntry) (evdev.InputKeymapEntry, error) {
	return getAny(
		dev.Fd(),
		evdev.EVIOCGKEYCODE_V2,
		&keymap,
		"failed to get keycodeV2 of evdev device",
	)
}

// SetScancode sets a keycode mapping for a scancode on the evdev device.
// The codes array holds two values:
// uint32[0] = Scancode to map.
// uint32[1] = Keycode to assign to that scancode.
func (dev *Device) SetScancode(codes [2]uint32) error {
	return setAny(
		dev.Fd(),
		evdev.EVIOCSKEYCODE,
		&codes,
		"failed to set keycode of evdev device",
	)
}

// SetScancodeV2 sets the keycode for a given keymap entry on the evdev
// device. The keymap parameter specifies the input index to map in its
// Index field and carries the keycode to assign in its Keycode field.
func (dev *Device) SetScancodeV2(keymap evdev.InputKeymapEntry) error {
	return setAny(
		dev.Fd(),
		evdev.EVIOCSKEYCODE_V2,
		&keymap,
		"failed to set index keycodeV2 of evdev device",
	)
}

// Name returns the device’s name as a string.
// bufSize specifies the maximum number of bytes to read. If unsure,
// use 256.
func (dev *Device) Name(bufSize uint32) (string, error) {
	return getStr(
		dev.Fd(),
		evdev.EVIOCGNAME,
		bufSize,
		"failed to get evdev device name",
	)
}

// PhysicalLocation returns the device’s physical location as a string.
// bufSize specifies the maximum number of bytes to read. If unsure,
// use 256.
func (dev *Device) PhysicalLocation(bufSize uint32) (string, error) {
	return getStr(
		dev.Fd(),
		evdev.EVIOCGPHYS,
		bufSize,
		"failed to get evdev device physical location",
	)
}

// UniqueID returns the device’s unique identifier as a string.
// bufSize specifies the maximum number of bytes to read. If unsure,
// use 256.
func (dev *Device) UniqueID(bufSize uint32) (string, error) {
	return getStr(
		dev.Fd(),
		evdev.EVIOCGUNIQ,
		bufSize,
		"failed to get evdev device unique id",
	)
}

// Properties returns the evdev device's input properties.
func (dev *Device) Properties() ([]evdev.InputPropCode, error) {
	return getBitmask(
		dev.Fd(),
		evdev.EVIOCGPROP,
		evdev.INPUT_PROP_CNT,
		"failed to get evdev device properties",
	)
}

// MTSlotValues retrieves the current multitouch slot values for a given
// [evdev.ABS_MT_SLOT] code.
func (dev *Device) MTSlotValues(abs evdev.InputAbsoluteCode) ([]int32, error) {
	var (
		absInfo evdev.InputAbsInfo
		length  uint32
		values  []int32
		err     error
	)

	if !evdev.IsMultiTouch(abs) {
		return nil, fmt.Errorf("code %d (%s): %w", abs, abs, ErrNotMultitouch)
	}

	absInfo, err = dev.AbsInfo(evdev.ABS_MT_SLOT)
	if err != nil {
		return nil, fmt.Errorf("failed to get evdev device multitouch slots: %w", err)
	}

	length = uint32(absInfo.Maximum) + 2
	values = make([]int32, length)
	values[0] = int32(abs)

	_, err = getAny(
		dev.Fd(),
		func() (uint32, error) {
			return evdev.EVIOCGMTSLOTS(length * uint32(unsafe.Sizeof(int32(0))))
		},
		&values[0],
		"failed to get evdev device multitouch slots",
	)
	if err != nil {
		return nil, err
	}

	return values[1:], nil
}

// EnabledKeycodes returns the evdev device's enabled key codes.
func (dev *Device) EnabledKeycodes() ([]evdev.InputKeyCode, error) {
	return getBitmask(
		dev.Fd(),
		evdev.EVIOCGKEY,
		evdev.KEY_CNT,
		"failed to get evdev device enabled keycodes",
	)
}

// EnabledLEDs returns the evdev device's enabled LED codes.
func (dev *Device) EnabledLEDs() ([]evdev.InputLEDCode, error) {
	return getBitmask(
		dev.Fd(),
		evdev.EVIOCGLED,
		evdev.LED_CNT,
		"failed to get evdev device enabled LEDs",
	)
}

// EnabledSounds returns the evdev device's enabled sound codes.
func (dev *Device) EnabledSounds() ([]evdev.InputSoundCode, error) {
	return getBitmask(
		dev.Fd(),
		evdev.EVIOCGSND,
		evdev.SND_CNT,
		"failed to get evdev device enabled sounds",
	)
}

// EnabledSwitches returns the evdev device's enabled switch codes.
func (dev *Device) EnabledSwitches() ([]evdev.InputSwitchCode, error) {
	return getBitmask(
		dev.Fd(),
		evdev.EVIOCGSW,
		evdev.SW_CNT,
		"failed to get evdev device enabled switches",
	)
}

// Codes returns the supported codes for the given
// [evdev.InputEventCode] by dispatching to the matching *Codes
// method and converting them to [evdev.InputCoder]. Errors from the
// underlying call are returned unchanged; unknown events return
// [ErrUnsupportedEvent]. Using Codes erases the concrete code type,
// so callers lose compile‑time type safety; prefer specific methods
// such as [Device.Keycodes] or [Device.RelativeCodes] when typed
// codes are required.
func (dev *Device) Codes(event evdev.InputEventCode) ([]evdev.InputCoder, error) {
	switch event {
	case evdev.EV_SYN:
		return asInputCoders(dev.SyncCodes())
	case evdev.EV_KEY:
		return asInputCoders(dev.Keycodes())
	case evdev.EV_REL:
		return asInputCoders(dev.RelativeCodes())
	case evdev.EV_ABS:
		return asInputCoders(dev.AbsoluteCodes())
	case evdev.EV_MSC:
		return asInputCoders(dev.MiscCodes())
	case evdev.EV_SW:
		return asInputCoders(dev.SwitchCodes())
	case evdev.EV_LED:
		return asInputCoders(dev.LEDCodes())
	case evdev.EV_SND:
		return asInputCoders(dev.SoundCodes())
	case evdev.EV_REP:
		return asInputCoders(dev.RepeatCodes())
	case evdev.EV_FF:
		return asInputCoders(dev.FFCodes())
	case evdev.EV_PWR:
		return asInputCoders(dev.PowerCodes())
	case evdev.EV_FF_STATUS:
		return asInputCoders(dev.FFStatusCodes())
	default:
		return nil, fmt.Errorf("event %d (%s): %w", event, event, ErrUnsupportedEvent)
	}
}

// EventCodes returns the evdev device's supported event codes.
func (dev *Device) EventCodes() ([]evdev.InputEventCode, error) {
	return getBitmask(
		dev.Fd(),
		evdev.BitmaskReq(0),
		evdev.EV_CNT,
		"failed to get evdev device supported event codes",
	)
}

// SyncCodes returns the evdev device's supported sync codes.
func (dev *Device) SyncCodes() ([]evdev.InputSyncCode, error) {
	return getBitmask(
		dev.Fd(),
		evdev.BitmaskReq(evdev.EV_SYN),
		evdev.SYN_CNT,
		"failed to get evdev device supported sync codes",
	)
}

// Keycodes returns the evdev device's supported keycodes.
func (dev *Device) Keycodes() ([]evdev.InputKeyCode, error) {
	return getBitmask(
		dev.Fd(),
		evdev.BitmaskReq(evdev.EV_KEY),
		evdev.KEY_CNT,
		"failed to get evdev device supported keycodes",
	)
}

// RelativeCodes returns the evdev device's supported relative codes.
func (dev *Device) RelativeCodes() ([]evdev.InputRelativeCode, error) {
	return getBitmask(
		dev.Fd(),
		evdev.BitmaskReq(evdev.EV_REL),
		evdev.REL_CNT,
		"failed to get evdev device supported relative codes",
	)
}

// AbsoluteCodes returns the evdev device's supported absolute codes.
func (dev *Device) AbsoluteCodes() ([]evdev.InputAbsoluteCode, error) {
	return getBitmask(
		dev.Fd(),
		evdev.BitmaskReq(evdev.EV_ABS),
		evdev.ABS_CNT,
		"failed to get evdev device supported absolute codes",
	)
}

// MiscCodes returns the evdev device's supported misc codes.
func (dev *Device) MiscCodes() ([]evdev.InputMiscCode, error) {
	return getBitmask(
		dev.Fd(),
		evdev.BitmaskReq(evdev.EV_MSC),
		evdev.MSC_CNT,
		"failed to get evdev device supported misc codes",
	)
}

// SwitchCodes returns the evdev device's supported switch codes.
func (dev *Device) SwitchCodes() ([]evdev.InputSwitchCode, error) {
	return getBitmask(
		dev.Fd(),
		evdev.BitmaskReq(evdev.EV_SW),
		evdev.SW_CNT,
		"failed to get evdev device supported switch codes",
	)
}

// LEDCodes returns the evdev device's supported LED codes.
func (dev *Device) LEDCodes() ([]evdev.InputLEDCode, error) {
	return getBitmask(
		dev.Fd(),
		evdev.BitmaskReq(evdev.EV_LED),
		evdev.LED_CNT,
		"failed to get evdev device supported LED codes",
	)
}

// SoundCodes returns the evdev device's supported sound codes.
func (dev *Device) SoundCodes() ([]evdev.InputSoundCode, error) {
	return getBitmask(
		dev.Fd(),
		evdev.BitmaskReq(evdev.EV_SND),
		evdev.SND_CNT,
		"failed to get evdev device supported sound codes",
	)
}

// RepeatCodes returns the evdev device's supported repeat codes.
func (dev *Device) RepeatCodes() ([]evdev.InputRepeatCode, error) {
	var err error

	_, err = dev.Repeat()
	if err != nil {
		return nil, fmt.Errorf("failed to get evdev device supported repeat codes: %w", err)
	}

	return []evdev.InputRepeatCode{
		evdev.REP_DELAY,
		evdev.REP_PERIOD,
	}, nil
}

// FFCodes returns the evdev device's supported force-feedback codes.
func (dev *Device) FFCodes() ([]evdev.InputFFCode, error) {
	return getBitmask(
		dev.Fd(),
		evdev.BitmaskReq(evdev.EV_FF),
		evdev.FF_CNT,
		"failed to get evdev device supported force-feedback codes",
	)
}

// PowerCodes returns the evdev device's supported power codes.
func (dev *Device) PowerCodes() ([]evdev.InputKeyCode, error) {
	return getBitmask(
		dev.Fd(),
		evdev.BitmaskReq(evdev.EV_PWR),
		evdev.KEY_CNT,
		"failed to get evdev device supported power codes",
	)
}

// FFStatusCodes returns the evdev device's supported force-feedback status
// codes.
func (dev *Device) FFStatusCodes() ([]evdev.InputFFStatusCode, error) {
	return getBitmask(
		dev.Fd(),
		evdev.BitmaskReq(evdev.EV_FF_STATUS),
		evdev.FF_STATUS_MAX,
		"failed to get evdev device supported force-feedback status codes",
	)
}

// AbsInfo returns the evdev device's absolute axis information
// corresponding to the provided [evdev.InputAbsoluteCode].
func (dev *Device) AbsInfo(abs evdev.InputAbsoluteCode) (evdev.InputAbsInfo, error) {
	return getAny(
		dev.Fd(),
		func() (uint32, error) {
			return evdev.EVIOCGABS(abs)
		},
		new(evdev.InputAbsInfo),
		"failed to get evdev device absolute axis info",
	)
}

// SetAbsInfo sets the evdev device's absolute axis information
// corresponding to the provided [evdev.InputAbsoluteCode].
func (dev *Device) SetAbsInfo(abs evdev.InputAbsoluteCode, absInfo evdev.InputAbsInfo) error {
	return setAny(
		dev.Fd(),
		func() (uint32, error) {
			return evdev.EVIOCSABS(abs)
		},
		&absInfo,
		"failed to set evdev device absolute axis info",
	)
}

// SendFF sends a force-feedback effect to the evdev device.
func (dev *Device) SendFF(effect evdev.FFEffect) error {
	return setAny(
		dev.Fd(),
		evdev.EVIOCSFF,
		&effect,
		"failed to send evdev device force-feedback",
	)
}

// RemoveFF removes a force-feedback effect of the evdev device.
func (dev *Device) RemoveFF(id int32) error {
	return setAny(
		dev.Fd(),
		evdev.EVIOCRMFF,
		&id,
		"failed to remove evdev device force-feedback",
	)
}

// FFEffects returns the amount of force-feedback effects of the evdev device.
func (dev *Device) FFEffects() (int32, error) {
	return getAny(
		dev.Fd(),
		evdev.EVIOCGEFFECTS,
		new(int32),
		"failed to get evdev device force-feedback effects",
	)
}

// Grab toggles exclusive access to the evdev device.
// Pass 1 to grab, 0 to release.
func (dev *Device) Grab(grab int32) error {
	return setAny(
		dev.Fd(),
		evdev.EVIOCGRAB,
		&grab,
		"failed to grab/release evdev device",
	)
}

// Release toggles exclusive access to the evdev device.
// Pass 1 to release, 0 to grab.
func (dev *Device) Release(release int32) error {
	return setAny(
		dev.Fd(),
		evdev.EVIOCREVOKE,
		&release,
		"failed to release/grab evdev device",
	)
}

// SetEventMask configures which input events the device will report by
// applying the given event mask. The mask is a bitmask of event types and
// codes defined in [evdev.InputMask].
func (dev *Device) SetEventMask(mask evdev.InputMask) error {
	return setAny(
		dev.Fd(),
		evdev.EVIOCSMASK,
		&mask,
		"failed to set evdev device event mask",
	)
}

// SetClockID sets the clock source used by the kernel when timestamping
// events read from the device. The clockID must be one of the standard
// clock constants.
func (dev *Device) SetClockID(clockID int32) error {
	return setAny(
		dev.Fd(),
		evdev.EVIOCSCLOCKID,
		&clockID,
		"failed to set evdev device clock id",
	)
}

// Close closes the evdev device by closing its underlying file handle.
func (dev *Device) Close() error {
	var err error

	err = dev.file.Close()
	if err != nil {
		return fmt.Errorf("failed to close event device: %w", err)
	}

	return nil
}

// getAny wraps [ioctl.GetAny], prefixing any error with errMsg.
func getAny[T any](
	fd uintptr,
	reqFn func() (uint32, error),
	arg *T,
	errMsg string,
) (T, error) {
	var (
		result T
		err    error
	)

	result, err = ioctl.GetAny(fd, reqFn, arg)
	if err != nil {
		return *new(T), fmt.Errorf("%s: %w", errMsg, err)
	}

	return result, nil
}

// setAny wraps [ioctl.GetAny], prefixing any error with errMsg.
func setAny[T any](
	fd uintptr,
	reqFn func() (uint32, error),
	arg *T,
	errMsg string,
) error {
	var err error

	_, err = ioctl.GetAny(fd, reqFn, arg)
	if err != nil {
		return fmt.Errorf("%s: %w", errMsg, err)
	}

	return nil
}

// getStr wraps [ioctl.GetStr], prefixing any error with errMsg.
func getStr(
	fd uintptr,
	reqFn func(length uint32) (uint32, error),
	bufSize uint32,
	errMsg string,
) (string, error) {
	var (
		str string
		err error
	)

	str, err = ioctl.GetStr(fd, reqFn, bufSize)
	if err != nil {
		return "", fmt.Errorf("%s: %w", errMsg, err)
	}

	return str, nil
}

// getBitmask wraps [ioctl.GetBitmask], prefixing any error with errMsg.
func getBitmask[T evdev.InputCode](
	fd uintptr,
	req func(length uint32) (uint32, error),
	count T,
	errMsg string,
) ([]T, error) {
	var (
		codes []T
		err   error
	)

	codes, err = evdev.GetBitmask(fd, req, count)
	if err != nil {
		return nil, fmt.Errorf("%s: %w", errMsg, err)
	}

	return codes, nil
}

// asInputCoders wraps [evdev.AsInputCoders], returning any error passed in.
func asInputCoders[T evdev.InputCode](codes []T, err error) ([]evdev.InputCoder, error) {
	if err != nil {
		return nil, err
	}

	return evdev.AsInputCoders(codes), nil
}
