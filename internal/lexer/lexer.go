package lexer

import (
	"arenac/internal/token"
	"bufio"
	"errors"
	"unicode"
)

type Lexer struct {
	r   *bufio.Reader
	pos int
}

func New(r *bufio.Reader) *Lexer {
	return &Lexer{r: r, pos: 0}
}

func (l *Lexer) Next() (*token.Token, error) {
	for {
		r, _, err := l.r.ReadRune()
		if err != nil {
			return nil, err
		}

		l.pos += 1

		switch r {
		case '\t', '\n', '\r', ' ':
			continue
		case '(':
			return &token.Token{Kind: token.START_PAREN, Value: "(", Pos: l.pos}, nil
		case ')':
			return &token.Token{Kind: token.END_PAREN, Value: ")", Pos: l.pos}, nil
		case '{':
			return &token.Token{Kind: token.START_BLOCK, Value: "{", Pos: l.pos}, nil
		case '}':
			return &token.Token{Kind: token.END_BLOCK, Value: "}", Pos: l.pos}, nil
		case ';':
			return &token.Token{Kind: token.SEMICOLON, Value: ";", Pos: l.pos}, nil
		case ':':
			return &token.Token{Kind: token.COLON, Value: ":", Pos: l.pos}, nil
		case ',':
			return &token.Token{Kind: token.COMMA, Value: ",", Pos: l.pos}, nil
		case 'p':
			err := l.r.UnreadRune()
			if err != nil {
				return nil, err
			}

			tok, err := l.tryKeyword("proc", token.PROC)
			if err != nil {
				return nil, err
			}

			if tok != nil {
				return tok, nil
			}

			fallthrough
		default:
			return l.ident(r)
		}
	}
}

func (l *Lexer) ident(start rune) (*token.Token, error) {
	pos := l.pos
	value := []rune{start}

	for {
		r, _, err := l.r.ReadRune()
		if err != nil {
			return nil, err
		}

		if !(unicode.IsUpper(r) || unicode.IsLower(r) || unicode.IsDigit(r)) {
			err := l.r.UnreadRune()
			if err != nil {
				return nil, err
			}

			return &token.Token{Kind: token.IDENT, Value: string(value), Pos: pos}, nil
		}

		l.pos += 1
	}
}

func (l *Lexer) tryKeyword(keyword string, kind token.Kind) (*token.Token, error) {
	tok, err := l.keyword("proc", token.PROC)
	if errors.Is(err, &NotExpectedError{}) {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}

	return tok, nil
}

func (l *Lexer) keyword(keyword string, kind token.Kind) (*token.Token, error) {
	pos := l.pos
	value := []rune{}

	for _, expected := range keyword {
		r, _, err := l.r.ReadRune()
		if err != nil {
			return nil, err
		}

		if r != expected {
			return nil, &NotExpectedError{
				Expected: []token.Kind{kind},
				Actual: token.Token{
					Kind:  token.IDENT,
					Value: string(value),
					Pos:   pos,
				},
			}
		}

		value = append(value, r)
	}

	return &token.Token{Kind: kind, Value: string(value), Pos: pos}, nil
}
