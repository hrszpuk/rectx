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
	utilities.Check(err, true, "Attempt to read templates/ failed despite existence checks.")

	var templateList []string

	for _, entry := range templateDir {
		if !entry.IsDir() {
			templateList = append(templateList, entry.Name())
		}
	}

	return templateList
}

func ListTemplates() []string {
	config.ValidateConfig()

	dir, err := os.ReadDir(utilities.GetRectxPath() + "/templates")
	utilities.Check(err, true, "Attempt to read templates/ failed despite existence checks.")

	var list []string

	for _, entry := range dir {
		list = append(list, entry.Name())
	}

	return list
}

func Test(templateName string) {
	file, err := os.ReadFile(utilities.GetRectxPath() + "/templates/" + templateName)
	utilities.Check(err, true, fmt.Sprintf("Attempt to read template \"%s\" failed.", templateName))

	tempDir := utilities.GetRectxPath() + "/.temp"
	err = os.Mkdir(tempDir, os.ModeDir)
	if err != nil {
		tempDir = ".temp"
		err = os.Mkdir(tempDir, os.ModeDir)
		utilities.Check(err, true, "Attempt to generate a temporary testing environment failed!")
	}

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
			statement.Generate(tempDir)
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

	err = os.RemoveAll(tempDir)
	message := "Attempt to remove testing folder failed."
	if tempDir == ".temp" {
		message += ".. You may have to remove this file manually. (sudo rm ./.temp -rf)"
	}
	utilities.Check(err, false, message)
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
		} else if !dir.IsDir() && func() bool {
			filename := strings.ToLower(name)
			if filename != "commands" {
				return false
			}
			fileExtension := strings.Split(filename, ".")
			if len(fileExtension) < 2 {
				return true
			}
			extension := fileExtension[1]
			for _, allowedExtensions := range []string{"txt", "rectx"} {
				if extension == allowedExtensions {
					return true
				}
			}
			return false
		}() {
			fileBytes, err := os.ReadFile(templateName + "/" + sourceDir + name)
			utilities.Check(err, true, fmt.Sprintf("Attempt to read \"%s\" for command statements failed.", name))

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
			utilities.Check(err, true, fmt.Sprintf("Attempt to read \"%s\" for content block failed.", name))

			fileContent = strings.Replace(fileContent, "@FILE_CONTENT_PLACEHOLDER@", string(fileBytes), 1)
			templateContents += fmt.Sprintf("FILE %s %s %s\n", name, sourceDir, fileContent)
		}
		return nil
	})
	utilities.Check(err, true, "Path walk failed.")
	templateContents += templateCommands
	file, err := os.Create(templateName + ".rectx.template")
	utilities.Check(err, true, "Attempt to create .rectx.template file failed for an unknown reason!")
	_, err = file.WriteString(templateContents)
	utilities.Check(err, true, "Attempt to write template contents to .rectx.template file failed.")
	utilities.Check(file.Close(), false, "Could not close generated template file!")
	fmt.Printf("Snapshot complete... Generated \"%s\"!\n", templateName+".rectx.template")
}
