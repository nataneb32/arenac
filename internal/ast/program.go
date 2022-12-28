package ast

import (
	"arenac/internal/ast/astKind"
	"arenac/internal/token"
)

type Program struct {
	start *token.Token
	end   *token.Token
	Decls []*Decl
}

func NewProgram(decls []*Decl, start *token.Token, end *token.Token) *Program {
	return &Program{
		Decls: decls,
		start: start,
		end:   end,
	}
}

func (program *Program) Start() *token.Token {
	return program.start
}
func (program *Program) End() *token.Token {
	return program.end
}

func (p *Program) Kind() astKind.Kind {
	return astKind.Program
}
