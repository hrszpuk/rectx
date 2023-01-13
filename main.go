package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"rectx/utilities"
)

var (
	newCmd      = flag.NewFlagSet("new", flag.ExitOnError)
	buildCmd    = flag.NewFlagSet("build", flag.ExitOnError)
	runCmd      = flag.NewFlagSet("run", flag.ExitOnError)
	templateCmd = flag.NewFlagSet("template", flag.ExitOnError)
	configCmd   = flag.NewFlagSet("config", flag.ExitOnError)
)

func main() {
	if len(os.Args) < 3 {
		fmt.Println("Usage: rectx <command> [flags] [arguments]")
		os.Exit(0)
	}

	switch os.Args[1] {
	case "new":
		utilities.Check(newCmd.Parse(os.Args[2:]))
		// project_manager.new
		// -n --name "project name"
		// -a --author "author"
		// -t --template "default"
		// -d --directory "create project path"
		// -l --license "obvious"
		// -v --version "0.1.0"
		// -vs --version-system "major.minor.patch"
		// -np --no-prompt
	case "build":
		utilities.Check(buildCmd.Parse(os.Args[2:]))
		// project_manager.build
		// -p --profile "custom build profile"
	case "run":
		utilities.Check(runCmd.Parse(os.Args[2:]))
		// project_manager.run
		// -p --profile "custom run profile"
	case "template":
		utilities.Check(templateCmd.Parse(os.Args[2:]))
		// templates. ...
		// list
		// add "path/to/file" "name"
		// snapshot "path/to/folder" "name"
		// setDefault "name of template"
		// rename "name" "new name"
	case "config":
		utilities.Check(configCmd.Parse(os.Args[2:]))
		// config. ...
		// validate
		// -cf --config-file
		// -t --templates
		// -l --licenses
		// regenerate
		// -cf --config-file
		// -t --templates
		// -l --licenses
		// reset "config.setting" "value"
		// --all
		// set "config.setting" "value"
	default:
		log.Fatalf("Unknown subcommand \"%s\"!\n", os.Args[1])
	}
}
