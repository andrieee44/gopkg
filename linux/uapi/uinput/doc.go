//go:build linux

// Package uinput provides Go bindings for the Linux uinput subsystem.
//
// It allows user-space programs to create and manage virtual input devices,
// such as keyboards, mice, touchscreens, and gamepads. The package defines
// constants and helpers for constructing ioctl request codes used to
// configure devices, enable event types, and emit input events.
//
// These bindings are designed for low-level control and integration with
// Linux input drivers. All ioctl codes are encoded using platform-aware
// bitfield logic to ensure compatibility across kernel versions.
package uinput
