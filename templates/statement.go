package templates

import (
	"fmt"
)

type Statement interface {
	Generate(projectName string)
}

type BadStatement struct {
	expectedToken *Token
	tokenFound    *Token
	line          int
	column        int
}

func NewBadStatement(expected, found *Token) *BadStatement {
	return &BadStatement{
		expectedToken: expected,
		tokenFound:    found,
	}
}

func (bs *BadStatement) Generate(_projectName string) {
	fmt.Printf(
		"ERROR[%d, %d]: Expected Token of type \"%s\" but found Token of type \"%s\" with values of \"%s\"!\n",
		bs.line,
		bs.column,
		bs.expectedToken.Kind,
		bs.tokenFound.Kind,
		bs.tokenFound.Value,
	)
}
