package main

import (
	"log"
	"os"
	"rectx/project_manager"
)

func main() {
	initFlags()

	if len(os.Args) < 2 {
		ShowHelpMenu()
		os.Exit(0)
	}

	switch os.Args[1] {
	case "new":
		handleParseErrorAndHelpFlag(newCmd, newCmd.Parse(os.Args[2:]), ShowNewHelpMenu)
		project_manager.New()
	case "build":
		handleParseErrorAndHelpFlag(buildCmd, buildCmd.Parse(os.Args[2:]), ShowBuildHelpMenu)
	case "run":
		handleParseErrorAndHelpFlag(runCmd, runCmd.Parse(os.Args[2:]), ShowRunHelpMenu)
	case "template":
		handleParseErrorAndHelpFlag(templateCmd, templateCmd.Parse(os.Args[2:]), ShowTemplateHelpMenu)
	case "config":
		handleParseErrorAndHelpFlag(configCmd, configCmd.Parse(os.Args[2:]), ShowConfigHelpMenu)
	case "help":
		fallthrough
	case "--help":
		fallthrough
	case "-h":
		ShowHelpMenu()
	default:
		log.Fatalf("Unknown command \"%s\"! If you're looking for a certain command try \"rectx --help\"!\n", os.Args[1])
	}
}
