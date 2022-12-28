package ast

import (
	"arenac/internal/ast/astKind"
	"arenac/internal/token"
)

type Decl struct {
	start *token.Token
	end   *token.Token

	Name string
	Value Expr
}

func NewDecl(name string, value Expr, start *token.Token, end *token.Token) *Decl {
	return &Decl{
		Name:  name,
		Value: value,
		start: start,
		end:   end,
	}
}

func (decl *Decl) Start() *token.Token {
	return decl.start
}

func (decl *Decl) End() *token.Token {
	return decl.end
}

func (decl *Decl) Kind() astKind.Kind {
	return astKind.Decl
}
