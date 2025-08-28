package input

import "fmt"

// Code represents any input event code as a generic constraint.
type Code interface {
	PropCode |
		EventCode |
		SyncCode |
		KeyCode |
		RelativeCode |
		AbsoluteCode |
		SwitchCode |
		MiscCode |
		LEDCode |
		RepeatCode |
		SoundCode |
		BusCode |
		MultiTouchCode |
		FFCode |
		FFStatusCode
}

// Coder represents any input event code.
type Coder interface {
	fmt.Stringer
	Value() uint32
}

// PropCode identifies a property supported by an input device.
//
//go:generate go run github.com/dmarkham/enumer -type=PropCode
type PropCode uint32

// EventCode identifies the broad event category for an input
// event.
//
//go:generate go run github.com/dmarkham/enumer -type=EventCode
type EventCode uint32

// SyncCode describes synchronization events that delimit
// packets of input data or change reporting mode.
//
//go:generate go run github.com/dmarkham/enumer -type=SyncCode
type SyncCode uint32

// KeyCode represents a keyboard or button key code.
//
//go:generate go run github.com/dmarkham/enumer -type=KeyCode
type KeyCode uint32

// RelativeCode describes relative axes, such as pointer
// movement, scroll wheels, and tilt sensors.
//
//go:generate go run github.com/dmarkham/enumer -type=RelativeCode
type RelativeCode uint32

// AbsoluteCode describes absolute axes, such as touch‐screen
// coordinates, joystick positions, or tablet pressure.
//
//go:generate go run github.com/dmarkham/enumer -type=AbsoluteCode
type AbsoluteCode uint32

// SwitchCode describes switch events, usually binary toggles
// like lid, tablet mode, or proximity sensors.
//
//go:generate go run github.com/dmarkham/enumer -type=SwitchCode
type SwitchCode uint32

// MiscCode covers miscellaneous event codes that don’t fit
// into other categories, such as drive insert/eject, auto-repeat
// toggle, or power events.
//
//go:generate go run github.com/dmarkham/enumer -type=MiscCode
type MiscCode uint32

// LEDCode represents status LEDs on a device, such as
// keyboard or system LEDs.
//
//go:generate go run github.com/dmarkham/enumer -type=LEDCode
type LEDCode uint32

// RepeatCode defines auto‐repeat settings for keys.
//
//go:generate go run github.com/dmarkham/enumer -type=RepeatCode
type RepeatCode uint32

// SoundCode describes simple tone and sound events,
// typically used for system beeps.
//
//go:generate go run github.com/dmarkham/enumer -type=SoundCode
type SoundCode uint32

// BusCode identifies the hardware bus (USB, PCI, Bluetooth, etc.)
//
//go:generate go run github.com/dmarkham/enumer -type=BusCode
type BusCode uint32

// MultiTouchCode represents multi-touch event codes
// (MT_SLOT, MT_POSITION_X, MT_TRACKING_ID, etc.)
//
//go:generate go run github.com/dmarkham/enumer -type=MultiTouchCode
type MultiTouchCode uint32

// FFCode denotes force-feedback effect types
// (FF_RUMBLE, FF_SPRING, FF_PERIODIC, etc.)
//
//go:generate go run github.com/dmarkham/enumer -type=FFCode
type FFCode uint32

// FFStatusCode holds the status value for an FF_STATUS event.
//
//go:generate go run github.com/dmarkham/enumer -type=FFStatusCode
type FFStatusCode uint32

// Value returns the uint32 numeric representation of the PropCode.
func (code PropCode) Value() uint32 {
	return uint32(code)
}

// Value returns the uint32 numeric representation of the EventCode.
func (code EventCode) Value() uint32 {
	return uint32(code)
}

// Value returns the uint32 numeric representation of the SyncCode.
func (code SyncCode) Value() uint32 {
	return uint32(code)
}

// Value returns the uint32 numeric representation of the KeyCode.
func (code KeyCode) Value() uint32 {
	return uint32(code)
}

// Value returns the uint32 numeric representation of the RelativeCode.
func (code RelativeCode) Value() uint32 {
	return uint32(code)
}

// Value returns the uint32 numeric representation of the AbsoluteCode.
func (code AbsoluteCode) Value() uint32 {
	return uint32(code)
}

// Value returns the uint32 numeric representation of the SwitchCode.
func (code SwitchCode) Value() uint32 {
	return uint32(code)
}

// Value returns the uint32 numeric representation of the MiscCode.
func (code MiscCode) Value() uint32 {
	return uint32(code)
}

// Value returns the uint32 numeric representation of the LEDCode.
func (code LEDCode) Value() uint32 {
	return uint32(code)
}

// Value returns the uint32 numeric representation of the RepeatCode.
func (code RepeatCode) Value() uint32 {
	return uint32(code)
}

// Value returns the uint32 numeric representation of the SoundCode.
func (code SoundCode) Value() uint32 {
	return uint32(code)
}

// Value returns the uint32 numeric representation of the BusCode.
func (code BusCode) Value() uint32 {
	return uint32(code)
}

// Value returns the uint32 numeric representation of the
// MultiTouchCode.
func (code MultiTouchCode) Value() uint32 {
	return uint32(code)
}

// Value returns the uint32 numeric representation of the FFCode.
func (code FFCode) Value() uint32 {
	return uint32(code)
}

// Value returns the uint32 numeric representation of the FFStatusCode.
func (code FFStatusCode) Value() uint32 {
	return uint32(code)
}
