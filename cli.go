package main

import (
	"flag"
	"os"
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
	templateCmd         = flag.NewFlagSet("template", flag.ExitOnError)
	templateSubcommands = [...]string{"list", "add", "snapshot", "setDefault", "rename"}

	// rectx config <subcommand> [optional]
	configCmd         = flag.NewFlagSet("config", flag.ExitOnError)
	configSubcommands = [...]string{"validate", "regenerate", "reset", "set"}
	configFile        bool // -c --config
	templates         bool // -t --templates
	licenses          bool // -l --licenses
	all               bool // -a --all

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

func ShowHelpMenu(visible bool) {
	if visible {
		for _, cmd := range CMDS {
			cmd.PrintDefaults()
		}
		os.Exit(0)
	}
}
