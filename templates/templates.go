package templates

import (
	"fmt"
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

func Test(templateName string) {
	file, err := os.ReadFile(utilities.GetRectxPath() + "/templates/" + templateName)
	utilities.Check(err)

	err = os.Mkdir(".temp", os.ModeDir)
	utilities.Check(err)

	BadStatementCounter := 0
	FileStatementCounter := 0
	FolderStatementCounter := 0
	CommandStatementCounter := 0

	parser := NewTemplateParser(string(file))
	stmts := parser.Parse()
	for _, statement := range stmts {
		if statement.GetType() == "ERROR" {
			BadStatementCounter++
			fmt.Printf("ERROR: %s\n", statement.GetName())
		} else {
			fmt.Printf("Generating \"%s\"...", statement.GetType())
			statement.Generate(utilities.GetRectxPath() + "/.temp/")
			fmt.Print("DONE! ")
			if statement.GetType() == "FILE" {
				FileStatementCounter++
				fmt.Printf("Created \"%s\"!\n", statement.GetName())
			} else if statement.GetType() == "FOLDER" {
				FolderStatementCounter++
				fmt.Printf("Created \"%s\"!\n", statement.GetName())
			} else {
				CommandStatementCounter++
				fmt.Printf("Executed \"%s\"!\n", statement.GetName())
			}
		}
	}

	err = os.RemoveAll(".temp")
	utilities.Check(err)
}

func Snapshot(path string) {
	// TODO
}
