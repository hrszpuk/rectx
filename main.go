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

	ifHelp := func(showMenu bool, helpMenu func()) {
		if showMenu {
			helpMenu()
		}
	}

	switch os.Args[1] {
	case "new":
		utilities.Check(newCmd.Parse(os.Args[2:]))
		ifHelp(help, ShowNewHelpMenu)
	case "build":
		utilities.Check(buildCmd.Parse(os.Args[2:]))

	case "run":
		utilities.Check(runCmd.Parse(os.Args[2:]))

	case "template":
		utilities.Check(templateCmd.Parse(os.Args[2:]))

	case "config":
		utilities.Check(configCmd.Parse(os.Args[2:]))

	case "help":
		fallthrough
	case "--help":
		fallthrough
	case "-h":
		ShowHelpMenu()

	default:
		log.Fatalf("Unknown command \"%s\"!\n", os.Args[1])
	}
}
