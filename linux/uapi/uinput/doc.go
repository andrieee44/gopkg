// Package uinput provides Go bindings for the Linux uinput subsystem,
// allowing user-space programs to create and manage virtual input devices.
// It defines constants and functions for constructing ioctl request codes
// used to configure devices, enable event types, and send input events.
// The package mirrors the uinput and evdev ioctl interfaces, enabling
// fine-grained control over virtual keyboards, mice, gamepads, and other
// input devices from Go.
package uinput
