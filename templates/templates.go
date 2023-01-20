package templates

import (
	"os"
	"rectx/config"
	"rectx/utilities"
)

// Check for template files in ~/.rectx/templates/
//		- ~/.rectx/templates/default
//		- ~/.rectx/templates/short
//		- ~/.rectx/templates/short_alt

func FetchTemplates() []string {
	config.ValidateConfig()
	templateDir, err := os.ReadDir(utilities.GetRectxPath() + "/templates")
	utilities.Check(err)

	var templateList []string

	for _, entry := range templateDir {
		if !entry.IsDir() {
			templateList = append(templateList, entry.Name())
		}
	}

	return templateList
}

func ListTemplates() []string {
	dir, err := os.ReadDir(utilities.GetRectxPath() + "/templates")
	utilities.Check(err)

	var list []string

	for _, entry := range dir {
		list = append(list, entry.Name())
	}

	return list
}
