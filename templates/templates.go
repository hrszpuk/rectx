package templates

import (
	"fmt"
	"io/fs"
	"os"
	"path/filepath"
	"rectx/config"
	"rectx/utilities"
	"strings"
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
	templateName := strings.Split(path, "/")[0]
	templateContents := "# This template was generated using ReCTx Template Snapshot!\n"
	err := filepath.WalkDir(path, func(xpath string, dir fs.DirEntry, err error) error {
		if dir.IsDir() {
			pathContents := strings.Split(xpath, "/")
			name := pathContents[len(pathContents)-1]
			sourceDir := ""
			if len(pathContents) > 1 {
				for i, pathSegment := range pathContents {
					if i == len(pathContents)-1 {
						break
					}
					sourceDir += pathSegment + "/"
				}
			}
			templateContents += fmt.Sprintf("FOLDER %s %s\n", name, sourceDir)
		} 
		return nil
	})
	fmt.Println(templateContents)
	utilities.Check(err)
}
