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

	if _, err := os.Stat(buildPath + "/" + name); os.IsExist(err) {
		utilities.Check(exec.Command("rm", buildPath+"/"+name, "-f").Run(), false, "Executable exists but was unable to remove it from build path!")

	}

	dir, err := os.ReadDir(sourcePath)
	msg := ""
	if sourcePath != "" {
		msg = fmt.Sprintf("Could not build %s directory!", sourcePath)
	} else {
		msg = "Source directory has not been set. Cannot build. Please check your project.rectx file."
	}
	utilities.Check(err, true, msg)

	if len(dir) < 1 {
		fmt.Printf("Found no files in %s!\n", sourcePath)
		os.Exit(1)
	}

	_, err = os.Stat(sourcePath + "/" + main)
	msg = fmt.Sprintf("Attempted to build %s but failed for unknown reasons!", sourcePath+"/"+main)
	utilities.Check(err, true, msg)

	utilities.Check(exec.Command(compiler, main, "-o="+buildPath+"/"+main).Run(), true, "Failed to build for unknown reasons!")
}
