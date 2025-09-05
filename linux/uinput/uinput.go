// Package uinput provides a pure-Go interface to the Linux uinput
// subsystem. It lets you create virtual input devices, emit events, and
// handle force feedback (FF) upload and erase requests from the kernel.
package uinput

import (
	"encoding/binary"
	"errors"
	"fmt"
	"io"
	"os"

	"github.com/andrieee44/gopkg/linux/internal/ioctlwrap"
	"github.com/andrieee44/gopkg/linux/uapi/input"
	"github.com/andrieee44/gopkg/linux/uapi/uinput"
)

// FFUploadEvent represents a force feedback upload request from the kernel.
// FF contains the [uinput.FFUpload] data to be stored, and Ret is a
// one‑shot channel the handler must send the updated [uinput.FFUpload] back
// on (with [uinput.FFUpload].Retval set).
type FFUploadEvent struct {
	FF  uinput.FFUpload
	Ret chan<- uinput.FFUpload
}

// FFEraseEvent represents a force feedback erase request from the kernel.
// [FF] contains the [uinput.FFErase] identifying the effect to remove, and
// [Ret] is a one‑shot channel the handler must send the updated
// [uinput.FFErase] back on (with [uinput.FFUpload].Retval set).
type FFEraseEvent struct {
	FF  uinput.FFErase
	Ret chan<- uinput.FFErase
}

// Device represents a virtual input device created through uinput. Use
// [NewDevice] to construct one, configure its capabilities with the Set*
// methods, then call [Device.Create] to activate it. When finished, call
// [Device.Destroy] to remove it from the system.
type Device struct {
	setup      uinput.Setup
	file       *os.File
	uploadChan chan FFUploadEvent
	eraseChan  chan FFEraseEvent
	errChan    chan error
}

// ErrNameTooLong is returned by [NewDevice] when the provided name exceeds
// [uinput.UINPUT_MAX_NAME_SIZE] bytes including the null terminator.
var ErrNameTooLong error = errors.New("name is too long")

// NewDevice returns a new [Device] with the given [input.ID] and name,
// opening /dev/uinput for writing. The name must not exceed
// [uinput.UINPUT_MAX_NAME_SIZE] bytes including the null terminator,
// otherwise [ErrNameTooLong] is returned. Before activating the [Device]
// with [Device.Create], clients should configure its capabilities using
// the provided Set* methods such as [Device.SetKeys] or [Device.SetAbsCodes].
func NewDevice(id input.ID, name string) (*Device, error) {
	var (
		buf [uinput.UINPUT_MAX_NAME_SIZE]byte
		dev *Device
		err error
	)

	if len(name)+1 > uinput.UINPUT_MAX_NAME_SIZE {
		return nil, fmt.Errorf(
			"failed to allocate uinput device: name is %d bytes long, max is %d bytes including the null terminator: %w",
			len(name),
			uinput.UINPUT_MAX_NAME_SIZE,
			ErrNameTooLong,
		)
	}

	copy(buf[:], name)

	dev = &Device{
		setup: uinput.Setup{
			ID:   id,
			Name: buf,
		},
	}

	dev.file, err = os.OpenFile("/dev/uinput", os.O_WRONLY, 0)
	if err != nil {
		return nil, err
	}

	return dev, nil
}

// Fd returns the uinput device's underlying file descriptor.
func (dev *Device) Fd() uintptr {
	return dev.file.Fd()
}

// Create activates the [Device], making it available for use by the system.
// Call this after adding all desired capabilities with the Set* methods.
// Returns an error if activation fails.
func (dev *Device) Create() error {
	return ioctlwrap.SetAny(
		dev.file,
		uinput.UI_DEV_CREATE,
		&dev.setup,
		"failed to create uinput device",
	)
}

// Destroy deactivates the [Device] and removes it from the system. Once
// destroyed, the device can no longer send events. Returns an error if
// deactivation fails.
func (dev *Device) Destroy() error {
	return ioctlwrap.Empty(
		dev.file,
		uinput.UI_DEV_DESTROY,
		"failed to destroy uinput device",
	)
}

// Name returns the system‑assigned name of the [Device], typically something
// like "eventN" where N is the event device number.
func (dev *Device) Name() (string, error) {
	return ioctlwrap.GetStr(
		dev.file,
		uinput.UI_GET_SYSNAME,
		256,
		"failed to get name of uinput device",
	)
}

// Version reports the uinput protocol version as defined by the running
// kernel.
func (dev *Device) Version() (uint32, error) {
	return ioctlwrap.GetAny(
		dev.file,
		uinput.UI_GET_VERSION,
		new(uint32),
		"failed to get uinput protocol version",
	)
}

// SetAbsInfos configures the absolute axis parameters for the [Device].
// Each [uinput.AbsSetup] specifies the axis code and its value range,
// fuzz, and flat settings.
func (dev *Device) SetAbsInfos(absInfos []uinput.AbsSetup) error {
	var (
		absInfo uinput.AbsSetup
		err     error
	)

	for _, absInfo = range absInfos {
		err = ioctlwrap.SetAny(
			dev.file,
			uinput.UI_ABS_SETUP,
			&absInfo,
			"failed to set uinput device absolute infos",
		)
		if err != nil {
			return err
		}
	}

	return nil
}

// SetAbsCodes enables reporting of the given absolute axis codes on the
// [Device], allowing it to send events such as joystick positions or touch
// coordinates.
func (dev *Device) SetAbsCodes(codes []input.AbsoluteCode) error {
	return setCodes(
		dev.file,
		uinput.UI_SET_ABSBIT,
		codes,
		"failed to set uinput device absolute codes",
	)
}

// SetEvents enables reporting of the given high‑level event types on the
// [Device], such as key, relative, or absolute events.
func (dev *Device) SetEvents(codes []input.EventCode) error {
	return setCodes(
		dev.file,
		uinput.UI_SET_EVBIT,
		codes,
		"failed to set uinput device event codes",
	)
}

// SetForceFeedbacks enables the given force‑feedback effect codes on the
// [Device], allowing it to send haptic feedback events.
func (dev *Device) SetForceFeedbacks(codes []input.FFCode) error {
	return setCodes(
		dev.file,
		uinput.UI_SET_FFBIT,
		codes,
		"failed to set uinput device force feedback codes",
	)
}

// SetKeys enables reporting of the given key codes on the [Device], allowing
// it to send key press and release events.
func (dev *Device) SetKeys(codes []input.KeyCode) error {
	return setCodes(
		dev.file,
		uinput.UI_SET_KEYBIT,
		codes,
		"failed to set uinput device key codes",
	)
}

// SetLEDs enables control of the given LED codes on the [Device], such as
// keyboard lock lights or controller indicators.
func (dev *Device) SetLEDs(codes []input.LEDCode) error {
	return setCodes(
		dev.file,
		uinput.UI_SET_LEDBIT,
		codes,
		"failed to set uinput device LED codes",
	)
}

// SetMiscs enables reporting of the given miscellaneous codes on the
// [Device], used for less common input events.
func (dev *Device) SetMiscs(codes []input.MiscCode) error {
	return setCodes(
		dev.file,
		uinput.UI_SET_MSCBIT,
		codes,
		"failed to set uinput device misc codes",
	)
}

// SetProps enables the given input property codes on the [Device], describing
// its capabilities (e.g., whether it has a direct touch surface).
func (dev *Device) SetProps(codes []input.PropCode) error {
	return setCodes(
		dev.file,
		uinput.UI_SET_PROPBIT,
		codes,
		"failed to set uinput device property codes",
	)
}

// SetRelatives enables reporting of the given relative axis codes on the
// [Device], such as mouse movement deltas.
func (dev *Device) SetRelatives(codes []input.RelativeCode) error {
	return setCodes(
		dev.file,
		uinput.UI_SET_RELBIT,
		codes,
		"failed to set uinput device relative codes",
	)
}

// SetSounds enables playback of the given sound codes on the [Device], such
// as clicks or tones.
func (dev *Device) SetSounds(codes []input.SoundCode) error {
	return setCodes(
		dev.file,
		uinput.UI_SET_SNDBIT,
		codes,
		"failed to set uinput device sound codes",
	)
}

// SetSwitches enables reporting of the given switch codes on the [Device],
// such as lid open/close or tablet mode switches.
func (dev *Device) SetSwitches(codes []input.SwitchCode) error {
	return setCodes(
		dev.file,
		uinput.UI_SET_SWBIT,
		codes,
		"failed to set uinput device switch codes",
	)
}

// SetFFEffectsMax sets the maximum number of force feedback effects the
// device can store and starts its event loop. It returns three channels:
// one for upload requests (FFUploadEvent), one for erase requests
// (FFEraseEvent), and one for asynchronous errors. The caller should read
// from the upload and erase channels, handle each request, and send the
// updated struct back on the provided Ret channel. Errors from the device
// are sent on the error channel.
func (dev *Device) SetFFEffectsMax(
	ffEffectsMax uint32,
) (<-chan FFUploadEvent, <-chan FFEraseEvent, <-chan error) {
	dev.setup.FFEffectsMax = ffEffectsMax
	dev.uploadChan = make(chan FFUploadEvent)
	dev.eraseChan = make(chan FFEraseEvent)
	dev.errChan = make(chan error)

	go dev.serve()

	return dev.uploadChan, dev.eraseChan, dev.errChan
}

func (dev *Device) uploadFF(value int32) error {
	var (
		ff      uinput.FFUpload
		retChan chan uinput.FFUpload
		err     error
	)

	ff.RequestID = uint32(value)
	retChan = make(chan uinput.FFUpload)

	ff, err = ioctlwrap.GetAny(
		dev.file,
		uinput.UI_BEGIN_FF_UPLOAD,
		&ff,
		"failed to begin uinput device force feedback upload",
	)
	if err != nil {
		return err
	}

	dev.uploadChan <- FFUploadEvent{
		FF:  ff,
		Ret: retChan,
	}

	ff = <-retChan

	err = ioctlwrap.SetAny(
		dev.file,
		uinput.UI_END_FF_UPLOAD,
		&ff,
		"failed to end uinput device force feedback upload",
	)
	if err != nil {
		return err
	}

	return nil
}

func (dev *Device) eraseFF(value int32) error {
	var (
		ff      uinput.FFErase
		retChan chan uinput.FFErase
		err     error
	)

	ff.RequestID = uint32(value)
	retChan = make(chan uinput.FFErase)

	ff, err = ioctlwrap.GetAny(
		dev.file,
		uinput.UI_BEGIN_FF_ERASE,
		&ff,
		"failed to begin uinput device force feedback erase",
	)
	if err != nil {
		return err
	}

	dev.eraseChan <- FFEraseEvent{
		FF:  ff,
		Ret: retChan,
	}

	ff = <-retChan

	err = ioctlwrap.SetAny(
		dev.file,
		uinput.UI_END_FF_ERASE,
		&ff,
		"failed to end uinput device force feedback erase",
	)
	if err != nil {
		return err
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

		if event.Type != uinput.EV_UINPUT {
			continue
		}

		err = nil

		switch event.Code {
		case uinput.UI_FF_UPLOAD:
			err = dev.uploadFF(event.Value)
		case uinput.UI_FF_ERASE:
			err = dev.eraseFF(event.Value)
		}

		if err != nil {
			dev.errChan <- err

			break
		}
	}

	close(dev.uploadChan)
	close(dev.eraseChan)
	close(dev.errChan)
}

func setCodes[T input.Code](
	file *os.File,
	fn func() (uint32, error),
	codes []T,
	errMsg string,
) error {
	var (
		code T
		err  error
	)

	for _, code = range codes {
		err = ioctlwrap.SetAny(file, fn, &code, errMsg)
		if err != nil {
			return err
		}
	}

	return nil
}
