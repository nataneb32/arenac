package ast

import "arenac/internal/ast/astKind"

func Walk(target Node, f func(Node) bool) {
	if !f(target) {
		return
	}

	switch target.Kind() {
	case astKind.ProcExpr:
		procexpr := target.(*ProcExpr)

		for _, param := range procexpr.Params {
			Walk(param, f)
		}

		Walk(procexpr.Value, f)
	case astKind.Program:
		program := target.(*Program)
		for _, decl := range program.Decls {
			Walk(decl, f)
		}
	case astKind.Decl:
		decl := target.(*Decl)
		Walk(decl.Value, f)
	default:
		panic("unreachable")
	}

}
