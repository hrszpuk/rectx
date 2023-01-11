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
	utilities.Check(os.Mkdir(GetRectxPath()+"/templates", os.ModePerm))

	domain := "https://hrszpuk.github.io/rectx/templates/"
	for _, name := range TEMPLATE {
		go utilities.DownloadFile(
			domain+name+".rectx.template",
			GetRectxPath()+"/templates/"+name+".rectx.template",
		)
	}
}
