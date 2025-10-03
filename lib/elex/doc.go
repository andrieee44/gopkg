// Package elex provides a lexer engine for building tokenizers.
//
// It abstracts away low-level concerns such as buffer management,
// UTF-8 decoding, and input handling. Clients define their own
// state machines and token types, while elex supplies the primitives
// and helpers needed to drive the process.
//
// The core type is [Lexer], which reads from an [io.Reader] and
// exposes operations for consuming runes and advancing through the
// input. Additional helpers simplify common lexer patterns,
// including predicate-driven matching for declarative rule
// construction.
//
// The package expects valid UTF-8 input and a well-behaved
// [io.Reader]. Errors are reported when the reader violates the
// [io.Reader] contract or when invalid UTF-8 sequences are
// encountered.
package elex
