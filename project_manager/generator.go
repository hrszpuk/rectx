package project_manager

import (
	"fmt"
	"os"
	"os/exec"
	projectConfig "rectx/project_manager/config"
	"rectx/templates"
	"rectx/utilities"
)

func CreateNewProject(config *projectConfig.ProjectConfig, templateName string) {
	fmt.Println("Generating project...")

	utilities.Check(os.Mkdir(config.Project.Name, 0750))

	f, err := os.ReadFile(utilities.GetRectxPath() + "/templates/" + templateName)
	utilities.Check(err)

	parser := templates.NewTemplateParser(string(f))
	statements := parser.Parse()

	_ = exec.Command("cd", config.Project.Name)
	for _, statement := range statements {
		statement.Generate(config.Project.Name)
	}

	fmt.Println("Done!")
}
