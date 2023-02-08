package projectManager

import (
	"fmt"
	"os"
	"os/exec"
	ProjectConfig "rectx/projectManager/ProjectConfig"
	"rectx/utilities"
)

func Run() {
	conf := ProjectConfig.CreateDefaultConfig()

	// TODO this error check should be moved into Load() later
	if _, err := os.Stat("project.rectx"); os.IsNotExist(err) {
		utilities.Check(err, true, "Could not find project.rectx ProjectConfig within current directory!")
		fmt.Println("Could not find project.rectx ProjectConfig within current directory!")
	}

	conf.Load("project.rectx")
	buildPath := conf.BuildProfile.BuildDirectory
	name := conf.BuildProfile.ExecutableName

	if _, err := os.Stat(buildPath + "/" + name); os.IsNotExist(err) {
		Build()
	}

	// TODO add executable arguments to run profile
	utilities.Check(exec.Command("./"+buildPath+name).Run(), true, "Attempt to run executable failed!")
}
