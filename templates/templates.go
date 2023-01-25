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
	fmt.Printf("Project: %s\n", templateName)
	templateContents := "# This template was generated using ReCTx Template Snapshot!\n"
	templateCommands := ""

	err := filepath.WalkDir(path, func(xpath string, dir fs.DirEntry, err error) error {
		if xpath == templateName {
			return nil
		}

		pathContents := strings.Split(xpath, "/")
		pathContents = pathContents[1:]
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

		if dir.IsDir() {
			templateContents += fmt.Sprintf("FOLDER %s %s\n", name, sourceDir)
		} else if !dir.IsDir() && strings.Split(strings.ToLower(name), ".")[0] == "commands" {
			fileBytes, err := os.ReadFile(templateName + "/" + sourceDir + name)
			utilities.Check(err)

			var buffer []byte
			for _, char := range fileBytes {
				buffer = append(buffer, char)
				if char == '\n' {
					templateCommands += fmt.Sprintf("COMMAND %s", string(buffer))
					buffer = []byte{}
				}
			}


		} else {
			fileContent := "{%@FILE_CONTENT_PLACEHOLDER@%}"
			fileBytes, err := os.ReadFile(templateName + "/" + sourceDir + name)
			utilities.Check(err)

			fileContent = strings.Replace(fileContent, "@FILE_CONTENT_PLACEHOLDER@", string(fileBytes), 1)
			templateContents += fmt.Sprintf("FILE %s %s %s\n", name, sourceDir, fileContent)
		}
		return nil
	})
	templateContents += templateCommands
	file, err := os.Create(templateName + ".rectx.template")
	utilities.Check(err)
	_, err = file.WriteString(templateContents)
	defer file.Close()
	utilities.Check(err)
	fmt.Printf("Snapshot complete... Generated \"%s\"\n!", templateName + ".rectx.template")
}
