package project_manager

import (
	"fmt"
	"rectx/licenses"
	projectConfig "rectx/project_manager/config"
	"rectx/templates"
	"strconv"
	"strings"
	"time"
)

func New() {
	pc := projectConfig.CreateDefaultConfig()

	fmt.Print("Project name: ")
	fmt.Scanln(&pc.Project.Name)

	var author string
	fmt.Print("Author: ")
	fmt.Scanln(&author)
	pc.Project.Authors = append(pc.Project.Authors, author)

	pc.Project.Version = GetVersion()
	pc.Project.License = licenses.Prompt()
	pc.Project.Template = GetTemplate()

	variables := make(map[string]string)
	variables["%PROJECT_NAME%"] = pc.Project.Name
	variables["%AUTHOR%"] = pc.Project.Authors[0]
	year, month, day := time.Now().Date()
	variables["%YEAR%"] = strconv.Itoa(year)
	variables["%MONTH%"] = month.String()
	variables["%DAY%"] = strconv.Itoa(day)

	CreateNewProject(pc, variables)

	pc.Dump(pc.Project.Name + "/project.rectx")
}

func GetTemplate() string {
	fmt.Println("Pick a template:")
	var templateList = templates.FetchTemplates()
	for _, name := range templateList {
		fmt.Printf("- %s\n", strings.Replace(name, ".rectx.template", "", 1))
	}

	var chosenTemplateName string
	validTemplateChosen := false
	fmt.Println("\nType the name of the template you want to use: ")
	fmt.Scanln(&chosenTemplateName)
	chosenTemplateName += ".rectx.template"
	for _, name := range templateList {
		if chosenTemplateName == name {
			validTemplateChosen = true
			break
		}
	}

	if !validTemplateChosen {
		fmt.Printf("\"%s\" is not a valid template!\n", chosenTemplateName)
		return GetTemplate()
	} else {
		return chosenTemplateName
	}
}

func GetVersion() string {
	version := ""
	fmt.Print("Version: ")
	fmt.Scanln(&version)
	badCharacters := "abcdefghijklmnopqrstuvwxyz"
	badCharacters += strings.ToUpper(badCharacters)
	badCharacters += "!£$%^&*()\"'@~#[]{}:;<>,-=+_/?|\\`¬"
	if strings.ContainsAny(version, badCharacters) {
		fmt.Println("Warning: Version number contained bad characters! " +
			"Version numbers can only contain numbers (0-9) and dots (.). " +
			"\nFor Example: 1.2.3, 0.0.1, 8.394.13305" +
			"\nPlease submit another version number")
		return GetVersion()
	} else {
		return version
	}
}
