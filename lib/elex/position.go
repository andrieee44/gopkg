package elex

// Position represents a location within the input stream. It tracks
// absolute and relative offsets, as well as line and column numbers
// for error reporting. Position values are immutable outside this
// package.
type Position struct {
	absolute, offset, line, column int
}

// NewPosition creates a new Position with the given absolute offset,
// buffer-relative offset, line, and column. This is primarily useful
// for testing or internal initialization.
func NewPosition(abs, off, line, col int) Position {
	return Position{
		absolute: abs,
		offset:   off,
		line:     line,
		column:   col,
	}
}

// Absolute returns the absolute byte offset from the start of the
// input stream.
func (p Position) Absolute() int {
	return p.absolute
}

// Offset returns the byte offset relative to the current buffer.
func (p Position) Offset() int {
	return p.offset
}

// Line returns the current line number, starting at 1.
func (p Position) Line() int {
	return p.line
}

// Column returns the current column number, starting at 1.
func (p Position) Column() int {
	return p.column
}
