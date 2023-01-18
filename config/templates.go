package config

import (
	"fmt"
	"os"
	"rectx/utilities"
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
