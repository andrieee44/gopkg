package elex

import "strings"

// PredicateFn tests whether a rune satisfies a condition and returns
// an error if evaluation fails.
type PredicateFn func(*Lexer, rune) (bool, error)

// Accept returns a predicate that matches any rune in the given set.
// Useful for matching character classes like digits or letters.
func Accept(match string) PredicateFn {
	return func(_ *Lexer, r rune) (bool, error) {
		return strings.ContainsRune(match, r), nil
	}
}

// AcceptRune returns a predicate that matches a single specific rune.
// Useful for matching fixed symbols or delimiters.
func AcceptRune(match rune) PredicateFn {
	return func(_ *Lexer, r rune) (bool, error) {
		return r == match, nil
	}
}

// Skip returns a predicate that matches any rune in the given set,
// and discards it. Useful for skipping ignorable characters like
// whitespace or separators.
func Skip(match string) PredicateFn {
	return func(lexer *Lexer, r rune) (bool, error) {
		var ok bool

		ok = strings.ContainsRune(match, r)
		if ok {
			lexer.Ignore()
		}

		return ok, nil
	}
}

// SkipRune returns a predicate that matches a specific rune and
// discards it. Useful for skipping fixed delimiters or padding.
func SkipRune(match rune) PredicateFn {
	return func(lexer *Lexer, r rune) (bool, error) {
		var ok bool

		ok = r == match
		if ok {
			lexer.Ignore()
		}

		return ok, nil
	}
}

// OptionalSep returns a predicate that matches a rune from the given
// set, optionally followed by a separator. Useful for grouped forms
// like 1_000_000, where digits may be split by a consistent delimiter.
// Returns true if the match succeeds, false otherwise, and reports
// any read errors encountered.
func OptionalSep(match string, sep rune) PredicateFn {
	return func(lexer *Lexer, r rune) (bool, error) {
		var err error

		if !strings.ContainsRune(match, r) {
			return false, nil
		}

		_, err = lexer.Oneshot(AcceptRune(sep))
		if err != nil {
			return false, err
		}

		return true, nil
	}
}
