package ast

import "testing"

func TestShouldCreateProgram(t *testing.T) {
	program := NewProgram([]*Decl{
		NewDecl("main", nil, nil, nil),
	}, nil, nil)
	
	if program == nil {
		t.Error("Expected program to be created")
	}

	t.Logf("%#v", program)
}
