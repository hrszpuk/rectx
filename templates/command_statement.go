package templates

import (
	"os/exec"
	"rectx/utilities"
	"strings"
)

type CommandStatement struct {
	Name    string
	Command string
	Args    []string
}

func NewCommandStatement(command string) *CommandStatement {
	args := strings.Split(command, " ")
	name := args[0]
	args = args[1:]
	return &CommandStatement{
		Name:    name,
		Command: command,
		Args:    args,
	}
}

func (command *CommandStatement) Generate(_projectName string) {
	cmd := exec.Command(command.Name, command.Args...)

	err := cmd.Run()

	utilities.Check(err)
}

func (command *CommandStatement) GetName() string {
	return command.Command
}

func (command *CommandStatement) GetType() string {
	return "COMMAND"
}
