package lexer

import (
	"errors"
	"fmt"
)

func operatorFn(r rune) stateFn {
	return func(_ *lexer) stateFn {
		switch r {
		case '+':
			return matchRuneFn(emitBeginFn(TokenAdd), map[rune]stateFn{
				'+': emitBeginFn(TokenInc),
				'=': emitBeginFn(TokenAddAssign),
			})
		case '-':
			return matchRuneFn(emitBeginFn(TokenSub), map[rune]stateFn{
				'-': emitBeginFn(TokenDec),
				'=': emitBeginFn(TokenSubAssign),
			})
		case '&':
			return matchRuneFn(emitBeginFn(TokenAnd), map[rune]stateFn{
				'=': emitBeginFn(TokenAndAssign),
				'&': emitBeginFn(TokenLogicalAnd),
				'^': matchRuneFn(emitBeginFn(TokenAndNot), map[rune]stateFn{
					'=': emitBeginFn(TokenAndNotAssign),
				}),
			})
		case '|':
			return matchRuneFn(emitBeginFn(TokenOr), map[rune]stateFn{
				'|': emitBeginFn(TokenLogicalOr),
				'=': emitBeginFn(TokenOrAssign),
			})
		case '=':
			return matchRuneFn(emitBeginFn(TokenAssign), map[rune]stateFn{
				'=': emitBeginFn(TokenEqual),
			})
		case '<':
			return matchRuneFn(emitBeginFn(TokenLess), map[rune]stateFn{
				'=': emitBeginFn(TokenLessEqual),
				'<': matchRuneFn(emitBeginFn(TokenLShift), map[rune]stateFn{
					'=': emitBeginFn(TokenLShiftAssign),
				}),
			})
		case '>':
			return matchRuneFn(emitBeginFn(TokenGreater), map[rune]stateFn{
				'=': emitBeginFn(TokenGreaterEqual),
				'>': matchRuneFn(emitBeginFn(TokenRShift), map[rune]stateFn{
					'=': emitBeginFn(TokenRShiftAssign),
				}),
			})
		case '*':
			return matchRuneFn(emitBeginFn(TokenMul), map[rune]stateFn{
				'=': emitBeginFn(TokenMulAssign),
			})
		case '/':
			return matchRuneFn(emitBeginFn(TokenDiv), map[rune]stateFn{
				'=': emitBeginFn(TokenDivAssign),
				'/': commentLine,
				'*': commentBlock,
			})
		case '^':
			return matchRuneFn(emitBeginFn(TokenXor), map[rune]stateFn{
				'=': emitBeginFn(TokenXorAssign),
			})
		case '%':
			return matchRuneFn(emitBeginFn(TokenMod), map[rune]stateFn{
				'=': emitBeginFn(TokenModAssign),
			})
		case '!':
			return matchRuneFn(emitBeginFn(TokenLogicalNot), map[rune]stateFn{
				'=': emitBeginFn(TokenNotEqual),
			})
		case '.':
			return matchRuneFn(emitBeginFn(TokenDot), map[rune]stateFn{
				'.': matchRuneFn(multiDot, map[rune]stateFn{
					'.': emitBeginFn(TokenEllipsis),
				}),
			})
		}

		return begin
	}
}

func multiDot(lx *lexer) stateFn {
	var (
		dot rune
		err error
	)

	lx.engine.Backup()
	lx.emit(TokenDot)

	dot, err = lx.engine.Next()
	if err != nil {
		return errorln(err)
	}

	if dot != '.' {
		panic(fmt.Errorf("expected '.', got %q instead", dot))
	}

	return emitBeginFn(TokenDot)
}

func commentBlock(lx *lexer) stateFn {
	var (
		ok  bool
		err error
	)

	ok, err = lx.engine.Delim("*/")
	if err != nil {
		return errorln(err)
	}

	if !ok {
		return errorln(errors.New("comment not terminated"))
	}

	lx.emit(TokenCommentBlock)

	return begin
}

func commentLine(lx *lexer) stateFn {
	var err error

	_, err = lx.engine.Delim("\n")
	if err != nil {
		return errorln(err)
	}

	lx.emit(TokenCommentLine)

	return begin
}
