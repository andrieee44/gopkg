package lexer

import (
	"strings"
	"unicode"

	"github.com/andrieee44/gopkg/lib/elex"
)

type stateFn func(*lexer) stateFn

type stringState struct {
	match string
	fn    stateFn
}

func stringStateCat(rStates ...stringState) map[rune]stateFn {
	var (
		result map[rune]stateFn
		rState stringState
		r      rune
		size   int
	)

	for _, rState = range rStates {
		size += len([]rune(rState.match))
	}

	result = make(map[rune]stateFn, size)

	for _, rState = range rStates {
		for _, r = range rState.match {
			result[r] = rState.fn
		}
	}

	return result
}

func matchRuneFn(null stateFn, matches map[rune]stateFn) stateFn {
	return func(lx *lexer) stateFn {
		var (
			r   rune
			fn  stateFn
			ok  bool
			err error
		)

		r, err = lx.engine.Next()
		if err != nil {
			return errorln(err)
		}

		fn, ok = matches[r]
		if !ok {
			fn = null

			lx.engine.Backup()
		}

		return fn
	}
}

func emitBeginFn(typ TokenType) stateFn {
	return func(lx *lexer) stateFn {
		lx.emit(typ)

		return begin
	}
}

func errorf(format string, args ...any) stateFn {
	return func(lx *lexer) stateFn {
		lx.ch <- lx.engine.ErrorToken(int(TokenError), format, args...)

		return nil
	}
}

func errorln(err error) stateFn {
	return errorf("%s", err.Error())
}

func begin(lx *lexer) stateFn {
	var (
		r   rune
		err error
	)

	for {
		r, err = lx.engine.Next()
		if err != nil {
			return errorln(err)
		}

		switch {
		case r == '\n':
			lx.engine.Ignore()

			if insertSemicolon(lx.lastTokenType) {
				lx.emit(TokenSemicolon)
			}
		case r == elex.EOF:
			if insertSemicolon(lx.lastTokenType) {
				lx.emit(TokenSemicolon)
			}

			lx.emit(TokenEOF)

			return nil
		case unicode.IsSpace(r):
			lx.engine.Ignore()
		case unicode.IsLetter(r) || r == '_':
			return identifier
		case strings.ContainsRune(".0123456789", r):
			return numberFn(r)
		case strings.ContainsRune("+-&|=<>*/^%!", r):
			return operatorFn(r)
		case strings.ContainsRune("(),;:{}[]", r):
			return punctuationFn(r)
		}
	}
}
