package astKind

type Kind int

const (
	Program Kind = iota

	TypeDecl
	Decl

	IfExpr
	ForExpr
	ProcExpr
)
