package lexer_test

import (
	"testing"

	"github.com/andrieee44/gopkg/cmd/evdevd/lexer"
	"github.com/andrieee44/gopkg/lib/elex"
)

type lexerTest struct {
	input    string
	expected []elex.Token
}

var lexerTests = []lexerTest{
	{
		input: "\t\t\t\n" +
			"//\t\n" +
			"////\t\n" +
			"\t\t//\t\n" +
			"// Comment 1\n" +
			"// Comment 2\n" +
			"/**/\n" +
			"/***/\n" +
			"/****/\n" +
			"/**a/*/\n" +
			"/* Multi Comment */ // Line /* */ Comment\n",
		expected: []elex.Token{
			newToken(lexer.TokenCommentLine, "//\t\n", elex.NewPosition(0, 4, 2, 1)),
			newToken(lexer.TokenCommentLine, "////\t\n", elex.NewPosition(0, 8, 3, 1)),
			newToken(lexer.TokenCommentLine, "//\t\n", elex.NewPosition(0, 16, 4, 3)),
			newToken(lexer.TokenCommentLine, "// Comment 1\n", elex.NewPosition(0, 20, 5, 1)),
			newToken(lexer.TokenCommentLine, "// Comment 2\n", elex.NewPosition(0, 33, 6, 1)),
			newToken(lexer.TokenCommentBlock, "/**/", elex.NewPosition(0, 46, 7, 1)),
			newToken(lexer.TokenCommentBlock, "/***/", elex.NewPosition(0, 51, 8, 1)),
			newToken(lexer.TokenCommentBlock, "/****/", elex.NewPosition(0, 57, 9, 1)),
			newToken(lexer.TokenCommentBlock, "/**a/*/", elex.NewPosition(0, 64, 10, 1)),
			newToken(lexer.TokenCommentBlock, "/* Multi Comment */", elex.NewPosition(0, 72, 11, 1)),
			newToken(lexer.TokenCommentLine, "// Line /* */ Comment\n", elex.NewPosition(0, 92, 11, 21)),
			newToken(lexer.TokenEOF, "", elex.NewPosition(0, 114, 12, 1)),
		},
	},
	{
		input: "+-*/%<>:==!=<=>=&&||!&|^&^<<>>:=+=-=*=/=%=&=|=^=&^=<<=>>=++--....:.=.....",
		expected: []elex.Token{
			newToken(lexer.TokenAdd, "+", elex.NewPosition(0, 0, 1, 1)),
			newToken(lexer.TokenSub, "-", elex.NewPosition(0, 1, 1, 2)),
			newToken(lexer.TokenMul, "*", elex.NewPosition(0, 2, 1, 3)),
			newToken(lexer.TokenDiv, "/", elex.NewPosition(0, 3, 1, 4)),
			newToken(lexer.TokenMod, "%", elex.NewPosition(0, 4, 1, 5)),
			newToken(lexer.TokenLess, "<", elex.NewPosition(0, 5, 1, 6)),
			newToken(lexer.TokenGreater, ">", elex.NewPosition(0, 6, 1, 7)),
			newToken(lexer.TokenColon, ":", elex.NewPosition(0, 7, 1, 8)),
			newToken(lexer.TokenEqual, "==", elex.NewPosition(0, 8, 1, 9)),
			newToken(lexer.TokenNotEqual, "!=", elex.NewPosition(0, 10, 1, 11)),
			newToken(lexer.TokenLessEqual, "<=", elex.NewPosition(0, 12, 1, 13)),
			newToken(lexer.TokenGreaterEqual, ">=", elex.NewPosition(0, 14, 1, 15)),
			newToken(lexer.TokenLogicalAnd, "&&", elex.NewPosition(0, 16, 1, 17)),
			newToken(lexer.TokenLogicalOr, "||", elex.NewPosition(0, 18, 1, 19)),
			newToken(lexer.TokenLogicalNot, "!", elex.NewPosition(0, 20, 1, 21)),
			newToken(lexer.TokenAnd, "&", elex.NewPosition(0, 21, 1, 22)),
			newToken(lexer.TokenOr, "|", elex.NewPosition(0, 22, 1, 23)),
			newToken(lexer.TokenXor, "^", elex.NewPosition(0, 23, 1, 24)),
			newToken(lexer.TokenAndNot, "&^", elex.NewPosition(0, 24, 1, 25)),
			newToken(lexer.TokenLShift, "<<", elex.NewPosition(0, 26, 1, 27)),
			newToken(lexer.TokenRShift, ">>", elex.NewPosition(0, 28, 1, 29)),
			newToken(lexer.TokenColon, ":", elex.NewPosition(0, 30, 1, 31)),
			newToken(lexer.TokenAssign, "=", elex.NewPosition(0, 31, 1, 32)),
			newToken(lexer.TokenAddAssign, "+=", elex.NewPosition(0, 32, 1, 33)),
			newToken(lexer.TokenSubAssign, "-=", elex.NewPosition(0, 34, 1, 35)),
			newToken(lexer.TokenMulAssign, "*=", elex.NewPosition(0, 36, 1, 37)),
			newToken(lexer.TokenDivAssign, "/=", elex.NewPosition(0, 38, 1, 39)),
			newToken(lexer.TokenModAssign, "%=", elex.NewPosition(0, 40, 1, 41)),
			newToken(lexer.TokenAndAssign, "&=", elex.NewPosition(0, 42, 1, 43)),
			newToken(lexer.TokenOrAssign, "|=", elex.NewPosition(0, 44, 1, 45)),
			newToken(lexer.TokenXorAssign, "^=", elex.NewPosition(0, 46, 1, 47)),
			newToken(lexer.TokenAndNotAssign, "&^=", elex.NewPosition(0, 48, 1, 49)),
			newToken(lexer.TokenLShiftAssign, "<<=", elex.NewPosition(0, 51, 1, 52)),
			newToken(lexer.TokenRShiftAssign, ">>=", elex.NewPosition(0, 54, 1, 55)),
			newToken(lexer.TokenInc, "++", elex.NewPosition(0, 57, 1, 58)),
			newToken(lexer.TokenDec, "--", elex.NewPosition(0, 59, 1, 60)),
			newToken(lexer.TokenEllipsis, "...", elex.NewPosition(0, 61, 1, 62)),
			newToken(lexer.TokenDot, ".", elex.NewPosition(0, 64, 1, 65)),
			newToken(lexer.TokenColon, ":", elex.NewPosition(0, 65, 1, 66)),
			newToken(lexer.TokenDot, ".", elex.NewPosition(0, 66, 1, 67)),
			newToken(lexer.TokenAssign, "=", elex.NewPosition(0, 67, 1, 68)),
			newToken(lexer.TokenEllipsis, "...", elex.NewPosition(0, 68, 1, 69)),
			newToken(lexer.TokenDot, ".", elex.NewPosition(0, 71, 1, 72)),
			newToken(lexer.TokenDot, ".", elex.NewPosition(0, 72, 1, 73)),
			newToken(lexer.TokenEOF, "", elex.NewPosition(0, 73, 1, 74)),
		},
	},
}

func newToken(typ lexer.TokenType, value string, pos elex.Position) elex.Token {
	return elex.NewToken(int(typ), value, pos)
}

func zeroedAbs(token elex.Token) elex.Token {
	return elex.NewToken(token.Type(), token.Value(), elex.NewPosition(0,
		token.Position().Offset(),
		token.Position().Line(),
		token.Position().Column(),
	))
}

func equalTokens(t *testing.T, input string, got, want elex.Token) {
	t.Helper()

	if zeroedAbs(got) != zeroedAbs(want) {
		t.Errorf(`
Input:    %q
          Got                                      Want
Type:     %-40s %-25s
Value:    %-40q %-25q
Offset:   %-5d                                    %-5d
Line:     %-5d                                    %-5d
Column:   %-5d                                    %-5d
`, input, lexer.TokenType(got.Type()), lexer.TokenType(want.Type()),
			got.Value(), want.Value(),
			got.Position().Offset(), want.Position().Offset(),
			got.Position().Line(), want.Position().Line(),
			got.Position().Column(), want.Position().Column(),
		)
	}
}

func TestLexer(t *testing.T) {
	var (
		test  lexerTest
		ch    <-chan elex.Token
		token elex.Token
		i     int
	)

	t.Parallel()

	for _, test = range lexerTests {
		ch = lexer.NewLexer(test.input)
		i = 0

		for token = range ch {
			equalTokens(t, test.input, token, test.expected[i])
			i++
		}
	}
}
