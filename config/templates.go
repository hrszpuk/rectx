package config

import (
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

	domain := "https://hrszpuk.github.io/rectx/templates/"
	for _, name := range TEMPLATE {
		utilities.DownloadFile(
			domain+name+".rectx.template",
			utilities.GetRectxPath()+"/templates/"+name+".rectx.template",
		)
	}
	// TODO validate /templates has template files in it
}
