package ast

import (
	"arenac/internal/ast/astKind"
	"arenac/internal/token"
)

type ProcExpr struct {
	start *token.Token
	end   *token.Token
	Params []ProcParam
	Value  Expr
}

func NewProcExpr(params []ProcParam, start *token.Token, end *token.Token) *ProcExpr {
	return &ProcExpr{
		start: start,
		end:   end,
		Params: params,
	}
}

func (procexpr *ProcExpr) Kind() astKind.Kind {
	return astKind.ProcExpr
}

func (procexpr *ProcExpr) Start() *token.Token {
	return procexpr.start
}
func (procexpr *ProcExpr) End() *token.Token {
	return procexpr.end
}
