package main

import (
	"fmt"
	"log"
	"os"
	"rectx/config"
	"rectx/projectManager"
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

	// rectx new|build|run
	// These commands have no subcommands and their flags are handled in the projectManager module.
	case "new":
		handleParseErrorAndHelpFlag(newCmd, newCmd.Parse(os.Args[2:]), ShowNewHelpMenu)
		projectManager.New()
	case "build":
		handleParseErrorAndHelpFlag(buildCmd, buildCmd.Parse(os.Args[2:]), ShowBuildHelpMenu)
		projectManager.Build()
	case "run":
		handleParseErrorAndHelpFlag(runCmd, runCmd.Parse(os.Args[2:]), ShowRunHelpMenu)
		projectManager.Run()

	// rectx template|templates
	// I allow both "templates" and "template" because both go well with the subcommands and I kept putting the wrong one.
	case "templates":
		fallthrough
	case "template":
		// TODO debug flag not working for template subcommands - Issue should probably be created.
		handleParseErrorAndHelpFlag(templateCmd, templateCmd.Parse(os.Args[2:]), ShowTemplateHelpMenu)
		EnsureArguments(3, "template")

		switch os.Args[2] {

		// `rectx template list` lists all the templates in the rectx template directory
		case "list":
			fmt.Println("Listing all templates found:")
			for i, files := range templates.ListTemplates() {
				fmt.Printf("%d. %s\n", i+1, files)
			}

		// `rectx template add <path/to/template>` copies a .rectx.template into the rectx template directory
		// Although this is a template subcommand, it uses the ProjectConfig package because it manages the ~/.rectx ProjectConfig directory where the templates are stored.
		case "add":
			EnsureArguments(4, "template add")
			config.AddTemplate(os.Args[3])

		// `rectx template test <template>` parses and generates using a template file in a temporary directory.
		// This command was designed to allow you to test for errors in a template file. Errors are reported to the command line.
		case "test":
			EnsureArguments(4, "template test")
			templates.Test(os.Args[3])

		// `rectx template default <template>` sets the provided template as the default template for project generation.
		// This means if you don't select a template this template will be used. The default template is `default.rect.template`.
		// Although this is a template subcommand, it uses the ProjectConfig package because that package manages the rectx ProjectConfig where the default template is stored.
		case "default":
			EnsureArguments(4, "template default")
			config.SetDefaultTemplate(os.Args[3])

		// `rectx template snapshot <path>` generates a .rectx.template file from the information provided in the directory provided.
		// NOTE: the template name will be taken from the directory name, commands should be held within a file called "commands"
		case "snapshot":
			// TODO maybe an exclude flag should be added so people can avoid files named "commands" that have other purposes?
			EnsureArguments(4, "template snapshot")
			templates.Snapshot(os.Args[3])

		// `rectx template rename <templateName> <newTemplateName>`
		case "rename":
			EnsureArguments(5, "template rename")
			config.RenameTemplate(os.Args[3], os.Args[4])
		default:
			fmt.Printf("Unknown subcommand \"%s\"! Maybe try rectx templates --help for a list of subcommands...\n", os.Args[2])
		}

	// rectx ProjectConfig
	// TODO I should probably start this before it's too late - Tokorv
	case "ProjectConfig":
		handleParseErrorAndHelpFlag(configCmd, configCmd.Parse(os.Args[2:]), ShowConfigHelpMenu)

	case "help":
		fallthrough
	case "--help":
		fallthrough
	case "-h":
		ShowHelpMenu()
	default:
		log.Fatalf("Unknown command \"%s\"!\nIf you're looking for a certain command try \"rectx --help\"!\n", os.Args[1])
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
		for _, COMMAND := range []string{"template", "templates", "ProjectConfig"} {
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
