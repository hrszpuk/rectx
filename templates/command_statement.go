package templates

import (
	"os/exec"
	"rectx/utilities"
)

type CommandStatement struct {
	Command string
}

func NewCommandStatement(command string) *CommandStatement {
	return &CommandStatement{
		Command: command,
	}
}

func (command *CommandStatement) Generate() {
	cmd := exec.Command(command.Command)

	err := cmd.Run()

	utilities.Check(err)
}
