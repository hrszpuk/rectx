package project_manager

import (
	"fmt"
	"os"
	projectConfig "rectx/project_manager/config"
	"rectx/utilities"
)

func CreateNewProject(config *projectConfig.ProjectConfig, templateName string) {
	fmt.Println("Generating project...")

	utilities.Check(os.Mkdir(config.Project.Name, 0750))

	fmt.Println("Done!")
}
