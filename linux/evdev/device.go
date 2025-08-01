package evdev

import (
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"syscall"
	"unsafe"

	"github.com/andrieee44/gopkg/lib/bits"
	"github.com/andrieee44/gopkg/linux/uapi"
	"golang.org/x/sys/unix"
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

// NewDevice opens the evdev device at the given path and returns a Device.
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
// and returns a slice of Device pointers. If any device fails to open,
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
	return ioctlGetAny(
		dev.Fd(),
		uapi.EVIOCGVERSION(),
		new(int32),
		"failed to get event device driver version",
	)
}

// ID returns the evdev device’s identifier.
func (dev *Device) ID() (uapi.InputID, error) {
	return ioctlGetAny(
		dev.Fd(),
		uapi.EVIOCGID(),
		new(uapi.InputID),
		"failed to get event device ID",
	)
}

// Repeat returns the current key‐repeat parameters of the evdev device in
// milliseconds. The returned array holds two values:
// uint32[0] = Delay before key repeat starts.
// uint32[1] = Period between repeats when a key is held.
func (dev *Device) Repeat() ([2]uint32, error) {
	return ioctlGetAny(
		dev.Fd(),
		uapi.EVIOCGREP(),
		new([2]uint32),
		"failed to get repeat settings of evdev device",
	)
}

// SetRepeat sets the key‐repeat parameters of the evdev device in
// milliseconds. The settings array holds two values:
// uint32[0] = Delay before key repeat starts.
// uint32[1] = Period between repeats when a key is held.
func (dev *Device) SetRepeat(settings [2]uint32) error {
	return ioctlSetAny(
		dev.Fd(),
		uapi.EVIOCSREP(),
		&settings,
		"failed to set repeat settings of evdev device",
	)
}

// Scancode retrieves the keycode for a scancode on the evdev device.
// The codes array holds two values:
// uint32[0] = Scancode to look up.
// uint32[1] = Populated keycode for the given scancode.
func (dev *Device) Scancode(codes [2]uint32) ([2]uint32, error) {
	return ioctlGetAny(
		dev.Fd(),
		uapi.EVIOCGKEYCODE(),
		&codes,
		"failed to get keycode of evdev device",
	)
}

// ScancodeV2 retrieves the keycode for a given keymap entry on this
// evdev device. The keymap parameter specifies the input index to query
// in its Index field and is updated with the retrieved keycode in its
// Keycode field.
func (dev *Device) ScancodeV2(keymap uapi.InputKeymapEntry) (uapi.InputKeymapEntry, error) {
	return ioctlGetAny(
		dev.Fd(),
		uapi.EVIOCGKEYCODE_V2(),
		&keymap,
		"failed to get keycodeV2 of evdev device",
	)
}

// SetScancode sets a keycode mapping for a scancode on the evdev device.
// The codes array holds two values:
// uint32[0] = Scancode to map.
// uint32[1] = Keycode to assign to that scancode.
func (dev *Device) SetScancode(codes [2]uint32) error {
	return ioctlSetAny(
		dev.Fd(),
		uapi.EVIOCSKEYCODE(),
		&codes,
		"failed to set keycode of evdev device",
	)
}

// SetScancodeV2 sets the keycode for a given keymap entry on the evdev
// device. The keymap parameter specifies the input index to map in its
// Index field and carries the keycode to assign in its Keycode field.
func (dev *Device) SetScancodeV2(keymap uapi.InputKeymapEntry) error {
	return ioctlSetAny(
		dev.Fd(),
		uapi.EVIOCSKEYCODE_V2(),
		&keymap,
		"failed to set index keycodeV2 of evdev device",
	)
}

// Name returns the evdev device’s name.
func (dev *Device) Name() (string, error) {
	return ioctlGetStr(
		dev.Fd(),
		uapi.EVIOCGNAME,
		"failed to get evdev device name",
	)
}

// PhysicalLocation returns the evdev device’s physical location.
// evdev.
func (dev *Device) PhysicalLocation() (string, error) {
	return ioctlGetStr(
		dev.Fd(),
		uapi.EVIOCGPHYS,
		"failed to get evdev device physical location",
	)
}

// UniqueID returns the evdev device’s unique identifier.
func (dev *Device) UniqueID() (string, error) {
	return ioctlGetStr(
		dev.Fd(),
		uapi.EVIOCGUNIQ,
		"failed to get evdev device unique id",
	)
}

// Properties returns the evdev device's input properties.
func (dev *Device) Properties() ([]uapi.InputCoder, error) {
	return ioctlGetBitmask(
		dev.Fd(),
		uapi.EVIOCGPROP,
		uapi.INPUT_PROP_CNT,
		"failed to get evdev device properties",
	)
}

// MTSlots retrieves the current multitouch slot values for a given
// [uapi.ABS_MT_SLOT] code.
func (dev *Device) MTSlots(code uapi.InputCoder) ([]int32, error) {
	var (
		abs     uapi.InputAbsoluteCode
		absInfo uapi.InputAbsInfo
		length  uint
		values  []int32
		ok      bool
		err     error
	)

	abs, ok = code.(uapi.InputAbsoluteCode)
	if !ok {
		return nil, fmt.Errorf("code %d (%s) is not an absolute code", code.Value(), code)
	}

	if !isMT(abs) {
		return nil, fmt.Errorf("code %d (%s): %w", abs, abs, ErrNotMultitouch)
	}

	absInfo, err = dev.AbsInfo(uapi.ABS_MT_SLOT)
	if err != nil {
		return nil, fmt.Errorf("failed to get evdev device multitouch slots: %w", err)
	}

	length = uint(absInfo.Maximum + 2)
	values = make([]int32, length)
	values[0] = int32(abs)

	_, err = ioctlGetAny(
		dev.Fd(),
		uapi.EVIOCGMTSLOTS(length*uint(unsafe.Sizeof(int32(0)))),
		&values[0],
		"failed to get evdev device multitouch slots",
	)
	if err != nil {
		return nil, err
	}

	return values[1:], nil
}

// EnabledKeycodes returns the evdev device's enabled key codes.
func (dev *Device) EnabledKeycodes() ([]uapi.InputCoder, error) {
	return ioctlGetBitmask(
		dev.Fd(),
		uapi.EVIOCGKEY,
		uapi.KEY_CNT,
		"failed to get evdev device enabled keycodes",
	)
}

// EnabledLEDs returns the evdev device's enabled LED codes.
func (dev *Device) EnabledLEDs() ([]uapi.InputCoder, error) {
	return ioctlGetBitmask(
		dev.Fd(),
		uapi.EVIOCGLED,
		uapi.LED_CNT,
		"failed to get evdev device enabled LEDs",
	)
}

// EnabledSounds returns the evdev device's enabled sound codes.
func (dev *Device) EnabledSounds() ([]uapi.InputCoder, error) {
	return ioctlGetBitmask(
		dev.Fd(),
		uapi.EVIOCGSND,
		uapi.SND_CNT,
		"failed to get evdev device enabled sounds",
	)
}

// EnabledSwitches returns the evdev device's enabled switch codes.
func (dev *Device) EnabledSwitches() ([]uapi.InputCoder, error) {
	return ioctlGetBitmask(
		dev.Fd(),
		uapi.EVIOCGSW,
		uapi.SW_CNT,
		"failed to get evdev device enabled switches",
	)
}

// Codes returns the list of [uapi.InputCoder] values supported by this
// device for the given [uapi.InputEventCode]. It dispatches to the
// appropriate helper method based on the event type.
func (dev *Device) Codes(event uapi.InputEventCode) ([]uapi.InputCoder, error) {
	switch event {
	case uapi.EV_SYN:
		return dev.SyncCodes()
	case uapi.EV_KEY:
		return dev.Keycodes()
	case uapi.EV_REL:
		return dev.RelativeCodes()
	case uapi.EV_ABS:
		return dev.AbsoluteCodes()
	case uapi.EV_MSC:
		return dev.MiscCodes()
	case uapi.EV_SW:
		return dev.SwitchCodes()
	case uapi.EV_LED:
		return dev.LEDCodes()
	case uapi.EV_SND:
		return dev.SoundCodes()
	case uapi.EV_REP:
		return dev.RepeatCodes()
	case uapi.EV_FF:
		return dev.FFCodes()
	case uapi.EV_PWR:
		return dev.PowerCodes()
	case uapi.EV_FF_STATUS:
		return dev.FFStatusCodes()
	default:
		return nil, fmt.Errorf("event %d (%s): %w", event, event, ErrUnsupportedEvent)
	}
}

// EventCodes returns the evdev device's supported event codes.
func (dev *Device) EventCodes() ([]uapi.InputCoder, error) {
	return ioctlGetBitmask(
		dev.Fd(),
		bitmaskReq(0),
		uapi.EV_CNT,
		"failed to get evdev device supported event codes",
	)
}

// SyncCodes returns the evdev device's supported sync codes.
func (dev *Device) SyncCodes() ([]uapi.InputCoder, error) {
	return ioctlGetBitmask(
		dev.Fd(),
		bitmaskReq(uapi.EV_SYN),
		uapi.SYN_CNT,
		"failed to get evdev device supported sync codes",
	)
}

// Keycodes returns the evdev device's supported keycodes.
func (dev *Device) Keycodes() ([]uapi.InputCoder, error) {
	return ioctlGetBitmask(
		dev.Fd(),
		bitmaskReq(uapi.EV_KEY),
		uapi.KEY_CNT,
		"failed to get evdev device supported keycodes",
	)
}

// RelativeCodes returns the evdev device's supported relative codes.
func (dev *Device) RelativeCodes() ([]uapi.InputCoder, error) {
	return ioctlGetBitmask(
		dev.Fd(),
		bitmaskReq(uapi.EV_REL),
		uapi.REL_CNT,
		"failed to get evdev device supported relative codes",
	)
}

// AbsoluteCodes returns the evdev device's supported absolute codes.
func (dev *Device) AbsoluteCodes() ([]uapi.InputCoder, error) {
	return ioctlGetBitmask(
		dev.Fd(),
		bitmaskReq(uapi.EV_ABS),
		uapi.ABS_CNT,
		"failed to get evdev device supported absolute codes",
	)
}

// MiscCodes returns the evdev device's supported misc codes.
func (dev *Device) MiscCodes() ([]uapi.InputCoder, error) {
	return ioctlGetBitmask(
		dev.Fd(),
		bitmaskReq(uapi.EV_MSC),
		uapi.MSC_CNT,
		"failed to get evdev device supported misc codes",
	)
}

// SwitchCodes returns the evdev device's supported switch codes.
func (dev *Device) SwitchCodes() ([]uapi.InputCoder, error) {
	return ioctlGetBitmask(
		dev.Fd(),
		bitmaskReq(uapi.EV_SW),
		uapi.SW_CNT,
		"failed to get evdev device supported switch codes",
	)
}

// LEDCodes returns the evdev device's supported LED codes.
func (dev *Device) LEDCodes() ([]uapi.InputCoder, error) {
	return ioctlGetBitmask(
		dev.Fd(),
		bitmaskReq(uapi.EV_LED),
		uapi.LED_CNT,
		"failed to get evdev device supported LED codes",
	)
}

// SoundCodes returns the evdev device's supported sound codes.
func (dev *Device) SoundCodes() ([]uapi.InputCoder, error) {
	return ioctlGetBitmask(
		dev.Fd(),
		bitmaskReq(uapi.EV_SND),
		uapi.SND_CNT,
		"failed to get evdev device supported sound codes",
	)
}

// RepeatCodes returns the evdev device's supported repeat codes.
func (dev *Device) RepeatCodes() ([]uapi.InputCoder, error) {
	var err error

	_, err = dev.Repeat()
	if err != nil {
		return nil, fmt.Errorf("failed to get evdev device supported repeat codes: %w", err)
	}

	return []uapi.InputCoder{uapi.REP_DELAY, uapi.REP_PERIOD}, nil
}

// FFCodes returns the evdev device's supported force-feedback codes.
func (dev *Device) FFCodes() ([]uapi.InputCoder, error) {
	return ioctlGetBitmask(
		dev.Fd(),
		bitmaskReq(uapi.EV_FF),
		uapi.FF_CNT,
		"failed to get evdev device supported force-feedback codes",
	)
}

// PowerCodes returns the evdev device's supported power codes.
func (dev *Device) PowerCodes() ([]uapi.InputCoder, error) {
	return ioctlGetBitmask(
		dev.Fd(),
		bitmaskReq(uapi.EV_PWR),
		uapi.KEY_CNT,
		"failed to get evdev device supported power codes",
	)
}

// FFStatusCodes returns the evdev device's supported force-feedback status
// codes.
func (dev *Device) FFStatusCodes() ([]uapi.InputCoder, error) {
	return ioctlGetBitmask(
		dev.Fd(),
		bitmaskReq(uapi.EV_FF_STATUS),
		uapi.FF_STATUS_MAX,
		"failed to get evdev device supported force-feedback status codes",
	)
}

// AbsInfo returns the evdev device's absolute axis information
// corresponding to the provided [uapi.InputAbsoluteCode].
func (dev *Device) AbsInfo(code uapi.InputCoder) (uapi.InputAbsInfo, error) {
	return ioctlGetAny(
		dev.Fd(),
		uapi.EVIOCGABS(uint(code.Value())),
		new(uapi.InputAbsInfo),
		"failed to get evdev device absolute axis info",
	)
}

// SetAbsInfo sets the evdev device's absolute axis information
// corresponding to the provided [uapi.InputAbsoluteCode].
func (dev *Device) SetAbsInfo(abs uapi.InputAbsoluteCode, absInfo uapi.InputAbsInfo) error {
	return ioctlSetAny(
		dev.Fd(),
		uapi.EVIOCGABS(uint(abs)),
		&absInfo,
		"failed to get evdev device absolute axis info",
	)
}

// SendFF sends a force-feedback effect to the evdev device.
func (dev *Device) SendFF(effect uapi.FFEffect) error {
	return ioctlSetAny(
		dev.Fd(),
		uapi.EVIOCSFF(),
		&effect,
		"failed to send evdev device force-feedback",
	)
}

// RemoveFF removes a force-feedback effect of the evdev device.
func (dev *Device) RemoveFF(id int32) error {
	return ioctlSetAny(
		dev.Fd(),
		uapi.EVIOCRMFF(),
		&id,
		"failed to remove evdev device force-feedback",
	)
}

// FFEffects returns the amount of force-feedback effects of the evdev device.
func (dev *Device) FFEffects() (int32, error) {
	return ioctlGetAny(
		dev.Fd(),
		uapi.EVIOCGEFFECTS(),
		new(int32),
		"failed to get evdev device force-feedback effects",
	)
}

// Grab toggles exclusive access to the evdev device.
// Pass 1 to grab, 0 to release.
func (dev *Device) Grab(grab int32) error {
	return ioctlSetAny(
		dev.Fd(),
		uapi.EVIOCGRAB(),
		&grab,
		"failed to grab/release evdev device",
	)
}

// Release toggles exclusive access to the evdev device.
// Pass 1 to release, 0 to grab.
func (dev *Device) Release(release int32) error {
	return ioctlSetAny(
		dev.Fd(),
		uapi.EVIOCREVOKE(),
		&release,
		"failed to release/grab evdev device",
	)
}

// SetEventMask configures which input events the device will report by
// applying the given event mask. The mask is a bitmask of event types and
// codes defined in [uapi.InputMask].
func (dev *Device) SetEventMask(mask uapi.InputMask) error {
	return ioctlSetAny(
		dev.Fd(),
		uapi.EVIOCSMASK(),
		&mask,
		"failed to set evdev device event mask",
	)
}

// SetClockID sets the clock source used by the kernel when timestamping
// events read from the device. The clockID must be one of the standard
// clock constants.
func (dev *Device) SetClockID(clockID int32) error {
	return ioctlSetAny(
		dev.Fd(),
		uapi.EVIOCSCLOCKID(),
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

func ioctlGetAny[T any](fd uintptr, req uint, arg *T, errMsg string) (T, error) {
	var errno syscall.Errno

	_, _, errno = unix.Syscall(
		unix.SYS_IOCTL,
		fd,
		uintptr(req),
		uintptr(unsafe.Pointer(arg)),
	)
	if errno != 0 {
		return *new(T), fmt.Errorf("%s: failed ioctl syscall: %w", errMsg, errno)
	}

	return *arg, nil
}

func ioctlSetAny[T any](fd uintptr, req uint, arg *T, errMsg string) error {
	var err error

	_, err = ioctlGetAny(fd, req, arg, errMsg)
	if err != nil {
		return err
	}

	return nil
}

func ioctlGetStr(fd uintptr, req func(uint) uint, errMsg string) (string, error) {
	var (
		buf []byte
		err error
	)

	buf = make([]byte, 256)

	_, err = ioctlGetAny(fd, req(256), &buf[0], errMsg)
	if err != nil {
		return "", err
	}

	return unix.ByteSliceToString(buf), nil
}

func ioctlGetBitmask[T uapi.InputCode](
	fd uintptr,
	req func(uint) uint,
	count T,
	errMsg string,
) ([]uapi.InputCoder, error) {
	var (
		buf   []byte
		codes []uapi.InputCoder
		code  T
		err   error
	)

	buf = bits.Bytes(int(count))

	_, err = ioctlGetAny(fd, req(uint(len(buf))), &buf[0], errMsg)
	if err != nil {
		return nil, err
	}

	codes = make([]uapi.InputCoder, 0, count)

	for code = range count {
		if !bits.Test(buf, int(code)) {
			continue
		}

		codes = append(codes, uapi.InputCoder(code))
	}

	return codes, nil
}

func bitmaskReq(event uapi.InputEventCode) func(uint) uint {
	return func(length uint) uint {
		return uapi.EVIOCGBIT(uint(event), length)
	}
}

func isMT(abs uapi.InputAbsoluteCode) bool {
	switch abs {
	case uapi.ABS_MT_TOUCH_MAJOR,
		uapi.ABS_MT_TOUCH_MINOR,
		uapi.ABS_MT_WIDTH_MAJOR,
		uapi.ABS_MT_WIDTH_MINOR,
		uapi.ABS_MT_ORIENTATION,
		uapi.ABS_MT_POSITION_X,
		uapi.ABS_MT_POSITION_Y,
		uapi.ABS_MT_TOOL_TYPE,
		uapi.ABS_MT_BLOB_ID,
		uapi.ABS_MT_TRACKING_ID,
		uapi.ABS_MT_PRESSURE,
		uapi.ABS_MT_DISTANCE,
		uapi.ABS_MT_TOOL_X,
		uapi.ABS_MT_TOOL_Y:
		return true
	default:
		return false
	}
}
