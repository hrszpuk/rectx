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

	pc.Dump(pc.Project.Name + "/project.rectx")
}

func GetTemplate() string {
	fmt.Println("Pick a template:")
	var templateList []string = templates.FetchTemplates()
	for _, name := range templateList {
		fmt.Printf("- %s\n", name)
	}

	var chosenTemplateName string
	validTemplateChosen := false
	fmt.Println("\nType the name of the template you want to use: ")
	fmt.Scanln(&chosenTemplateName)
	for _, name := range templateList {
		if chosenTemplateName == name {
			validTemplateChosen = true
			break
		}
	}

	if !validTemplateChosen {
		fmt.Printf("\"%s\" is not a valid template!", chosenTemplateName)
		return GetTemplate()
	} else {
		return chosenTemplateName
	}
}
