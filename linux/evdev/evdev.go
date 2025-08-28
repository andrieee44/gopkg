// Package evdev provides a pure-Go interface to the Linux evdev (event
// device) subsystem. It lets you open input devices (keyboards, mice,
// gamepads, etc.), read raw input events, and interpret them as high-level
// actions.
package evdev

import (
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"unsafe"

	"github.com/andrieee44/gopkg/linux/uapi/input"
	"github.com/andrieee44/gopkg/linux/uapi/ioctl"
)

var (
	// ErrUnsupportedEvent is returned when the device does not support
	// the specified input event code. In such cases, callers should
	// invoke the underlying syscall directly.
	ErrUnsupportedEvent error = errors.New("unsupported event, use syscall instead")

	// ErrNotMultiTouch is returned when the provided absolute event code
	// does not correspond to any multitouch axis or slot index.
	ErrNotMultiTouch error = errors.New("is not a multitouch code")
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

// Devices attempts to open all discovered input devices and returns
// those that were opened successfully, along with an error slice
// containing the result of each attempt.
//
// The function does not stop at the first failure — it continues
// processing all paths and aggregates any errors into the errs slice.
func Devices() ([]*Device, []error) {
	var (
		devices []*Device
		device  *Device
		paths   []string
		path    string
		errs    []error
		err     error
	)

	paths, err = filepath.Glob("/dev/input/event*")
	if err != nil {
		return nil, []error{
			fmt.Errorf("failed to enumerate event devices at /dev/input/event*: %w", err),
		}
	}

	devices = make([]*Device, 0, len(paths))
	errs = make([]error, 0, len(paths))

	for _, path = range paths {
		device, err = NewDevice(path)
		if err != nil {
			errs = append(errs, err)
		} else {
			devices = append(devices, device)
		}
	}

	return devices, errs
}

func (dev *Device) Filename() string {
	return dev.file.Name()
}

// Fd returns the evdev device's underlying file descriptor.
func (dev *Device) Fd() uintptr {
	return dev.file.Fd()
}

// Version returns the evdev device's driver version.
func (dev *Device) Version() (int32, error) {
	return getAny(
		dev,
		input.EVIOCGVERSION,
		new(int32),
		"failed to get event device driver version",
	)
}

// ID returns the evdev device’s identifier.
func (dev *Device) ID() (input.ID, error) {
	return getAny(
		dev,
		input.EVIOCGID,
		new(input.ID),
		"failed to get event device ID",
	)
}

// Repeat returns the current key‐repeat parameters of the evdev device in
// milliseconds. The returned array holds two values:
// uint32[0] = Delay before key repeat starts.
// uint32[1] = Period between repeats when a key is held.
func (dev *Device) Repeat() ([2]uint32, error) {
	return getAny(
		dev,
		input.EVIOCGREP,
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
		dev,
		input.EVIOCSREP,
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
		dev,
		input.EVIOCGKEYCODE,
		&codes,
		"failed to get keycode of evdev device",
	)
}

// ScancodeV2 retrieves the keycode for a given keymap entry on this
// evdev device. The keymap parameter specifies the input index to query
// in its Index field and is updated with the retrieved keycode in its
// Keycode field.
func (dev *Device) ScancodeV2(keymap input.KeymapEntry) (input.KeymapEntry, error) {
	return getAny(
		dev,
		input.EVIOCGKEYCODE_V2,
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
		dev,
		input.EVIOCSKEYCODE,
		&codes,
		"failed to set keycode of evdev device",
	)
}

// SetScancodeV2 sets the keycode for a given keymap entry on the evdev
// device. The keymap parameter specifies the input index to map in its
// Index field and carries the keycode to assign in its Keycode field.
func (dev *Device) SetScancodeV2(keymap input.KeymapEntry) error {
	return setAny(
		dev,
		input.EVIOCSKEYCODE_V2,
		&keymap,
		"failed to set index keycodeV2 of evdev device",
	)
}

// Name returns the device’s name as a string.
// bufSize specifies the maximum number of bytes to read. If unsure,
// use 256.
func (dev *Device) Name(bufSize uint32) (string, error) {
	return getStr(
		dev,
		input.EVIOCGNAME,
		bufSize,
		"failed to get evdev device name",
	)
}

// PhysicalLocation returns the device’s physical location as a string.
// bufSize specifies the maximum number of bytes to read. If unsure,
// use 256.
func (dev *Device) PhysicalLocation(bufSize uint32) (string, error) {
	return getStr(
		dev,
		input.EVIOCGPHYS,
		bufSize,
		"failed to get evdev device physical location",
	)
}

// UniqueID returns the device’s unique identifier as a string.
// bufSize specifies the maximum number of bytes to read. If unsure,
// use 256.
func (dev *Device) UniqueID(bufSize uint32) (string, error) {
	return getStr(
		dev,
		input.EVIOCGUNIQ,
		bufSize,
		"failed to get evdev device unique id",
	)
}

// Properties returns the evdev device's input properties.
func (dev *Device) Properties() ([]input.PropCode, error) {
	return getBitmask(
		dev,
		input.EVIOCGPROP,
		input.INPUT_PROP_CNT,
		"failed to get evdev device properties",
	)
}

// MTSlotValues retrieves the current multitouch slot values for a given
// [input.AbsoluteCode].
func (dev *Device) MTSlotValues(abs input.AbsoluteCode) ([]int32, error) {
	var (
		absInfo input.AbsInfo
		length  uint32
		values  []int32
		err     error
	)

	if !input.IsMultiTouch(abs) {
		return nil, fmt.Errorf("%s: code %d (%s): %w", dev.Filename(), abs, abs, ErrNotMultiTouch)
	}

	absInfo, err = dev.AbsInfo(input.ABS_MT_SLOT)
	if err != nil {
		return nil, fmt.Errorf("failed to get evdev device multitouch slots: %w", err)
	}

	length = uint32(absInfo.Maximum) + 2
	values = make([]int32, length)
	values[0] = int32(abs)

	_, err = getAny(
		dev,
		func() (uint32, error) {
			return input.EVIOCGMTSLOTS(length * uint32(unsafe.Sizeof(int32(0))))
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
func (dev *Device) EnabledKeycodes() ([]input.KeyCode, error) {
	return getBitmask(
		dev,
		input.EVIOCGKEY,
		input.KEY_CNT,
		"failed to get evdev device enabled keycodes",
	)
}

// EnabledLEDs returns the evdev device's enabled LED codes.
func (dev *Device) EnabledLEDs() ([]input.LEDCode, error) {
	return getBitmask(
		dev,
		input.EVIOCGLED,
		input.LED_CNT,
		"failed to get evdev device enabled LEDs",
	)
}

// EnabledSounds returns the evdev device's enabled sound codes.
func (dev *Device) EnabledSounds() ([]input.SoundCode, error) {
	return getBitmask(
		dev,
		input.EVIOCGSND,
		input.SND_CNT,
		"failed to get evdev device enabled sounds",
	)
}

// EnabledSwitches returns the evdev device's enabled switch codes.
func (dev *Device) EnabledSwitches() ([]input.SwitchCode, error) {
	return getBitmask(
		dev,
		input.EVIOCGSW,
		input.SW_CNT,
		"failed to get evdev device enabled switches",
	)
}

// EnabledCodes returns the enabled codes for the given [input.EventCode]
// by dispatching to the matching Enabled* method and converting them to
// [input.Coder]. Only codes that are both supported by the device and
// currently enabled are returned. Errors from the underlying call are
// returned unchanged; unknown events return [ErrUnsupportedEvent].
//
// Using EnabledCodes erases the concrete code type, so callers lose
// compile‑time type safety; prefer specific methods such as
// [Device.EnabledKeycodes] or [Device.EnabledSwitches] when typed codes
// are required.
func (dev *Device) EnabledCodes(event input.EventCode) ([]input.Coder, error) {
	switch event {
	case input.EV_KEY:
		return asInputCoders(dev.EnabledKeycodes())
	case input.EV_SW:
		return asInputCoders(dev.EnabledSwitches())
	case input.EV_LED:
		return asInputCoders(dev.EnabledLEDs())
	case input.EV_SND:
		return asInputCoders(dev.EnabledSounds())
	default:
		return nil, fmt.Errorf(
			"%s: event %d (%s): %w",
			dev.Filename(),
			event,
			event,
			ErrUnsupportedEvent,
		)
	}
}

// EventCodes returns the evdev device's supported event codes.
func (dev *Device) EventCodes() ([]input.EventCode, error) {
	return getBitmask(
		dev,
		input.BitmaskReq(0),
		input.EV_CNT,
		"failed to get evdev device supported event codes",
	)
}

// SyncCodes returns the evdev device's supported sync codes.
func (dev *Device) SyncCodes() ([]input.SyncCode, error) {
	return getBitmask(
		dev,
		input.BitmaskReq(input.EV_SYN),
		input.SYN_CNT,
		"failed to get evdev device supported sync codes",
	)
}

// Keycodes returns the evdev device's supported keycodes.
func (dev *Device) Keycodes() ([]input.KeyCode, error) {
	return getBitmask(
		dev,
		input.BitmaskReq(input.EV_KEY),
		input.KEY_CNT,
		"failed to get evdev device supported keycodes",
	)
}

// RelativeCodes returns the evdev device's supported relative codes.
func (dev *Device) RelativeCodes() ([]input.RelativeCode, error) {
	return getBitmask(
		dev,
		input.BitmaskReq(input.EV_REL),
		input.REL_CNT,
		"failed to get evdev device supported relative codes",
	)
}

// AbsoluteCodes returns the evdev device's supported absolute codes.
func (dev *Device) AbsoluteCodes() ([]input.AbsoluteCode, error) {
	return getBitmask(
		dev,
		input.BitmaskReq(input.EV_ABS),
		input.ABS_CNT,
		"failed to get evdev device supported absolute codes",
	)
}

// MiscCodes returns the evdev device's supported misc codes.
func (dev *Device) MiscCodes() ([]input.MiscCode, error) {
	return getBitmask(
		dev,
		input.BitmaskReq(input.EV_MSC),
		input.MSC_CNT,
		"failed to get evdev device supported misc codes",
	)
}

// SwitchCodes returns the evdev device's supported switch codes.
func (dev *Device) SwitchCodes() ([]input.SwitchCode, error) {
	return getBitmask(
		dev,
		input.BitmaskReq(input.EV_SW),
		input.SW_CNT,
		"failed to get evdev device supported switch codes",
	)
}

// LEDCodes returns the evdev device's supported LED codes.
func (dev *Device) LEDCodes() ([]input.LEDCode, error) {
	return getBitmask(
		dev,
		input.BitmaskReq(input.EV_LED),
		input.LED_CNT,
		"failed to get evdev device supported LED codes",
	)
}

// SoundCodes returns the evdev device's supported sound codes.
func (dev *Device) SoundCodes() ([]input.SoundCode, error) {
	return getBitmask(
		dev,
		input.BitmaskReq(input.EV_SND),
		input.SND_CNT,
		"failed to get evdev device supported sound codes",
	)
}

// RepeatCodes returns the evdev device's supported repeat codes.
func (dev *Device) RepeatCodes() ([]input.RepeatCode, error) {
	var err error

	_, err = dev.Repeat()
	if err != nil {
		return nil, fmt.Errorf("failed to get evdev device supported repeat codes: %w", err)
	}

	return []input.RepeatCode{
		input.REP_DELAY,
		input.REP_PERIOD,
	}, nil
}

// FFCodes returns the evdev device's supported force-feedback codes.
func (dev *Device) FFCodes() ([]input.FFCode, error) {
	return getBitmask(
		dev,
		input.BitmaskReq(input.EV_FF),
		input.FF_CNT,
		"failed to get evdev device supported force-feedback codes",
	)
}

// PowerCodes returns the evdev device's supported power codes.
func (dev *Device) PowerCodes() ([]input.KeyCode, error) {
	return getBitmask(
		dev,
		input.BitmaskReq(input.EV_PWR),
		input.KEY_CNT,
		"failed to get evdev device supported power codes",
	)
}

// FFStatusCodes returns the evdev device's supported force-feedback status
// codes.
func (dev *Device) FFStatusCodes() ([]input.FFStatusCode, error) {
	return getBitmask(
		dev,
		input.BitmaskReq(input.EV_FF_STATUS),
		input.FF_STATUS_MAX,
		"failed to get evdev device supported force-feedback status codes",
	)
}

// Codes returns the supported codes for the given
// [input.EventCode] by dispatching to the matching *Codes
// method and converting them to [input.Coder]. Errors from the
// underlying call are returned unchanged; unknown events return
// [ErrUnsupportedEvent]. Using Codes erases the concrete code type,
// so callers lose compile‑time type safety; prefer specific methods
// such as [Device.Keycodes] or [Device.RelativeCodes] when typed
// codes are required.
func (dev *Device) Codes(event input.EventCode) ([]input.Coder, error) {
	switch event {
	case input.EV_SYN:
		return asInputCoders(dev.SyncCodes())
	case input.EV_KEY:
		return asInputCoders(dev.Keycodes())
	case input.EV_REL:
		return asInputCoders(dev.RelativeCodes())
	case input.EV_ABS:
		return asInputCoders(dev.AbsoluteCodes())
	case input.EV_MSC:
		return asInputCoders(dev.MiscCodes())
	case input.EV_SW:
		return asInputCoders(dev.SwitchCodes())
	case input.EV_LED:
		return asInputCoders(dev.LEDCodes())
	case input.EV_SND:
		return asInputCoders(dev.SoundCodes())
	case input.EV_REP:
		return asInputCoders(dev.RepeatCodes())
	case input.EV_FF:
		return asInputCoders(dev.FFCodes())
	case input.EV_PWR:
		return asInputCoders(dev.PowerCodes())
	case input.EV_FF_STATUS:
		return asInputCoders(dev.FFStatusCodes())
	default:
		return nil, fmt.Errorf(
			"%s: event %d (%s): %w",
			dev.Filename(),
			event,
			event,
			ErrUnsupportedEvent,
		)
	}
}

// AbsInfo returns the evdev device's absolute axis information
// corresponding to the provided [input.AbsoluteCode].
func (dev *Device) AbsInfo(abs input.AbsoluteCode) (input.AbsInfo, error) {
	return getAny(
		dev,
		func() (uint32, error) {
			return input.EVIOCGABS(abs)
		},
		new(input.AbsInfo),
		"failed to get evdev device absolute axis info",
	)
}

// SetAbsInfo sets the evdev device's absolute axis information
// corresponding to the provided [input.AbsoluteCode].
func (dev *Device) SetAbsInfo(abs input.AbsoluteCode, absInfo input.AbsInfo) error {
	return setAny(
		dev,
		func() (uint32, error) {
			return input.EVIOCSABS(abs)
		},
		&absInfo,
		"failed to set evdev device absolute axis info",
	)
}

// SendFF sends a force-feedback effect to the evdev device.
func (dev *Device) SendFF(effect input.FFEffect) error {
	return setAny(
		dev,
		input.EVIOCSFF,
		&effect,
		"failed to send evdev device force-feedback",
	)
}

// RemoveFF removes a force-feedback effect of the evdev device.
func (dev *Device) RemoveFF(id int32) error {
	return setAny(
		dev,
		input.EVIOCRMFF,
		&id,
		"failed to remove evdev device force-feedback",
	)
}

// FFEffects returns the amount of force-feedback effects of the evdev device.
func (dev *Device) FFEffects() (int32, error) {
	return getAny(
		dev,
		input.EVIOCGEFFECTS,
		new(int32),
		"failed to get evdev device force-feedback effects",
	)
}

// Grab toggles exclusive access to the evdev device.
// Pass 1 to grab, 0 to release.
func (dev *Device) Grab(grab int32) error {
	return setAny(
		dev,
		input.EVIOCGRAB,
		&grab,
		"failed to grab/release evdev device",
	)
}

// Release toggles exclusive access to the evdev device.
// Pass 1 to release, 0 to grab.
func (dev *Device) Release(release int32) error {
	return setAny(
		dev,
		input.EVIOCREVOKE,
		&release,
		"failed to release/grab evdev device",
	)
}

// SetEventMask configures which input events the device will report by
// applying the given event mask. The mask is a bitmask of event types and
// codes defined in [input.Mask].
func (dev *Device) SetEventMask(mask input.Mask) error {
	return setAny(
		dev,
		input.EVIOCSMASK,
		&mask,
		"failed to set evdev device event mask",
	)
}

// SetClockID sets the clock source used by the kernel when timestamping
// events read from the device. The clockID must be one of the standard
// clock constants.
func (dev *Device) SetClockID(clockID int32) error {
	return setAny(
		dev,
		input.EVIOCSCLOCKID,
		&clockID,
		"failed to set evdev device clock id",
	)
}

func (dev *Device) Snapshot() (*Snapshot, error) {
	var (
		info *Snapshot
		err  error
	)

	info, err = newSnapshot(dev)
	if err != nil {
		return nil, fmt.Errorf("failed to create evdev device snapshot: %w", err)
	}

	return info, nil
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

func getAny[T any](
	dev *Device,
	reqFn func() (uint32, error),
	arg *T,
	errMsg string,
) (T, error) {
	var (
		result T
		err    error
	)

	result, err = ioctl.GetAny(dev.Fd(), reqFn, arg)
	if err != nil {
		return *new(T), fmt.Errorf("%s: %s: %w", dev.Filename(), errMsg, err)
	}

	return result, nil
}

func setAny[T any](
	dev *Device,
	reqFn func() (uint32, error),
	arg *T,
	errMsg string,
) error {
	var err error

	_, err = getAny(dev, reqFn, arg, errMsg)
	if err != nil {
		return err
	}

	return nil
}

func getStr(
	dev *Device,
	reqFn func(length uint32) (uint32, error),
	bufSize uint32,
	errMsg string,
) (string, error) {
	var (
		str string
		err error
	)

	str, err = ioctl.GetStr(dev.Fd(), reqFn, bufSize)
	if err != nil {
		return "", fmt.Errorf("%s: %s: %w", dev.Filename(), errMsg, err)
	}

	return str, nil
}

func getBitmask[T input.Code](
	dev *Device,
	req func(length uint32) (uint32, error),
	count T,
	errMsg string,
) ([]T, error) {
	var (
		codes []T
		err   error
	)

	codes, err = input.GetBitmask(dev.Fd(), req, count)
	if err != nil {
		return nil, fmt.Errorf("%s: %s: %w", dev.Filename(), errMsg, err)
	}

	return codes, nil
}

func asInputCoders[T input.Code](codes []T, err error) ([]input.Coder, error) {
	if err != nil {
		return nil, err
	}

	return input.AsCoders(codes), nil
}
