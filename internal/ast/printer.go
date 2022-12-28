package ast

import "arenac/internal/ast/astKind"

func PrintTree(node Node) string {
	value := ""
	nest := 0
	Walk(node, func(node Node) bool {
		switch node.Kind() {
		case astKind.ProcExpr:
			procexpr := node.(*ProcExpr)
			value += "ProcExpr {\n"
			nest++
			for _, param := range procexpr.Params {
				value += printNestedTree(param, nest)
			}
		default:
			panic("unreachable")
		}

		return true
	})

	return value
}

func printNestedTree(value string, nest int) {
	ret := ""
	for i := 0; i < nest; i++ {
		ret += "--"
	}
	return ret + " " + value
}
