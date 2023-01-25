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
		EnsureArguments(3, "template")
		switch os.Args[2] {
		case "list":
			fmt.Println("Listing all templates found:")
			for i, files := range templates.ListTemplates() {
				fmt.Printf("%d. %s\n", i+1, files)
			}
		case "add":
			EnsureArguments(4, "template add")
			// TODO Maybe check if file exists first?
			if !strings.HasSuffix(os.Args[3], ".rectx.template") {
				fmt.Printf("Whoops \"%s\" isn't a .rectx.template file!\n", os.Args[3])
				os.Exit(1)
			} else {
				config.AddTemplate(os.Args[3])
			}
		case "test":
			EnsureArguments(4, "template test")
			templates.Test(os.Args[3])
		case "default":
			EnsureArguments(4, "template default")
			config.SetDefaultTemplate(os.Args[3])
		case "snapshot":
			EnsureArguments(4, "template snapshot")
			templates.Snapshot(os.Args[3])
		case "rename":
			EnsureArguments(5, "template rename")
			config.RenameTemplate(os.Args[3], os.Args[4])
		default:
			fmt.Printf("Unknown subcommand \"%s\"! Maybe try rectx templates --help for a list of subcommands...\n", os.Args[2])
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

// This function is used to check if the required number of arguments is met. 
// If so, the function will do nothing, otherwise it will display an error message, and a usage for the command provided.
// If only a command that requires a subcommand is supplied (via command argument) then a <subcommand> will be displayed
// in the usage message. Otherwise, only <arg> symbols will be displayed.
func EnsureArguments(requiredArgumentCount int, command string) {

	if requiredArgumentCount > len(os.Args) {
		argumentCount := requiredArgumentCount - len(os.Args)
		displayUsageWithSubcommand := false

		fmt.Printf("Not enough arguments to run \"rectx %s\"!\n", command)
		fmt.Print("Usage: rectx ")

		// We print a <subcommand> only if command is a command that requires subcommands.
		// Otheriwse, we can just display the actual command + subcommand used.
		for _, COMMAND := range []string{"template", "templates", "config"} {
			if command == COMMAND {
				fmt.Print(command + " <subcommand> ")
				displayUsageWithSubcommand = true
			}
		}
		if displayUsageWithSubcommand {
			argumentCount--
		} else {
			fmt.Print(command)
		}

		fmt.Printf(" %s\n", strings.Repeat("<arg> ", argumentCount))
		os.Exit(1)
	}
}
