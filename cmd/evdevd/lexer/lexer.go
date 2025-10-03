package lexer

import (
	"strings"

	"github.com/andrieee44/gopkg/lib/elex"
)

type lexer struct {
	engine        *elex.Lexer
	ch            chan<- elex.Token
	lastTokenType TokenType
}

func NewLexer(input string) <-chan elex.Token {
	var ch chan elex.Token

	ch = make(chan elex.Token)

	go (&lexer{
		engine: elex.NewLexer(strings.NewReader(input)),
		ch:     ch,
	}).run()

	return ch
}

func (lx *lexer) emit(typ TokenType) {
	lx.lastTokenType = typ
	lx.ch <- lx.engine.Emit(int(typ))
}

func (lx *lexer) run() {
	var state stateFn

	for state = begin; state != nil; {
		state = state(lx)
	}

	close(lx.ch)
}
