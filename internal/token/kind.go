package token

type Kind int

const (
	IDENT Kind = iota
	PROC
	FN

	START_BLOCK
	END_BLOCK

	START_QUOTE
	END_QUOTE

	START_PAREN
	END_PAREN

	SEMICOLON
	COLON
	COMMA
)

func (k Kind) String() string {
	switch k {
	case IDENT:
		return "IDENT"
	case PROC:
		return "PROC"
	case FN:
		return "FN"
	case START_BLOCK:
		return "START_BLOCK"
	case END_BLOCK:
		return "END_BLOCK"
	case START_QUOTE:
		return "START_QUOTE"
	case END_QUOTE:
		return "END_QUOTE"
	case START_PAREN:
		return "START_PAREN"
	case END_PAREN:
		return "END_PAREN"
	case SEMICOLON:
		return "SEMICOLON"
	case COLON:
		return "COLON"
	case COMMA:
		return "COMMA"
	default:
		return "UNKNOWN"
	}
}
