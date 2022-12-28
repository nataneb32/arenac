package ast

import (
	"arenac/internal/token"
	"arenac/internal/ast/astKind"
)

// Node is an interface that all AST nodes implement.
type Node interface {
	Kind() astKind.Kind
	Start() *token.Token
	End() *token.Token
}
