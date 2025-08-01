package uapi

import "fmt"

// InputCode represents any input event code as a generic constraint.
type InputCode interface {
	InputPropCode |
		InputEventCode |
		InputSyncCode |
		InputKeyCode |
		InputRelativeCode |
		InputAbsoluteCode |
		InputSwitchCode |
		InputMiscCode |
		InputLEDCode |
		InputRepeatCode |
		InputSoundCode |
		InputBusCode |
		InputMultiTouchCode |
		InputFFCode |
		InputFFStatusCode
}

// InputCoder represents any input event code.
type InputCoder interface {
	fmt.Stringer
	Value() uint32
}

// InputPropCode identifies a property supported by an input device.
//
//go:generate go run github.com/dmarkham/enumer -type=InputPropCode
type InputPropCode uint32

// InputEventCode identifies the broad event category for an input
// event.
//
//go:generate go run github.com/dmarkham/enumer -type=InputEventCode
type InputEventCode uint32

// InputSyncCode describes synchronization events that delimit
// packets of input data or change reporting mode.
//
//go:generate go run github.com/dmarkham/enumer -type=InputSyncCode
type InputSyncCode uint32

// InputKeyCode represents a keyboard or button key code.
//
//go:generate go run github.com/dmarkham/enumer -type=InputKeyCode
type InputKeyCode uint32

// InputRelativeCode describes relative axes, such as pointer
// movement, scroll wheels, and tilt sensors.
//
//go:generate go run github.com/dmarkham/enumer -type=InputRelativeCode
type InputRelativeCode uint32

// InputAbsoluteCode describes absolute axes, such as touch‐screen
// coordinates, joystick positions, or tablet pressure.
//
//go:generate go run github.com/dmarkham/enumer -type=InputAbsoluteCode
type InputAbsoluteCode uint32

// InputSwitchCode describes switch events, usually binary toggles
// like lid, tablet mode, or proximity sensors.
//
//go:generate go run github.com/dmarkham/enumer -type=InputSwitchCode
type InputSwitchCode uint32

// InputMiscCode covers miscellaneous event codes that don’t fit
// into other categories, such as drive insert/eject, auto-repeat
// toggle, or power events.
//
//go:generate go run github.com/dmarkham/enumer -type=InputMiscCode
type InputMiscCode uint32

// InputLEDCode represents status LEDs on a device, such as
// keyboard or system LEDs.
//
//go:generate go run github.com/dmarkham/enumer -type=InputLEDCode
type InputLEDCode uint32

// InputRepeatCode defines auto‐repeat settings for keys.
//
//go:generate go run github.com/dmarkham/enumer -type=InputRepeatCode
type InputRepeatCode uint32

// InputSoundCode describes simple tone and sound events,
// typically used for system beeps.
//
//go:generate go run github.com/dmarkham/enumer -type=InputSoundCode
type InputSoundCode uint32

// InputBusCode identifies the hardware bus (USB, PCI, Bluetooth, etc.)
//
//go:generate go run github.com/dmarkham/enumer -type=InputBusCode
type InputBusCode uint32

// InputMultiTouchCode represents multi-touch event codes
// (MT_SLOT, MT_POSITION_X, MT_TRACKING_ID, etc.)
//
//go:generate go run github.com/dmarkham/enumer -type=InputMultiTouchCode
type InputMultiTouchCode uint32

// InputFFCode denotes force-feedback effect types
// (FF_RUMBLE, FF_SPRING, FF_PERIODIC, etc.)
//
//go:generate go run github.com/dmarkham/enumer -type=InputFFCode
type InputFFCode uint32

// InputFFStatusCode holds the status value for an FF_STATUS event.
//
//go:generate go run github.com/dmarkham/enumer -type=InputFFStatusCode
type InputFFStatusCode uint32

// Value returns the uint32 numeric representation of the InputPropCode.
func (code InputPropCode) Value() uint32 {
	return uint32(code)
}

// Value returns the uint32 numeric representation of the InputEventCode.
func (code InputEventCode) Value() uint32 {
	return uint32(code)
}

// Value returns the uint32 numeric representation of the InputSyncCode.
func (code InputSyncCode) Value() uint32 {
	return uint32(code)
}

// Value returns the uint32 numeric representation of the InputKeyCode.
func (code InputKeyCode) Value() uint32 {
	return uint32(code)
}

// Value returns the uint32 numeric representation of the InputRelativeCode.
func (code InputRelativeCode) Value() uint32 {
	return uint32(code)
}

// Value returns the uint32 numeric representation of the InputAbsoluteCode.
func (code InputAbsoluteCode) Value() uint32 {
	return uint32(code)
}

// Value returns the uint32 numeric representation of the InputSwitchCode.
func (code InputSwitchCode) Value() uint32 {
	return uint32(code)
}

// Value returns the uint32 numeric representation of the InputMiscCode.
func (code InputMiscCode) Value() uint32 {
	return uint32(code)
}

// Value returns the uint32 numeric representation of the InputLEDCode.
func (code InputLEDCode) Value() uint32 {
	return uint32(code)
}

// Value returns the uint32 numeric representation of the InputRepeatCode.
func (code InputRepeatCode) Value() uint32 {
	return uint32(code)
}

// Value returns the uint32 numeric representation of the InputSoundCode.
func (code InputSoundCode) Value() uint32 {
	return uint32(code)
}

// Value returns the uint32 numeric representation of the InputBusCode.
func (code InputBusCode) Value() uint32 {
	return uint32(code)
}

// Value returns the uint32 numeric representation of the
// InputMultiTouchCode.
func (code InputMultiTouchCode) Value() uint32 {
	return uint32(code)
}

// Value returns the uint32 numeric representation of the InputFFCode.
func (code InputFFCode) Value() uint32 {
	return uint32(code)
}

// Value returns the uint32 numeric representation of the InputFFStatusCode.
func (code InputFFStatusCode) Value() uint32 {
	return uint32(code)
}
