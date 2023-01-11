package main

import (
	"fmt"
	"os"
	"rectx/config"
)

func main() {
	arguments := os.Args
	arguments = arguments[1:]
	if len(arguments) < 1 {
		fmt.Println("Usage: rectx <command> [arguments]")
		return
	} else if arguments[0] == "new" {
		fmt.Println("Creating a new project")
		var name string
		var authors string

		fmt.Print("Project name: ")
		fmt.Scanln(&name)

		fmt.Print("Project authors (space separated): ")
		fmt.Scanln(&authors)

		fmt.Println("Creating project directory")
		os.Mkdir(name, os.ModePerm)
		os.Mkdir(fmt.Sprintf("%s/Source", name), os.ModePerm)
		os.Mkdir(fmt.Sprintf("%s/Build", name), os.ModePerm)
		os.Mkdir(fmt.Sprintf("%s/Resources", name), os.ModePerm)
		f, _ := os.Create(fmt.Sprintf("%s/README.md", name))
		f.Close()
		f, _ = os.Create(fmt.Sprintf("%s/Source/main.rct", name))
		f.Close()
		f, _ = os.Create(fmt.Sprintf("%s/.gitignore", name))
		f.Close()
		f, _ = os.Create(fmt.Sprintf("%s/project.rectx", name))
		f.Close()
	} else if arguments[0] == "validate" {
		config.ValidateConfigFile()
	}
}
