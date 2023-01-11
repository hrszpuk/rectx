package templates

type TokenKind string

const (
	KEYWORD_TKN TokenKind = "KEYWORD"
	STRING_TKN            = "STRING"
	CONTENT_TKN           = "CONTENT"
)

type Token struct {
	Value string
	Kind  TokenKind
}

func NewToken(value string, kind TokenKind) *Token {
	return &Token{
		value, kind,
	}
}
