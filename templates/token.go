package templates

type TokenKind int

const (
	FOLDER_TKN TokenKind = iota
	FILE_TKN
	STRING_TKN
	LEFT_BRACE_TKN
	RIGHT_BRACE_TKN
	CONTENT_TKN
	COMMAND_TKN
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
