package lexer

import (
	"arenac/internal/token"
	"fmt"
)

type NotExpectedError struct {
	Expected []token.Kind
	Actual   token.Token
}

func (e *NotExpectedError) Error() string {
	return fmt.Sprintf("expected one of %#v, got %v", e.Expected, e.Actual)
}
