package main

import (
	"fmt"
	"log"
	"os"
	"rectx/config"
	"rectx/project_manager"
	"rectx/templates"
	"strings"
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
		project_manager.Build()
	case "run":
		handleParseErrorAndHelpFlag(runCmd, runCmd.Parse(os.Args[2:]), ShowRunHelpMenu)
		project_manager.Run()
	case "templates":
		fallthrough
	case "template":
		handleParseErrorAndHelpFlag(templateCmd, templateCmd.Parse(os.Args[2:]), ShowTemplateHelpMenu)
		if len(os.Args) == 3 {
			if os.Args[2] == "list" {
				fmt.Println("Listing all templates found:")
				for i, files := range templates.ListTemplates() {
					fmt.Printf("%d. %s\n", i+1, files)
				}
			}
		} else if len(os.Args) == 4 {
			switch os.Args[2] {
			case "add":
				if !strings.HasSuffix(os.Args[3], ".rectx.template") {
					fmt.Printf("Whoops \"%s\" isn't a .rectx.template file!\n", os.Args[3])
					os.Exit(1)
				} else {
					config.AddTemplate(os.Args[3])
				}
			case "test":
				templates.Test(os.Args[3])
			default:
				fmt.Printf("Unknown subcommand \"%s\"! Maybe try rectx templates --help for a list of subcommands.", os.Args[3])
			}
		} else if len(os.Args) == 5 {
			switch os.Args[2] {
			case "rename":
				config.RenameTemplate(os.Args[3], os.Args[4])
			}
		}
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
