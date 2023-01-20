package main

import (
	"flag"
	"fmt"
)

var (
	// rectx new [optional]
	newCmd          = flag.NewFlagSet("new", flag.ExitOnError)
	projectNameFlag string // -n --name
	authorFlag      string // -a --author
	templateFlag    string // -t --template
	pathFlag        string // -p --path
	licenseFlag     string // -l --license
	versionFlag     string // -v --version
	noPromptFlag    bool   // -np --no-prompt

	// rectx build [optional]
	buildCmd         = flag.NewFlagSet("build", flag.ExitOnError)
	buildProfileFlag string // -p --profile

	// rectx run [optional]
	runCmd         = flag.NewFlagSet("run", flag.ExitOnError)
	runProfileFlag string // -p --profile

	// rectx template <subcommand> [optional]
	templateCmd               = flag.NewFlagSet("template", flag.ExitOnError)
	templateSubcommands       = []string{"list", "add", "snapshot", "setDefault", "rename"}
	templateSubcommandDetails = []string{
		"Lists all the templates in the rectx config directory.",
		"Adds a new template file (.rectx.template required).",
		"Reads the folders/files of a directory and generates a .rectx.template file.",
		"Set a default template that will be auto selected for your projects.",
		"Change the name of a template.",
	}
	templateSubcommandArguments = []string{
		"(None)", "path/to/file", "path/to/folder", "template-name", "template-name new-template-name",
	}

	// rectx config <subcommand> [optional]
	configCmd               = flag.NewFlagSet("config", flag.ExitOnError)
	configSubcommands       = []string{"validate", "regenerate", "reset", "set"}
	configSubcommandDetails = []string{
		"Checks rectx global config data does not contain any errors.",
		"Downloads any missing rectx config data.",
		"Reset a value to it's default in the rectx global config.",
		"Change a value in the rectx global config.",
	}
	configSubcommandArguments = []string{
		"(None)", "(None)", "key-name", "key-name new-key-value",
	}
	configFileFlag bool // -c --config
	templatesFlag  bool // -t --templates
	licensesFlag   bool // -l --licenses
	allFlag        bool // -a --all

	helpFlag bool

	CMDS = [...]*flag.FlagSet{newCmd, buildCmd, runCmd, templateCmd, configCmd}

	// mega mega mega
	megaSpace = "       "
)

func initFlags() {
	initNewFlags()
	initBuildFlags()
	initRunFlags()
	initConfigFlags()

	for _, cmd := range CMDS {
		cmd.BoolVar(&helpFlag, "help", false, "Shows a specific help message for the command used.")
	}
}

func initNewFlags() {
	newCmd.StringVar(&projectNameFlag, "name", "Untitled", "Specify what you want your project to be called.")
	newCmd.StringVar(&authorFlag, "author", "", "Specify who is creating the project.")
	newCmd.StringVar(&templateFlag, "template", "default", "Specify how you want the project structure to look.")
	newCmd.StringVar(&pathFlag, "path", "Untitled", "Specify where to put the project.")
	newCmd.StringVar(&licenseFlag, "license", "", "Specify which license you want your project to use.")
	newCmd.StringVar(&versionFlag, "version", "0.1.0", "Specify what version to start the project at.")
	newCmd.BoolVar(&noPromptFlag, "noPrompt", false, "Don't show the project prompt (generate based off defaults and provided flags).")
}

func initConfigFlags() {
	configCmd.BoolVar(&configFileFlag, "config", false, "Specifies the rectx config file specifically.")
	configCmd.BoolVar(&templatesFlag, "templates", false, "Specifies the rectx templates specifically.")
	configCmd.BoolVar(&licensesFlag, "licenses", false, "Specifies the rectx licenses specifically.")
	configCmd.BoolVar(&allFlag, "all", false, "Specifically validate/regenerate the entire rectx config directory.")
}

func initBuildFlags() {
	buildCmd.StringVar(&buildProfileFlag, "profile", "", "Specify a custom build profile for the project (must be declared in the project.rectx).")
}

func initRunFlags() {
	runCmd.StringVar(&runProfileFlag, "profile", "", "Specify a custom run profile for the project (must be declared in the project.rectx).")
}

func ShowUsage(command string, showSubcommands bool) {
	if showSubcommands {
		command += " [subcommand]"
	}
	fmt.Printf("\n  Usage: rectx %s [flags] [arguments]\n", command)
}

func ShowHelpMenu() {
	ShowUsage("<command>", true)
	for _, command := range CMDS {
		name := command.Name()
		if name == "template" || name == "config" {
			name += " [subcommand]"
		}
		fmt.Printf("\n> rectx %s [flags] [arguments]\n\n", name)
		if name == "template [subcommand]" {
			fmt.Printf("  [subcommands]\n")
			for i, c := range templateSubcommands {
				fmt.Printf("   %s\n", c)
				fmt.Printf("         %s\n", templateSubcommandDetails[i])
			}
			fmt.Println()
		} else if name == "config [subcommand]" {
			fmt.Printf("  [subcommands]\n")
			for i, c := range configSubcommands {
				fmt.Printf("  %s\n", c)
				fmt.Printf("         %s\n", configSubcommandDetails[i])
			}
			fmt.Println()
		}
		fmt.Println("  [flags]")
		command.PrintDefaults()
	}

}

func ShowNewHelpMenu() {
	ShowUsage("new", false)
	fmt.Printf(
		"\n  [details]\n  Used to create a new project! \n  This command will prompt you questions about your project" +
			"and then generate all the project files you need to get started." +
			"\n  Optionally, you can pass flags such as --name=\"borgor\" to quickly assign values without the prompt!" +
			"\n  You can also add default values for authorFlag, license, template, which will make the prompt much faster to fill out!\n\n")
	fmt.Println("  [flags]")
	newCmd.PrintDefaults()
}

func ShowRunHelpMenu() {
	ShowUsage("run", false)
	fmt.Printf("\n  [details]\n  Runs the project's source code. \n  This function will run the executable found in the project." +
		"\n  If no executable exists, or if edits have been made to the project's source code since the last build, " +
		"\n  then this command will automatically build/re-build the executable.\n\n")
	fmt.Println("  [flags]")
	runCmd.PrintDefaults()
}

func ShowBuildHelpMenu() {
	ShowUsage("build", false)
	fmt.Printf("\n  [details]\n  Builds the project's source code. \n  This function will build an executable but will not run the executable." +
		"\n  If you wish to run the executable then please check out \"rectx run --help\".\n\n")
	fmt.Println("  [flags]")
	buildCmd.PrintDefaults()
}

func ShowTemplateHelpMenu() {
	ShowUsage("template", true)
	fmt.Printf("\n  [details]\n  Manage project generation templates." +
		"\n  These templates describe how a project is to be generated." +
		"\n  You can easily create you own template using a simple syntax." +
		"\n  Defining your own template allows you to work on projects with a structure you enjoy and are familiar with!\n\n")
	PrintSubcommands(templateSubcommands, templateSubcommandDetails, templateSubcommandArguments)
	fmt.Printf("\n  [flags]\n")
	templateCmd.PrintDefaults()
}

func ShowConfigHelpMenu() {
	ShowUsage("config", true)
	fmt.Printf("\n  [details]\n  Manage the rectx global config." +
		"\n  This subcommands has little to do with project.rectx files." +
		"\n  This is for the global rectx config used to configure rectx itself." +
		"\n  Most importantly, the config subcommands can be used to fix rectx issues with templates, licenses, and more.\n\n")
	PrintSubcommands(configSubcommands, configSubcommandDetails, configSubcommandArguments)
	fmt.Println("\n  [flags]")
	configCmd.PrintDefaults()
}

func PrintSubcommands(subcommands, details []string, arguments []string) {
	fmt.Println("  [subcommands]")
	for i, command := range subcommands {
		fmt.Printf("  %s\n%sArguments: %s\n%s%s\n", command, megaSpace, arguments[i], megaSpace, details[i])
	}
}

func handleParseErrorAndHelpFlag(command *flag.FlagSet, err error, helpMenu func()) {
	if err != nil {
		name := command.Name()
		fmt.Println("An unexpected error occurred during flag parsing!")
		fmt.Printf("Please try \"rectx %s --help\" for more information on the %s command!\n", name, name)
	} else if helpFlag {
		helpMenu()
	}
}
