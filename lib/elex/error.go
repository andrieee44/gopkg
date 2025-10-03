package elex

import "errors"

// InvalidUTF8Error reports that the lexer encountered a byte
// sequence that is not valid UTF-8. This error includes the
// absolute offset in the input stream and the offending byte.
type InvalidUTF8Error struct {
	offset int
	b      byte
}

// ErrBogusReader is returned when an [io.Reader] violates the
// [io.Reader] contract by reporting an impossible number of bytes
// read (negative or greater than the provided buffer length).
//
// This indicates a programming error in the reader implementation,
// not in the lexer itself.
var ErrBogusReader = errors.New("lexe.Lexer: bogus io.Reader")

// Offset returns the absolute byte offset in the input stream where the
// error occurred.
func (err InvalidUTF8Error) Offset() int {
	return err.offset
}

// Byte returns the offending byte value that could not be decoded as UTF-8.
func (err InvalidUTF8Error) Byte() byte {
	return err.b
}

// Error implements the error interface for [InvalidUTF8Error].
func (InvalidUTF8Error) Error() string {
	return "lexe.Lexer: invalid UTF-8 byte"
}
