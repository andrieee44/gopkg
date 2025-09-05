// Package evdev provides a pure-Go interface to the Linux evdev (event
// device) subsystem. It lets you open input devices (keyboards, mice,
// gamepads, etc.), read raw input events, and interpret them as high-level
// actions.
package evdev

import (
	"encoding/binary"
	"errors"
	"fmt"
	"io"
	"math"
	"os"
	"path/filepath"
	"unsafe"

	"github.com/andrieee44/gopkg/linux/internal/inputwrap"
	"github.com/andrieee44/gopkg/linux/internal/ioctlwrap"
	"github.com/andrieee44/gopkg/linux/uapi/input"
	"github.com/andrieee44/gopkg/linux/uapi/ioctl"
)

// ErrNotMultiTouch is returned when the provided absolute event code
// does not correspond to any multitouch axis or slot index.
var ErrNotMultiTouch error = errors.New("is not a multitouch code")

// Device represents an evdev device.
// It wraps the opened /dev/input/eventN file.
type Device struct {
	file       *os.File
	eventsChan chan input.Event
	errChan    chan error
}

// NewDevice opens the evdev device at the given path and returns a [Device].
// The device file is opened in read-write mode. The caller is responsible
// for releasing resources by calling [Device.Close] when the device is no
// longer needed.
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
// containing the result of each attempt. The function does not stop at
// the first failure, it continues processing all paths and aggregates
// any errors into the errs slice.
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

			continue
		}

		devices = append(devices, device)
	}

	return devices, errs
}

// Filename returns the name of the underlying file.
func (dev *Device) Filename() string {
	return dev.file.Name()
}

// Fd returns the evdev device's underlying file descriptor.
func (dev *Device) Fd() uintptr {
	return dev.file.Fd()
}

// Version returns the evdev device's driver version.
func (dev *Device) Version() (int32, error) {
	return ioctlwrap.GetAny(
		dev.file,
		input.EVIOCGVERSION,
		new(int32),
		"failed to get event device driver version",
	)
}

// ID returns the evdev device’s identifier.
func (dev *Device) ID() (input.ID, error) {
	return ioctlwrap.GetAny(
		dev.file,
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
	return ioctlwrap.GetAny(
		dev.file,
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
	return ioctlwrap.SetAny(
		dev.file,
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
	return ioctlwrap.GetAny(
		dev.file,
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
	return ioctlwrap.GetAny(
		dev.file,
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
	return ioctlwrap.SetAny(
		dev.file,
		input.EVIOCSKEYCODE,
		&codes,
		"failed to set keycode of evdev device",
	)
}

// SetScancodeV2 sets the keycode for a given keymap entry on the evdev
// device. The keymap parameter specifies the input index to map in its
// Index field and carries the keycode to assign in its Keycode field.
func (dev *Device) SetScancodeV2(keymap input.KeymapEntry) error {
	return ioctlwrap.SetAny(
		dev.file,
		input.EVIOCSKEYCODE_V2,
		&keymap,
		"failed to set index keycodeV2 of evdev device",
	)
}

// Name returns the evdev device’s name as a string.
// bufSize specifies the maximum number of bytes to read. If unsure,
// use 256.
func (dev *Device) Name(bufSize uint32) (string, error) {
	return ioctlwrap.GetStr(
		dev.file,
		input.EVIOCGNAME,
		bufSize,
		"failed to get evdev device name",
	)
}

// PhysicalLocation returns the device’s physical location as a string.
// bufSize specifies the maximum number of bytes to read. If unsure,
// use 256.
func (dev *Device) PhysicalLocation(bufSize uint32) (string, error) {
	return ioctlwrap.GetStr(
		dev.file,
		input.EVIOCGPHYS,
		bufSize,
		"failed to get evdev device physical location",
	)
}

// UniqueID returns the device’s unique identifier as a string.
// bufSize specifies the maximum number of bytes to read. If unsure,
// use 256.
func (dev *Device) UniqueID(bufSize uint32) (string, error) {
	return ioctlwrap.GetStr(
		dev.file,
		input.EVIOCGUNIQ,
		bufSize,
		"failed to get evdev device unique id",
	)
}

// Properties returns the evdev device's input properties.
func (dev *Device) Properties() ([]input.PropCode, error) {
	return inputwrap.GetBitmask(
		dev.file,
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
		return nil, fmt.Errorf("%s: code %s: %w", dev.Filename(), abs.Pretty(), ErrNotMultiTouch)
	}

	absInfo, err = dev.AbsInfo(input.ABS_MT_SLOT)
	if err != nil {
		return nil, fmt.Errorf("failed to get evdev device multitouch slots: %w", err)
	}

	length = uint32(absInfo.Maximum) + 2
	values = make([]int32, length)
	values[0] = int32(abs)

	_, err = ioctlwrap.GetAny(
		dev.file,
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
	return inputwrap.GetBitmask(
		dev.file,
		input.EVIOCGKEY,
		input.KEY_CNT,
		"failed to get evdev device enabled keycodes",
	)
}

// EnabledLEDs returns the evdev device's enabled LED codes.
func (dev *Device) EnabledLEDs() ([]input.LEDCode, error) {
	return inputwrap.GetBitmask(
		dev.file,
		input.EVIOCGLED,
		input.LED_CNT,
		"failed to get evdev device enabled LEDs",
	)
}

// EnabledSounds returns the evdev device's enabled sound codes.
func (dev *Device) EnabledSounds() ([]input.SoundCode, error) {
	return inputwrap.GetBitmask(
		dev.file,
		input.EVIOCGSND,
		input.SND_CNT,
		"failed to get evdev device enabled sounds",
	)
}

// EnabledSwitches returns the evdev device's enabled switch codes.
func (dev *Device) EnabledSwitches() ([]input.SwitchCode, error) {
	return inputwrap.GetBitmask(
		dev.file,
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
		return inputwrap.AsInputCoders(dev.EnabledKeycodes)
	case input.EV_SW:
		return inputwrap.AsInputCoders(dev.EnabledSwitches)
	case input.EV_LED:
		return inputwrap.AsInputCoders(dev.EnabledLEDs)
	case input.EV_SND:
		return inputwrap.AsInputCoders(dev.EnabledSounds)
	default:
		return nil, fmt.Errorf(
			"%s: event %s: %w",
			dev.Filename(),
			event.Pretty(),
			input.ErrUnsupportedEvent,
		)
	}
}

// Events returns the evdev device's supported event codes.
func (dev *Device) Events() ([]input.EventCode, error) {
	return inputwrap.GetBitmask(
		dev.file,
		input.BitmaskReq(0),
		input.EV_CNT,
		"failed to get evdev device supported event codes",
	)
}

// Syncs returns the evdev device's supported sync codes.
func (dev *Device) Syncs() ([]input.SyncCode, error) {
	return inputwrap.GetBitmask(
		dev.file,
		input.BitmaskReq(input.EV_SYN),
		input.SYN_CNT,
		"failed to get evdev device supported sync codes",
	)
}

// Keys returns the evdev device's supported keycodes.
func (dev *Device) Keys() ([]input.KeyCode, error) {
	return inputwrap.GetBitmask(
		dev.file,
		input.BitmaskReq(input.EV_KEY),
		input.KEY_CNT,
		"failed to get evdev device supported keycodes",
	)
}

// Relatives returns the evdev device's supported relative codes.
func (dev *Device) Relatives() ([]input.RelativeCode, error) {
	return inputwrap.GetBitmask(
		dev.file,
		input.BitmaskReq(input.EV_REL),
		input.REL_CNT,
		"failed to get evdev device supported relative codes",
	)
}

// Absolutes returns the evdev device's supported absolute codes.
func (dev *Device) Absolutes() ([]input.AbsoluteCode, error) {
	return inputwrap.GetBitmask(
		dev.file,
		input.BitmaskReq(input.EV_ABS),
		input.ABS_CNT,
		"failed to get evdev device supported absolute codes",
	)
}

// Miscs returns the evdev device's supported misc codes.
func (dev *Device) Miscs() ([]input.MiscCode, error) {
	return inputwrap.GetBitmask(
		dev.file,
		input.BitmaskReq(input.EV_MSC),
		input.MSC_CNT,
		"failed to get evdev device supported misc codes",
	)
}

// Switches returns the evdev device's supported switch codes.
func (dev *Device) Switches() ([]input.SwitchCode, error) {
	return inputwrap.GetBitmask(
		dev.file,
		input.BitmaskReq(input.EV_SW),
		input.SW_CNT,
		"failed to get evdev device supported switch codes",
	)
}

// LEDs returns the evdev device's supported LED codes.
func (dev *Device) LEDs() ([]input.LEDCode, error) {
	return inputwrap.GetBitmask(
		dev.file,
		input.BitmaskReq(input.EV_LED),
		input.LED_CNT,
		"failed to get evdev device supported LED codes",
	)
}

// Sounds returns the evdev device's supported sound codes.
func (dev *Device) Sounds() ([]input.SoundCode, error) {
	return inputwrap.GetBitmask(
		dev.file,
		input.BitmaskReq(input.EV_SND),
		input.SND_CNT,
		"failed to get evdev device supported sound codes",
	)
}

// Repeats returns the evdev device's supported repeat codes.
func (dev *Device) Repeats() ([]input.RepeatCode, error) {
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

// ForceFeedbacks returns the evdev device's supported force-feedback codes.
func (dev *Device) ForceFeedbacks() ([]input.FFCode, error) {
	return inputwrap.GetBitmask(
		dev.file,
		input.BitmaskReq(input.EV_FF),
		input.FF_CNT,
		"failed to get evdev device supported force-feedback codes",
	)
}

// Powers returns the evdev device's supported power codes.
func (dev *Device) Powers() ([]input.KeyCode, error) {
	return inputwrap.GetBitmask(
		dev.file,
		input.BitmaskReq(input.EV_PWR),
		input.KEY_CNT,
		"failed to get evdev device supported power codes",
	)
}

// FFStatuses returns the evdev device's supported force-feedback status
// codes.
func (dev *Device) FFStatuses() ([]input.FFStatusCode, error) {
	return inputwrap.GetBitmask(
		dev.file,
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
// such as [Device.Keys] or [Device.Relatives] when typed
// codes are required.
func (dev *Device) Codes(event input.EventCode) ([]input.Coder, error) {
	switch event {
	case input.EV_SYN:
		return inputwrap.AsInputCoders(dev.Syncs)
	case input.EV_KEY:
		return inputwrap.AsInputCoders(dev.Keys)
	case input.EV_REL:
		return inputwrap.AsInputCoders(dev.Relatives)
	case input.EV_ABS:
		return inputwrap.AsInputCoders(dev.Absolutes)
	case input.EV_MSC:
		return inputwrap.AsInputCoders(dev.Miscs)
	case input.EV_SW:
		return inputwrap.AsInputCoders(dev.Switches)
	case input.EV_LED:
		return inputwrap.AsInputCoders(dev.LEDs)
	case input.EV_SND:
		return inputwrap.AsInputCoders(dev.Sounds)
	case input.EV_REP:
		return inputwrap.AsInputCoders(dev.Repeats)
	case input.EV_FF:
		return inputwrap.AsInputCoders(dev.ForceFeedbacks)
	case input.EV_PWR:
		return inputwrap.AsInputCoders(dev.Powers)
	case input.EV_FF_STATUS:
		return inputwrap.AsInputCoders(dev.FFStatuses)
	default:
		return nil, fmt.Errorf(
			"%s: event %s: %w",
			dev.Filename(),
			event.Pretty(),
			input.ErrUnsupportedEvent,
		)
	}
}

// AbsInfo returns the evdev device's absolute axis information
// corresponding to the provided [input.AbsoluteCode].
func (dev *Device) AbsInfo(abs input.AbsoluteCode) (input.AbsInfo, error) {
	return ioctlwrap.GetAny(
		dev.file,
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
	return ioctlwrap.SetAny(
		dev.file,
		func() (uint32, error) {
			return input.EVIOCSABS(abs)
		},
		&absInfo,
		"failed to set evdev device absolute axis info",
	)
}

// SendFF uploads a force feedback effect to the device's internal buffer.
// To request a new effect ID, set the ID field of [input.FFEffect] to -1
// before calling this method. This does not play the effect; use [PlayFF]
// to trigger playback.
func (dev *Device) SendFF(effect input.FFEffect) error {
	return ioctlwrap.SetAny(
		dev.file,
		input.EVIOCSFF,
		&effect,
		"failed to upload force feedback to evdev device buffer",
	)
}

// RemoveFF deletes a force feedback effect from the device's internal buffer.
// The effect is identified by its ID. Once removed, the effect cannot be
// played again unless re-uploaded using [SendFF].
func (dev *Device) RemoveFF(id int32) error {
	return ioctlwrap.SetAny(
		dev.file,
		input.EVIOCRMFF,
		&id,
		"failed to remove force feedback in evdev device buffer",
	)
}

// FFEffects reports how many force feedback effects the device can store.
// This represents the size of the device's internal buffer for force
// feedback effects, not how many are currently loaded.
func (dev *Device) FFEffects() (int32, error) {
	return ioctlwrap.GetAny(
		dev.file,
		input.EVIOCGEFFECTS,
		new(int32),
		"failed to get evdev device force feedback buffer size",
	)
}

// Grab toggles exclusive access to the evdev device.
// Pass 1 to grab, 0 to release.
func (dev *Device) Grab(grab int32) error {
	return ioctlwrap.SetAny(
		dev.file,
		input.EVIOCGRAB,
		&grab,
		"failed to grab/release evdev device",
	)
}

// Release toggles exclusive access to the evdev device.
// Pass 1 to release, 0 to grab.
func (dev *Device) Release(release int32) error {
	return ioctlwrap.SetAny(
		dev.file,
		input.EVIOCREVOKE,
		&release,
		"failed to release/grab evdev device",
	)
}

// SetEventMask configures which input events the device will report by
// applying the given event mask. The mask is a bitmask of event types and
// codes defined in [input.Mask].
func (dev *Device) SetEventMask(mask input.Mask) error {
	return ioctlwrap.SetAny(
		dev.file,
		input.EVIOCSMASK,
		&mask,
		"failed to set evdev device event mask",
	)
}

// SetClockID sets the clock source used by the kernel when timestamping
// events read from the device. The clockID must be one of the standard
// clock constants such as [sys/unix.CLOCK_MONOTONIC].
func (dev *Device) SetClockID(clockID int32) error {
	return ioctlwrap.SetAny(
		dev.file,
		input.EVIOCSCLOCKID,
		&clockID,
		"failed to set evdev device clock id",
	)
}

// Snapshot returns a point-in-time capture of dev's current state and
// capabilities. The returned Snapshot reflects the device at the moment
// of the call, including identifiers, enabled and supported events,
// repeat settings, absolute axis details, multi-touch information, and
// descriptor fields such as Name, Filename, and Version. If the snapshot
// cannot be created, Snapshot returns a non-nil error.
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

// ReadEvents returns two read-only channels for receiving input events
// and errors.
// The first channel delivers [input.Event] values generated by the device.
// The second channel reports errors encountered during event processing.
func (dev *Device) ReadEvents() (<-chan input.Event, <-chan error) {
	dev.eventsChan = make(chan input.Event)
	dev.errChan = make(chan error)

	go dev.serve()

	return dev.eventsChan, dev.errChan
}

// PlayFF triggers playback of a force feedback effect previously uploaded.
// The effect is identified by its ID. The value specifies how many times the
// effect should repeat. Note that while effect IDs are handled as int32 when
// uploading or removing, only the lower 16 bits are used when playing—due to
// the input_event.code field being a uint16. IDs outside the 0–65535 range
// cannot be played and will return an error.
func (dev *Device) PlayFF(id, value int32) error {
	var err error

	if id < 0 || id > math.MaxUint16 {
		return fmt.Errorf(
			"%s: force feedback id %d exceeds uint16 limit: failed to play force feedback: %w",
			dev.Filename(),
			id,
			ioctl.ErrSizeOverflow,
		)
	}

	err = binary.Write(dev.file, binary.NativeEndian, input.Event{
		Type:  input.EV_FF,
		Code:  uint16(id),
		Value: value,
	})
	if err != nil {
		return fmt.Errorf("%s: failed to play force feedback: %w", dev.Filename(), err)
	}

	return nil
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

func (dev *Device) serve() {
	var (
		event input.Event
		err   error
	)

	for {
		err = binary.Read(dev.file, binary.NativeEndian, &event)
		if err != nil {
			if errors.Is(err, io.EOF) {
				break
			}

			dev.errChan <- err

			break
		}

		dev.eventsChan <- event
	}

	close(dev.eventsChan)
	close(dev.errChan)
}
