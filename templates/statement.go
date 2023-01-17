package templates

import "fmt"

type Statement interface {
	Generate()
}

type BadStatement struct {
	expectedToken *Token
	tokenFound    *Token
}

func NewBadStatement(expected, found *Token) *BadStatement {
	return &BadStatement{
		expectedToken: expected,
		tokenFound:    found,
	}
}

func (bs *BadStatement) Generate() {
	fmt.Printf(
		"ERROR: Expected Token of type \"%s\" but found Token of type \"%s\" with values of \"%s\"!\n",
		bs.expectedToken.Kind,
		bs.tokenFound.Kind,
		bs.tokenFound.Value,
	)
}
