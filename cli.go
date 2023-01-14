package main

import (
	"flag"
	"fmt"
)

var (
	// rectx new [optional]
	newCmd      = flag.NewFlagSet("new", flag.ExitOnError)
	projectName string // -n --name
	author      string // -a --author
	template    string // -t --template
	path        string // -p --path
	license     string // -l --license
	version     string // -v --version
	noPrompt    bool   // -np --no-prompt

	// rectx build [optional]
	buildCmd     = flag.NewFlagSet("build", flag.ExitOnError)
	buildProfile string // -p --profile

	// rectx run [optional]
	runCmd     = flag.NewFlagSet("run", flag.ExitOnError)
	runProfile string // -p --profile

	// rectx template <subcommand> [optional]
	templateCmd               = flag.NewFlagSet("template", flag.ExitOnError)
	templateSubcommands       = [...]string{"list", "add", "snapshot", "setDefault", "rename"}
	templateSubcommandDetails = [...]string{
		"Lists all the templates in the rectx config directory",
		"Adds a new template file (.rectx.template required)",
		"Reads the folders/files of a directory and generates a .rectx.template file",
		"Set a default template that will be auto selected for your projects",
		"Change the name of a template",
	}

	// rectx config <subcommand> [optional]
	configCmd               = flag.NewFlagSet("config", flag.ExitOnError)
	configSubcommands       = [...]string{"validate", "regenerate", "reset", "set"}
	configSubcommandDetails = [...]string{
		"Checks rectx config data does not contain any errors",
		"Downloads any missing rectx config data",
		"Reset a value to it's default in the rectx global config",
		"Change a value in the rectx global config",
	}
	configFile bool // -c --config
	templates  bool // -t --templates
	licenses   bool // -l --licenses
	all        bool // -a --all

	help bool

	CMDS = [...]*flag.FlagSet{newCmd, buildCmd, runCmd, templateCmd, configCmd}
)

func initFlags() {
	initNewFlags()
	initBuildFlags()
	initRunFlags()
	initConfigFlags()

	for _, cmd := range CMDS {
		cmd.BoolVar(&help, "help", false, "shows a help message with a list of all the commands")
	}
}

func initNewFlags() {
	newCmd.StringVar(&projectName, "name", "Untitled", "specify what you want your project to be called")
	newCmd.StringVar(&author, "author", "", "specify who is creating the project")
	newCmd.StringVar(&template, "template", "default", "specify how you want the project sturcture to look")
	newCmd.StringVar(&path, "path", "Untitled", "specify where to put the project")
	newCmd.StringVar(&license, "license", "", "specify which license you want your project to use")
	newCmd.StringVar(&version, "version", "0.1.0", "specify what version to start the project at")
	newCmd.BoolVar(&noPrompt, "noPrompt", false, "don't show the project prompt (generate based off defaults and provided flags)")
}

func initConfigFlags() {
	configCmd.BoolVar(&configFile, "config", false, "specifically the rectx config file")
	configCmd.BoolVar(&templates, "templates", false, "specifically the rectx templates")
	configCmd.BoolVar(&licenses, "licenses", false, "specifically the rectx licenses")
	configCmd.BoolVar(&all, "all", false, "specifically validate/regenerate the entire rectx config directory")
}

func initBuildFlags() {
	buildCmd.StringVar(&buildProfile, "profile", "", "specify a custom build profile for the project")
}

func initRunFlags() {
	runCmd.StringVar(&runProfile, "profile", "", "specify a custom run profile for the project")
}

func ShowUsage() {
	fmt.Printf("Usage: rectx <command> [subcommand] [flags] [arguments]\n")
}

func ShowHelpMenu() {
	fmt.Println()
	ShowUsage()
	for _, command := range CMDS {
		name := command.Name()
		if name == "template" || name == "config" {
			name += " [subcommand]"
		}
		fmt.Printf("\nrectx %s [flags] [arguments]\n\n", name)
		if name == "template [subcommand]" {
			fmt.Printf("  [subcommands]\n")
			for i, c := range templateSubcommands {
				fmt.Printf("   %s\n", c)
				fmt.Printf("         %s\n", templateSubcommandDetails[i])
			}
		} else if name == "config [subcommand]" {
			fmt.Printf("  [subcommands]\n")
			for i, c := range configSubcommands {
				fmt.Printf("  %s\n", c)
				fmt.Printf("         %s\n", configSubcommandDetails[i])
			}
		}
		fmt.Printf("  [flags]\n")
		command.PrintDefaults()
	}

}

func ShowNewHelpMenu() {
	fmt.Printf("\nUsage: rectx new [flags] [arguments]\n\n")
	fmt.Printf(
		"  [details]\nUsed to create a new project! This command will prompt you questions about your project" +
			"and then generate all the project files you need to get started." +
			"Optionally, you can pass flags such as --name=\"borgor\" to quickly assign values without the prompt!" +
			"You can also add default values for author, license, template, which will make the prompt much faster to fill out!\n\n")
	fmt.Printf("  [flags]\n")
	newCmd.PrintDefaults()
}
