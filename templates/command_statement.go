package templates

import (
	"os/exec"
	"rectx/utilities"
	"strings"
)

type CommandStatement struct {
	Command string
}

func NewCommandStatement(command string) *CommandStatement {
	return &CommandStatement{
		Command: command,
	}
}

func (command *CommandStatement) Generate(_projectName string) {
	args := strings.Split(command.Command, " ")
	command.Command = args[0]
	args = args[1:]

	cmd := exec.Command(command.Command, args...)

	err := cmd.Run()

	utilities.Check(err)
}
