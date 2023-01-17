package project_manager

import (
	"fmt"
	"os"
	"os/exec"
	ProjectConfig "rectx/project_manager/config"
	"rectx/utilities"
)

func Build() {
	conf := ProjectConfig.CreateDefaultConfig()

	// TODO this error check should be moved into Load() later
	if _, err := os.Stat("project.rectx"); os.IsNotExist(err) {
		fmt.Println("Could not find project.rectx config within current directory!")
		os.Exit(1)
	}

	conf.Load("project.rectx")
	sourcePath := conf.BuildProfile.BuildDirectory
	buildPath := conf.BuildProfile.SourceDirectory
	name := conf.BuildProfile.ExecutableName
	compiler := conf.BuildProfile.Compiler
	main := "main.rct" // TODO add entry source file to build profile

	if _, err := os.Stat(buildPath + "/" + name); !os.IsNotExist(err) {
		utilities.Check(exec.Command("rm", buildPath+"/"+name, "-f").Run())

	}

	dir, err := os.ReadDir(sourcePath)
	utilities.Check(err)

	if len(dir) < 1 {
		fmt.Printf("No files within %s/!\n", sourcePath)
		os.Exit(1)
	}

	if _, err := os.Stat(sourcePath + "/" + main); !os.IsNotExist(err) {
		fmt.Printf("Could not find main.rct in %s/!\n", sourcePath)
		os.Exit(1)
	}

	utilities.Check(exec.Command(compiler, main, "-o="+buildPath+"/"+main).Run())
}
