package project_manager

import (
	"fmt"
	projectConfig "rectx/project_manager/config"
	"rectx/templates"
	"strings"
)

func New() {
	pc := projectConfig.CreateDefaultConfig()

	fmt.Println("Project name: ")
	fmt.Scanln(&pc.Project.Name)

	var authors string
	fmt.Println("Authors (comma separated): ")
	fmt.Scanln(&authors)
	pc.Project.Authors = strings.Split(authors, ",")

	fmt.Println("Version (MAJOR.MINOR.PATCH): ")
	fmt.Scanln(&pc.Project.Version)

	templateName := GetTemplate()

	CreateNewProject(pc, templateName)
}

func build() {

}

func run() {

}
