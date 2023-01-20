package templates

type TokenKind string

const (
	KEYWORD_TKN TokenKind = "KEYWORD"
	STRING_TKN            = "STRING"
	CONTENT_TKN           = "CONTENT"
)

type Token struct {
	Value  string
	Kind   TokenKind
	line   int
	column int
}

func NewToken(value string, kind TokenKind, line, column int) *Token {
	return &Token{
		value, kind, line, column,
	}
}
