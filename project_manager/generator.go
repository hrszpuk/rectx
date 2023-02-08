package project_manager

import (
	"fmt"
	"os"
	"os/exec"
	"rectx/licenses"
	projectConfig "rectx/project_manager/config"
	"rectx/templates"
	"rectx/utilities"
	"strconv"
	"strings"
	"time"
)

func CreateNewProject(config *projectConfig.ProjectConfig) {
	fmt.Print("Generating project... ")

	variables := make(map[string]string)
	variables["%PROJECT_NAME%"] = config.Project.Name
	variables["%AUTHOR%"] = config.Project.Authors[0]
	year, month, day := time.Now().Date()
	variables["%YEAR%"] = strconv.Itoa(year)
	variables["%MONTH%"] = month.String()
	variables["%DAY%"] = strconv.Itoa(day)

	if file, err := os.Stat(config.Project.Name); err == nil {
		tyype := ""
		if file.IsDir() {
			tyype = "directory"
		} else {
			tyype = "file"
		}

		answer := ""
		fmt.Printf("A %s by the name of \"%s\" already exists would you like to overwrite it? [Y/n]: ", tyype, config.Project.Name)
		fmt.Scanln(&answer)

		for _, item := range []string{"Y", "YES", "YEAH", "MHM", "SURE", "1", "TRUE"} {
			if answer == item || answer == strings.ToLower(item) {
				utilities.Check(os.RemoveAll(config.Project.Name), true, "Could not remove pre-existing project directory in attempt to overwrite!")
			}
		}
	}
	utilities.Check(os.Mkdir(config.Project.Name, 0750), true, "Attempt to create project directory failed... Permission levels may not be sufficient?")

	f, err := os.ReadFile(utilities.GetRectxPath() + "/templates/" + config.Project.Template)
	if os.IsNotExist(err) {
		utilities.Check(err, true, "Attempt to read template file failed because it doesn't exist... What?")
	} else {
		utilities.Check(err, true, "Attempt to read template file failed.")
	}

	if config.Project.License != "None" {
		licenses.GenerateLicense(config.Project.License, variables)
	}

	parser := templates.NewTemplateParser(string(f))
	statements := parser.Parse()

	_ = exec.Command("cd", config.Project.Name)
	for _, statement := range statements {
		statement.Generate(config.Project.Name)
	}

	fmt.Println("Done!")
}
