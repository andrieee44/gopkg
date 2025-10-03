package lexer

func punctuationFn(r rune) stateFn {
	return func(lx *lexer) stateFn {
		switch r {
		case '(':
			lx.emit(TokenLParen)
		case ')':
			lx.emit(TokenRParen)
		case ',':
			lx.emit(TokenComma)
		case ';':
			lx.emit(TokenSemicolon)
		case ':':
			lx.emit(TokenColon)
		case '{':
			lx.emit(TokenLBrace)
		case '}':
			lx.emit(TokenRBrace)
		case '[':
			lx.emit(TokenLBracket)
		case ']':
			lx.emit(TokenRBracket)
		}

		return begin
	}
}
