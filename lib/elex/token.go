package elex

// Token represents a single token produced by the lexer engine.
// It contains the token type, raw value, and source position.
type Token struct {
	typ   int
	value string
	pos   Position
}

// NewToken creates a new Token with the given type, value, and position.
// This is typically used internally by the lexer engine to construct tokens
// during parsing. It can also be used in tests to manually create token
// instances for assertions and comparisons.
func NewToken(typ int, value string, pos Position) Token {
	return Token{
		typ:   typ,
		value: value,
		pos:   pos,
	}
}

// Type returns the token's type identifier.
// This is used to distinguish between different token kinds
// (e.g. identifier, integer, operator).
func (t Token) Type() int {
	return t.typ
}

// Value returns the raw string value of the token as it appeared in the
// input stream.
func (t Token) Value() string {
	return t.value
}

// Position returns the starting position of the token in the input stream.
// Useful for diagnostics, error reporting, and source mapping.
func (t Token) Position() Position {
	return t.pos
}
