package lexer

import (
	"arenac/internal/token"
	"bufio"
	"strings"
	"testing"
)

func TestLexer(t *testing.T) {
	input := "(a b c )"
	expected := []token.Kind{
		token.START_PAREN,
		token.IDENT,
		token.IDENT,
		token.IDENT,
		token.END_PAREN,
	}
	
	RunLexerTest(t, input, expected)
}

func RunLexerTest(t *testing.T, input string, expected []token.Kind) {
	r := bufio.NewReader(strings.NewReader(input))
	l := New(r)

	for _, kind := range expected {
		tok, err := l.Next()
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}
		
		if tok.Kind != kind {
			t.Fatalf("expected token kind %s, got %s", kind, tok.Kind)
		}
		t.Logf("got token: %#v %s", tok, tok.Kind)
	}
}
