package projectManager

import (
	"fmt"
	"os"
	"os/exec"
	ProjectConfig "rectx/projectManager/config"
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
		utilities.Check(exec.Command("rm", buildPath+"/"+name, "-f").Run(), false, "Attempted to add ")

	}

	dir, err := os.ReadDir(sourcePath)
	utilities.ErrCheckReadDir(err, sourcePath, func() {
		if dir, err := os.Stat(sourcePath); !dir.IsDir() {
			utilities.Check(err, true, fmt.Sprintf("Could not build \"%s\" because it is not a directory!", sourcePath))
		} else if os.IsNotExist(err) {
			utilities.Check(err, true, fmt.Sprintf("Could not build \"%s\" because it does not exist!", sourcePath))
		} else if os.IsPermission(err) {
			utilities.Check(err, true, fmt.Sprintf("Could not build \"%s\" due to a lack of permissions!", sourcePath))
		}
	})

	if len(dir) < 1 {
		fmt.Printf("Found no files in %s!\n", sourcePath)
		os.Exit(1)
	}

	_, err = os.Stat(sourcePath + "/" + main)
	utilities.Check(err, true, "Attempted to build %s but failed for unknown reasons!")

	utilities.Check(exec.Command(compiler, main, "-o="+buildPath+"/"+main).Run(), true, "Failed to build for unknown reasons!")
}
