package main

import (
	"log"
	"os"
	"rectx/utilities"
)

func main() {
	if len(os.Args) < 2 {
		ShowUsage()
		os.Exit(0)
	}

	initFlags()

	switch os.Args[1] {
	case "new":
		utilities.Check(newCmd.Parse(os.Args[2:]))
		if help {
			ShowNewHelpMenu()
		}
	case "build":
		utilities.Check(buildCmd.Parse(os.Args[2:]))
		if help {
			ShowBuildHelpMenu()
		}

	case "run":
		utilities.Check(runCmd.Parse(os.Args[2:]))
		if help {
			ShowRunHelpMenu()
		}

	case "template":
		utilities.Check(templateCmd.Parse(os.Args[2:]))
		if help {
			ShowTemplateHelpMenu()
		}

	case "config":
		utilities.Check(configCmd.Parse(os.Args[2:]))
		if help {
			ShowConfigHelpMenu()
		}

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
