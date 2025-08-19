//go:build 386 || arm || mips || mipsle

package evdev

// InputEventTime stores the timestamp of an input event for 32‑bit
// architectures such as 386 and arm. It matches the layout used
// by the Linux kernel’s input_event struct for these platforms.
type InputEventTime struct {
	// Sec is the number of whole seconds since the Unix epoch
	// (January 1, 1970 UTC) at which the event occurred.
	Sec uint64

	// Usec is the additional offset in microseconds past the value
	// in Sec. The combination of Sec and Usec provides microsecond
	// precision for the event timestamp.
	Usec uint64
}
