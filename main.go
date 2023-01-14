package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"rectx/utilities"
)

var (
	// rectx new [optional]
	newCmd        = flag.NewFlagSet("new", flag.ExitOnError)
	projectName   string // -n --name
	author        string // -a --author
	template      string // -t --template
	path          string // -p --path
	license       string // -l --license
	version       string // -v --version
	versionSystem string // -vs --version-system
	noPrompt      bool   // -np --no-prompt

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
	all               bool // -a --author
)

func initFlags() {
	// TODO
	newCmd.StringVar(&projectName, "name", "Untitled", "specify what you want your project to be called")
	newCmd.StringVar(&author, "author", "", "specify who is creating the project")
	newCmd.StringVar(&template, "template", "default", "usage")
	newCmd.StringVar(&path, "path", "Untitled", "usage")
	newCmd.StringVar(&license, "license", "", "usage")
	newCmd.StringVar(&version, "version", "0.1.0", "usage")
	newCmd.StringVar(&versionSystem, "version-system", "major.minor.patch", "usage")
	newCmd.BoolVar(&noPrompt, "noPrompt", false, "usage")
}

func main() {
	if len(os.Args) < 3 {
		fmt.Println("Usage: rectx <command> [flags] [arguments]")
		os.Exit(0)
	}

	switch os.Args[1] {
	case "new":
		utilities.Check(newCmd.Parse(os.Args[2:]))

	case "build":
		utilities.Check(buildCmd.Parse(os.Args[2:]))

	case "run":
		utilities.Check(runCmd.Parse(os.Args[2:]))

	case "template":
		utilities.Check(templateCmd.Parse(os.Args[2:]))

	case "config":
		utilities.Check(configCmd.Parse(os.Args[2:]))

	default:
		log.Fatalf("Unknown subcommand \"%s\"! Maybe try rectx --help if you're looking for a certain command!\n", os.Args[1])
	}
}
