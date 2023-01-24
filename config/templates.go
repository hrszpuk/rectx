package config

import (
	"fmt"
	"os"
	"rectx/utilities"
	"strings"
)

var TEMPLATE = [...]string{
	"default",
	"short",
	"short_with_build",
}

func GenerateTemplates() {
	utilities.Check(os.Mkdir(utilities.GetRectxPath()+"/templates", os.ModePerm))

	DownloadTemplates(utilities.GetRectxPath() + "/templates/")
	ValidateTemplates()
}

func DownloadTemplates(path string) {
	domain := "https://hrszpuk.github.io/rectx/templates/"
	for _, name := range TEMPLATE {
		utilities.DownloadFile(
			domain+name+".rectx.template",
			path+name+".rectx.template",
		)
	}
}

func ValidateTemplates() {
	dir, err := os.ReadDir(utilities.GetRectxPath() + "/templates")
	utilities.Check(err)

	if len(dir) < 1 {
		DownloadTemplates(utilities.GetRectxPath() + "/templates/")
		dir, err = os.ReadDir(utilities.GetRectxPath() + "/templates")
		utilities.Check(err)

		if len(dir) < 1 {
			fmt.Println("ERROR: Could not download templates for an unknown reason!")
		}
	}

	if len(dir) < 3 {
		fmt.Printf("ERROR: Expected at least %d templates but only found %d! You may want to regenerate the template files using \"rectx config regenerate --templates\"!\n", len(TEMPLATE), len(dir))
	}
}

func AddTemplate(path string) {
	if !strings.HasSuffix(path, ".rectx.template") {
		fmt.Printf("Unable to add template because \"%s\" is not a rectx template file!", path)
		os.Exit(1)
	}

	bytes, err := os.ReadFile(path)
	utilities.Check(err)

	ValidateTemplates()
	pathSplit := strings.Split(path, "/")

	file, err := os.Create(utilities.GetRectxPath() + "/templates/" + pathSplit[len(pathSplit)-1])
	_, err = file.WriteString(string(bytes))
	utilities.Check(err)

	utilities.Check(file.Close())
	fmt.Printf("Added new template called \"%s\"!", pathSplit)
	fmt.Println(
		"If you want to rename this template use: rectx template rename <name> <newName>",
		"For more information on templates please use rectx template --help",
	)
}
