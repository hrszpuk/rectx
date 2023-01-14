package main

import (
	"fmt"
	"log"
	"os"
	"rectx/utilities"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: rectx <command> [flags] [arguments]")
		os.Exit(0)
	}

	initFlags()

	switch os.Args[1] {
	case "new":
		utilities.Check(newCmd.Parse(os.Args[2:]))
		ShowHelpMenu(help)

	case "build":
		utilities.Check(buildCmd.Parse(os.Args[2:]))
		ShowHelpMenu(help)

	case "run":
		utilities.Check(runCmd.Parse(os.Args[2:]))
		ShowHelpMenu(help)

	case "template":
		utilities.Check(templateCmd.Parse(os.Args[2:]))
		ShowHelpMenu(help)

	case "config":
		utilities.Check(configCmd.Parse(os.Args[2:]))
		ShowHelpMenu(help)

	case "help":
		fallthrough
	case "--help":
		fallthrough
	case "-h":
		ShowHelpMenu(true)

	default:
		log.Fatalf("Unknown command \"%s\"!\n", os.Args[1])
	}
}
