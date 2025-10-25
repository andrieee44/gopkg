//go:build linux

// Package ioctl exposes Go bindings for the Linux ioctl userspace API.
//
// It mirrors constants and function names from <linux/ioctl.h>. Struct
// names follow Go idioms: redundant prefixes are removed and names are
// camel-cased.
//
// Constants and functions retain their original C names, except for leading
// underscores, which are removed to comply with Go's export rules.
//
// The package extends the raw API with strict argument validation and clear
// error reporting. It also provides helpers for common ioctl patterns to
// reduce boilerplate and encourage consistent usage.
package ioctl
