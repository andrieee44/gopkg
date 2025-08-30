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
	Pretty() string
	Value() uint16
}

// PropCode identifies a property supported by an input device.
//
//go:generate go run github.com/dmarkham/enumer -type=PropCode
type PropCode uint16

// EventCode identifies the broad event category for an input
// event.
//
//go:generate go run github.com/dmarkham/enumer -type=EventCode
type EventCode uint16

// SyncCode describes synchronization events that delimit
// packets of input data or change reporting mode.
//
//go:generate go run github.com/dmarkham/enumer -type=SyncCode
type SyncCode uint16

// KeyCode represents a keyboard or button key code.
//
//go:generate go run github.com/dmarkham/enumer -type=KeyCode
type KeyCode uint16

// RelativeCode describes relative axes, such as pointer
// movement, scroll wheels, and tilt sensors.
//
//go:generate go run github.com/dmarkham/enumer -type=RelativeCode
type RelativeCode uint16

// AbsoluteCode describes absolute axes, such as touch‐screen
// coordinates, joystick positions, or tablet pressure.
//
//go:generate go run github.com/dmarkham/enumer -type=AbsoluteCode
type AbsoluteCode uint16

// SwitchCode describes switch events, usually binary toggles
// like lid, tablet mode, or proximity sensors.
//
//go:generate go run github.com/dmarkham/enumer -type=SwitchCode
type SwitchCode uint16

// MiscCode covers miscellaneous event codes that don’t fit
// into other categories, such as drive insert/eject, auto-repeat
// toggle, or power events.
//
//go:generate go run github.com/dmarkham/enumer -type=MiscCode
type MiscCode uint16

// LEDCode represents status LEDs on a device, such as
// keyboard or system LEDs.
//
//go:generate go run github.com/dmarkham/enumer -type=LEDCode
type LEDCode uint16

// RepeatCode defines auto‐repeat settings for keys.
//
//go:generate go run github.com/dmarkham/enumer -type=RepeatCode
type RepeatCode uint16

// SoundCode describes simple tone and sound events,
// typically used for system beeps.
//
//go:generate go run github.com/dmarkham/enumer -type=SoundCode
type SoundCode uint16

// BusCode identifies the hardware bus (USB, PCI, Bluetooth, etc.)
//
//go:generate go run github.com/dmarkham/enumer -type=BusCode
type BusCode uint16

// MultiTouchCode represents multi-touch event codes
// (MT_SLOT, MT_POSITION_X, MT_TRACKING_ID, etc.)
//
//go:generate go run github.com/dmarkham/enumer -type=MultiTouchCode
type MultiTouchCode uint16

// FFCode denotes force-feedback effect types
// (FF_RUMBLE, FF_SPRING, FF_PERIODIC, etc.)
//
//go:generate go run github.com/dmarkham/enumer -type=FFCode
type FFCode uint16

// FFStatusCode holds the status value for an FF_STATUS event.
//
//go:generate go run github.com/dmarkham/enumer -type=FFStatusCode
type FFStatusCode uint16

// Value returns the uint16 numeric representation of the PropCode.
func (code PropCode) Value() uint16 {
	return uint16(code)
}

// Value returns the uint16 numeric representation of the EventCode.
func (code EventCode) Value() uint16 {
	return uint16(code)
}

// Value returns the uint16 numeric representation of the SyncCode.
func (code SyncCode) Value() uint16 {
	return uint16(code)
}

// Value returns the uint16 numeric representation of the KeyCode.
func (code KeyCode) Value() uint16 {
	return uint16(code)
}

// Value returns the uint16 numeric representation of the RelativeCode.
func (code RelativeCode) Value() uint16 {
	return uint16(code)
}

// Value returns the uint16 numeric representation of the AbsoluteCode.
func (code AbsoluteCode) Value() uint16 {
	return uint16(code)
}

// Value returns the uint16 numeric representation of the SwitchCode.
func (code SwitchCode) Value() uint16 {
	return uint16(code)
}

// Value returns the uint16 numeric representation of the MiscCode.
func (code MiscCode) Value() uint16 {
	return uint16(code)
}

// Value returns the uint16 numeric representation of the LEDCode.
func (code LEDCode) Value() uint16 {
	return uint16(code)
}

// Value returns the uint16 numeric representation of the RepeatCode.
func (code RepeatCode) Value() uint16 {
	return uint16(code)
}

// Value returns the uint16 numeric representation of the SoundCode.
func (code SoundCode) Value() uint16 {
	return uint16(code)
}

// Value returns the uint16 numeric representation of the BusCode.
func (code BusCode) Value() uint16 {
	return uint16(code)
}

// Value returns the uint16 numeric representation of the
// MultiTouchCode.
func (code MultiTouchCode) Value() uint16 {
	return uint16(code)
}

// Value returns the uint16 numeric representation of the FFCode.
func (code FFCode) Value() uint16 {
	return uint16(code)
}

// Value returns the uint16 numeric representation of the FFStatusCode.
func (code FFStatusCode) Value() uint16 {
	return uint16(code)
}

// Pretty returns the [PropCode] in the form "value (name)".
func (code PropCode) Pretty() string {
	return fmt.Sprintf("%d (%s)", code.Value(), code.String())
}

// Pretty returns the [EventCode] in the form "value (name)".
func (code EventCode) Pretty() string {
	return fmt.Sprintf("%d (%s)", code.Value(), code.String())
}

// Pretty returns the [SyncCode] in the form "value (name)".
func (code SyncCode) Pretty() string {
	return fmt.Sprintf("%d (%s)", code.Value(), code.String())
}

// Pretty returns the [KeyCode] in the form "value (name)".
func (code KeyCode) Pretty() string {
	return fmt.Sprintf("%d (%s)", code.Value(), code.String())
}

// Pretty returns the [RelativeCode] in the form "value (name)".
func (code RelativeCode) Pretty() string {
	return fmt.Sprintf("%d (%s)", code.Value(), code.String())
}

// Pretty returns the [AbsoluteCode] in the form "value (name)".
func (code AbsoluteCode) Pretty() string {
	return fmt.Sprintf("%d (%s)", code.Value(), code.String())
}

// Pretty returns the [SwitchCode] in the form "value (name)".
func (code SwitchCode) Pretty() string {
	return fmt.Sprintf("%d (%s)", code.Value(), code.String())
}

// Pretty returns the [MiscCode] in the form "value (name)".
func (code MiscCode) Pretty() string {
	return fmt.Sprintf("%d (%s)", code.Value(), code.String())
}

// Pretty returns the [LEDCode] in the form "value (name)".
func (code LEDCode) Pretty() string {
	return fmt.Sprintf("%d (%s)", code.Value(), code.String())
}

// Pretty returns the [RepeatCode] in the form "value (name)".
func (code RepeatCode) Pretty() string {
	return fmt.Sprintf("%d (%s)", code.Value(), code.String())
}

// Pretty returns the [SoundCode] in the form "value (name)".
func (code SoundCode) Pretty() string {
	return fmt.Sprintf("%d (%s)", code.Value(), code.String())
}

// Pretty returns the [BusCode] in the form "value (name)".
func (code BusCode) Pretty() string {
	return fmt.Sprintf("%d (%s)", code.Value(), code.String())
}

// Pretty returns the [MultiTouchCode] in the form "value (name)".
func (code MultiTouchCode) Pretty() string {
	return fmt.Sprintf("%d (%s)", code.Value(), code.String())
}

// Pretty returns the [FFCode] in the form "value (name)".
func (code FFCode) Pretty() string {
	return fmt.Sprintf("%d (%s)", code.Value(), code.String())
}

// Pretty returns the [FFStatusCode] in the form "value (name)".
func (code FFStatusCode) Pretty() string {
	return fmt.Sprintf("%d (%s)", code.Value(), code.String())
}
