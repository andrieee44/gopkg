package lexer

//go:generate go run github.com/dmarkham/enumer -type=TokenType
type TokenType int

const (
	TokenError TokenType = iota
	TokenEOF
	TokenCommentLine
	TokenCommentBlock
	TokenIdentifier
	TokenBreak
	TokenConst
	TokenContinue
	TokenElse
	TokenFallthrough
	TokenFor
	TokenFunc
	TokenIf
	TokenImport
	TokenMap
	TokenRange
	TokenReturn
	TokenSwitch
	TokenAdd
	TokenSub
	TokenMul
	TokenDiv
	TokenMod
	TokenAnd
	TokenOr
	TokenXor
	TokenLShift
	TokenRShift
	TokenAddAssign
	TokenSubAssign
	TokenMulAssign
	TokenDivAssign
	TokenModAssign
	TokenAndAssign
	TokenOrAssign
	TokenXorAssign
	TokenLShiftAssign
	TokenRShiftAssign
	TokenLogicalAnd
	TokenLogicalOr
	TokenGreater
	TokenInc
	TokenDec
	TokenEqual
	TokenLess
	TokenGreaterEqual
	TokenAssign
	TokenLogicalNot
	TokenNotEqual
	TokenLessEqual
	TokenAndNotAssign
	TokenComma
	TokenEllipsis
	TokenLParen
	TokenLBracket
	TokenLBrace
	TokenSemicolon
	TokenDot
	TokenRParen
	TokenRBracket
	TokenRBrace
	TokenAndNot
	TokenColon
	TokenDecimalInt
	TokenBinary
	TokenOctal
	TokenHexInt
	TokenDecimalFloat
	TokenHexFloat
	TokenRawString
	TokenInterpretedString
)

func insertSemicolon(typ TokenType) bool {
	switch typ {
	case TokenIdentifier,
		TokenDecimalInt,
		TokenBinary,
		TokenOctal,
		TokenDecimalFloat,
		TokenHexFloat,
		TokenRawString,
		TokenInterpretedString,
		TokenBreak,
		TokenContinue,
		TokenFallthrough,
		TokenReturn,
		TokenInc,
		TokenDec,
		TokenRParen,
		TokenRBracket,
		TokenRBrace:
		return true
	default:
		return false
	}
}
