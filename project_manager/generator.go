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
	fmt.Print("Generating project... ")

	utilities.Check(os.Mkdir(config.Project.Name, 0750), true, "Attempt to create project directory failed... Permission levels may not be sufficient?")

	f, err := os.ReadFile(utilities.GetRectxPath() + "/templates/" + templateName)
	if os.IsNotExist(err) {
		utilities.Check(err, true, "Attempt to read template file failed because it doesn't exist... What?")
	} else {
		utilities.Check(err, true, "Attempt to read template file failed.")
	}

	parser := templates.NewTemplateParser(string(f))
	statements := parser.Parse()

	_ = exec.Command("cd", config.Project.Name)
	for _, statement := range statements {
		statement.Generate(config.Project.Name)
	}

	fmt.Println("Done!")
}
