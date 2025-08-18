// Package evdev provides Go bindings for the Linux input event subsystem,
// mirroring the constants, types, and structures defined in the kernel’s
// <linux/input.h> header. It exposes the event types, codes, and related
// definitions used to interpret data from /dev/input/event* devices,
// enabling direct handling of keyboard, mouse, touchscreen, joystick,
// and other input device events in Go programs.
//
// In addition to constant and type mappings, evdev preserves the original
// numeric values from the kernel headers to ensure binary compatibility
// with the Linux input API. The package also includes architecture‑specific
// representations of kernel structs so input event data can be read and
// parsed accurately across supported platforms.
package evdev
