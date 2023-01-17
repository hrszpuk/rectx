package project_manager

import (
	"fmt"
	"os"
	"os/exec"
	ProjectConfig "rectx/project_manager/config"
	"rectx/utilities"
)

func Run() {
	conf := ProjectConfig.CreateDefaultConfig()

	// TODO this error check should be moved into Load() later
	if _, err := os.Stat("project.rectx"); os.IsNotExist(err) {
		fmt.Println("Could not find project.rectx config within current directory!")
		os.Exit(1)
	}

	conf.Load("project.rectx")
	buildPath := conf.BuildProfile.SourceDirectory
	name := conf.BuildProfile.ExecutableName

	if _, err := os.Stat(buildPath + "/" + name); os.IsNotExist(err) {
		Build()
	}

	// TODO add executable arguments to run profile
	utilities.Check(exec.Command(name).Run())
}
