//go:build linux

// Package uinput provides Go bindings for the Linux uinput subsystem.
//
// It allows user-space programs to create and manage virtual input devices,
// such as keyboards, mice, touchscreens, and gamepads.
//
// The package defines constants and helpers for constructing ioctl request
// codes used to configure devices, enable event types, and emit input events.
//
// All ioctl codes are encoded using platform-aware bitfield logic to ensure
// compatibility across kernel versions.
//
// Struct names follow Go idioms: redundant prefixes are removed and names
// are camel-cased. Constants and functions retain their original C names,
// except for leading underscores, which are removed to comply with Go's
// export rules.
package uinput
