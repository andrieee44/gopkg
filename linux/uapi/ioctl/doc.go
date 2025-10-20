//go:build linux

// Package ioctl provides a Go interface to the Linux ioctl userspace API,
// mirroring constants and function names from the corresponding C headers.
//
// It retains leading underscores found in the C definitions for accuracy,
// and extends the raw API with strict argument validation and clear error
// reporting. In addition to the low‑level bindings, it offers exported
// helpers for common ioctl patterns to reduce boilerplate and encourage
// consistent usage.
package ioctl
