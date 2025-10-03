package elex

import (
	"errors"
	"fmt"
	"io"
	"unicode/utf8"
)

// Lexer reads runes from an [io.Reader]. It hides buffer management
// and UTF-8 decoding so clients can focus on higher-level tokenization.
//
// A Lexer is not safe for concurrent use. It should be confined to
// a single goroutine.
type Lexer struct {
	rd                   io.Reader
	buf                  []byte
	read, eof            int
	start, prev, current Position
}

const (
	// EOF is the sentinel rune returned by [Lexer.Next] when the input
	// stream has been fully consumed.
	EOF = -1

	defaultBufSize = 4096
)

// NewLexer creates a new Lexer that reads from the given io.Reader.
// The reader must deliver valid UTF-8. The Lexer starts at the
// beginning of the stream and is ready to consume input and produce
// tokens.
func NewLexer(rd io.Reader) *Lexer {
	var start Position

	start = NewPosition(0, 0, 1, 1)

	return &Lexer{
		rd:      rd,
		buf:     make([]byte, defaultBufSize),
		eof:     -1,
		start:   start,
		prev:    start,
		current: start,
	}
}

// Next returns the next rune from the input stream.
// On success it returns the decoded rune and a nil error.
// When the end of input is reached, it returns [EOF] and a nil error.
// If the input contains invalid UTF-8 or the underlying reader
// misbehaves, Next returns [utf8.RuneError] along with a non-nil error.
func (lx *Lexer) Next() (rune, error) {
	var (
		r    rune
		size int
		err  error
	)

	for lx.eof == -1 && lx.read-lx.current.offset < utf8.UTFMax {
		err = lx.fill()
		if err != nil {
			return utf8.RuneError, err
		}
	}

	if lx.eof != -1 && lx.current.offset >= lx.eof {
		return EOF, nil
	}

	r, size = utf8.DecodeRune(lx.buf[lx.current.offset:lx.read])
	if r == utf8.RuneError {
		return utf8.RuneError, InvalidUTF8Error{
			offset: lx.current.absolute,
			b:      lx.buf[lx.current.offset],
		}
	}

	lx.prev = lx.current
	lx.current.absolute += size
	lx.current.offset += size

	if r == '\n' {
		lx.current.line++
		lx.current.column = 0
	}

	lx.current.column++

	return r, nil
}

// Peek returns the next rune without advancing the input. Peek returns
// the rune and any error encountered while reading.
func (lx *Lexer) Peek() (rune, error) {
	var (
		r   rune
		err error
	)

	r, err = lx.Next()
	if err != nil {
		return r, err
	}

	lx.Backup()

	return r, nil
}

// PeekLiteral returns the current buffered input between the start and
// current positions, without advancing the lexer. Useful for inspecting
// the matched text before emitting a token or performing classification.
func (lx *Lexer) PeekLiteral() string {
	return string(lx.buf[lx.start.offset:lx.current.offset])
}

// Backup undoes the most recent call to [Lexer.Next].
// [Lexer] only tracks the position of the last
// rune returned by [Lexer.Next]. You can only back up once.
// Calling Backup multiple times in a row will not rewind further.
func (lx *Lexer) Backup() {
	lx.current = lx.prev
}

// Ignore discards all input read since the last token boundary.
// Useful for skipping whitespace, comments, or delimiters before
// emitting the next token.
func (lx *Lexer) Ignore() {
	lx.start = lx.current
}

// ErrorToken creates a token of the given error type with a formatted
// message. The token covers the input span from the last token boundary.
// Useful for reporting unexpected input or malformed constructs.
func (lx *Lexer) ErrorToken(errorType int, format string, args ...any) Token {
	return NewToken(errorType, fmt.Sprintf(format, args...), lx.start)
}

// Emit returns a token using the input consumed since the last
// Emit or Ignore. Clients typically call Emit after matching a
// meaningful sequence. Tokens can be collected in a slice, sent
// into a channel, or forwarded directly to a parser for streaming.
func (lx *Lexer) Emit(typ int) Token {
	var token Token

	token = NewToken(
		typ,
		lx.PeekLiteral(),
		lx.start,
	)

	lx.start = lx.current

	return token
}

// Delim consumes input until the given delimiter string is found,
// or until EOF. Useful for skipping over multi-line constructs like
// block comments or string literals. Returns true if the delimiter
// was found, false if EOF was reached before matching, and an error
// if reading fails. Panics if the provided delimiter string is empty.
func (lx *Lexer) Delim(delim string) (bool, error) {
	var (
		buf, delimRunes []rune
		r               rune
		full            bool
		idx, i, j       int
		err             error
	)

	if delim == "" {
		panic("elex.Lexer.Delim: delimiter string must not be empty")
	}

	delimRunes = []rune(delim)
	buf = make([]rune, len(delimRunes))

	for {
		r, err = lx.Next()
		if err != nil {
			return false, err
		}

		if r == EOF {
			return false, nil
		}

		if idx+1 == len(buf) {
			full = true
		}

		buf[idx] = r
		idx = (idx + 1) % len(buf)

		if !full {
			continue
		}

		i = 0
		j = idx
		for {
			if delimRunes[i] != buf[j] {
				break
			}

			i++
			j = (j + 1) % len(buf)

			if i == len(delimRunes) {
				return true, nil
			}
		}
	}
}

// Oneshot tests the next rune with a predicate and consumes it if
// matched. If the rune fails the predicate, it is backed up. Returns
// true if matched, false if not, and an error if reading fails.
func (lx *Lexer) Oneshot(fn PredicateFn) (bool, error) {
	var (
		r   rune
		ok  bool
		err error
	)

	r, err = lx.Next()
	if err != nil {
		return false, err
	}

	ok, err = fn(lx, r)
	if err != nil {
		return false, err
	}

	if !ok {
		lx.Backup()
	}

	return ok, nil
}

// Iterate consumes runes while they satisfy the given predicate.
// Returns the number of runes matched, or an error if reading fails.
// Stops when a rune fails the predicate and backs up before it.
func (lx *Lexer) Iterate(fn PredicateFn) (int, error) {
	var (
		n   int
		ok  bool
		err error
	)

	for {
		ok, err = lx.Oneshot(fn)
		if err != nil {
			return 0, err
		}

		if !ok {
			return n, nil
		}

		n++
	}
}

func (lx *Lexer) fill() error {
	var (
		n   int
		err error
	)

	if lx.read == len(lx.buf) {
		if lx.start.offset == 0 {
			lx.buf = append(lx.buf, make([]byte, len(lx.buf))...)
		} else {
			copy(lx.buf, lx.buf[lx.start.offset:lx.read])
			lx.read -= lx.start.offset
			lx.prev.offset -= lx.start.offset
			lx.current.offset -= lx.start.offset
			lx.start.offset = 0
		}
	}

	n, err = lx.rd.Read(lx.buf[lx.read:])
	if n < 0 || n > len(lx.buf[lx.read:]) {
		return fmt.Errorf(
			"lexe.Lexer: bogus io.Reader, bytes read %d out of %d",
			n,
			len(lx.buf[lx.read:]),
		)
	}

	if n == 0 && err == nil {
		return io.ErrNoProgress
	}

	lx.read += n

	if err != nil {
		if errors.Is(err, io.EOF) {
			lx.eof = lx.read

			return nil
		}

		return err
	}

	return nil
}
