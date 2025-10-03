package lexer

import "github.com/andrieee44/gopkg/lib/elex"

const (
	decimalDigits string = "0123456789"
	binaryDigits  string = "01"
	octalDigits   string = "01234567"
	hexDigits     string = "0123456789AaBbCcDdEeFf"
)

func intFn(digits string, fn stateFn) stateFn {
	return func(lx *lexer) stateFn {
		var err error

		_, err = lx.engine.Oneshot(elex.AcceptRune('_'))
		if err != nil {
			return errorln(err)
		}

		_, err = lx.engine.Iterate(elex.OptionalSep(digits, '_'))
		if err != nil {
			return errorln(err)
		}

		return fn
	}
}

func floatFn(digits, expPrefix string, floatType TokenType, expFn stateFn) stateFn {
	return func(lx *lexer) stateFn {
		var (
			ok  bool
			err error
		)

		_, err = lx.engine.Iterate(elex.OptionalSep(digits, '_'))
		if err != nil {
			return errorln(err)
		}

		ok, err = lx.engine.Oneshot(elex.Accept(expPrefix))
		if err != nil {
			return errorln(err)
		}

		if ok {
			return expFn
		}

		lx.emit(floatType)

		return begin
	}
}

func exponentFn(digits string, typ TokenType) stateFn {
	return func(lx *lexer) stateFn {
		var err error

		_, err = lx.engine.Oneshot(elex.Accept("+-"))
		if err != nil {
			return errorln(err)
		}

		_, err = lx.engine.Iterate(elex.OptionalSep(digits, '_'))
		if err != nil {
			return errorln(err)
		}

		lx.emit(typ)

		return begin
	}
}

func numberFn(r rune) stateFn {
	return func(_ *lexer) stateFn {
		switch r {
		case '0':
			return matchRuneFn(emitBeginFn(TokenDecimalInt), stringStateCat(
				stringState{match: ".", fn: decimalFloat},
				stringState{match: "Ee", fn: decimalExponent},
				stringState{match: "Bb", fn: binary},
				stringState{match: "Oo", fn: octal},
				stringState{match: "Xx", fn: hex},
				stringState{match: "_" + octalDigits, fn: octal},
			))
		case '.':
			return matchRuneFn(operatorFn('.'), stringStateCat(
				stringState{match: decimalDigits, fn: decimalFloat},
			))
		default:
			return decimal
		}
	}
}

func decimalExponent(_ *lexer) stateFn {
	return exponentFn(decimalDigits, TokenDecimalFloat)
}

func decimalFloat(_ *lexer) stateFn {
	return floatFn(decimalDigits, "Ee", TokenDecimalFloat, decimalExponent)
}

func decimal(_ *lexer) stateFn {
	return intFn(
		decimalDigits,
		matchRuneFn(emitBeginFn(TokenDecimalInt), stringStateCat(
			stringState{match: ".", fn: decimalFloat},
			stringState{match: "Ee", fn: decimalExponent},
		)),
	)
}

func binary(_ *lexer) stateFn {
	return intFn(binaryDigits, emitBeginFn(TokenBinary))
}

func octal(_ *lexer) stateFn {
	return intFn(octalDigits, emitBeginFn(TokenOctal))
}

func hexExponent(_ *lexer) stateFn {
	return exponentFn(hexDigits, TokenHexFloat)
}

func hexFloat(_ *lexer) stateFn {
	return floatFn(hexDigits, "Pp", TokenHexFloat, hexExponent)
}

func hex(_ *lexer) stateFn {
	return intFn(
		hexDigits,
		matchRuneFn(emitBeginFn(TokenHexInt), map[rune]stateFn{
			'.': hexFloat,
		}),
	)
}
