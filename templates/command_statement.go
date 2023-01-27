package templates

import (
	"fmt"
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

	message := fmt.Sprintf("Attempted to run \"%s\" from during template-based project generation failed.", command.Command)
	utilities.Check(err, true, message)
}

func (command *CommandStatement) GetName() string {
	return command.Command
}

func (command *CommandStatement) GetType() string {
	return "COMMAND"
}
