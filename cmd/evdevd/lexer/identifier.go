package lexer

import (
	"unicode"

	"github.com/andrieee44/gopkg/lib/elex"
)

func identifier(lx *lexer) stateFn {
	var err error

	_, err = lx.engine.Iterate(func(_ *elex.Lexer, r rune) (bool, error) {
		return unicode.IsLetter(r) || unicode.IsNumber(r) || r == '_', nil
	})
	if err != nil {
		return errorln(err)
	}

	switch lx.engine.PeekLiteral() {
	case "break":
		lx.emit(TokenBreak)
	case "const":
		lx.emit(TokenConst)
	case "continue":
		lx.emit(TokenContinue)
	case "else":
		lx.emit(TokenElse)
	case "fallthrough":
		lx.emit(TokenFallthrough)
	case "for":
		lx.emit(TokenFor)
	case "func":
		lx.emit(TokenFunc)
	case "if":
		lx.emit(TokenIf)
	case "import":
		lx.emit(TokenImport)
	case "map":
		lx.emit(TokenMap)
	case "range":
		lx.emit(TokenRange)
	case "return":
		lx.emit(TokenReturn)
	case "switch":
		lx.emit(TokenSwitch)
	default:
		lx.emit(TokenIdentifier)
	}

	return begin
}
