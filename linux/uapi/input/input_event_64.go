//go:build amd64 || arm64 || loong64 || mips64 || mips64le || ppc64 || ppc64le || riscv64 || s390x || wasm

package input

// EventTime stores the timestamp of an input event for 64‑bit
// architectures such as amd64 and arm64. It matches the layout used
// by the Linux kernel’s input_event struct for these platforms.
type EventTime struct {
	// Sec is the number of whole seconds since the Unix epoch
	// (January 1, 1970 UTC) at which the event occurred.
	Sec int64

	// Usec is the additional offset in microseconds past the value
	// in Sec. The combination of Sec and Usec provides microsecond
	// precision for the event timestamp.
	Usec int64
}
