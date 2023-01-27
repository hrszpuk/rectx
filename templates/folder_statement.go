package templates

import (
	"fmt"
	"os"
	"rectx/utilities"
)

type FolderStatement struct {
	Name      string
	SourceDir string
}

func NewFolderStatement(name string, sourceDir string) *FolderStatement {
	return &FolderStatement{
		Name:      name,
		SourceDir: sourceDir,
	}
}

func (folder *FolderStatement) Generate(projectName string) {
	message := fmt.Sprintf("Attempt to generate folder \"%s\" during template-based project generation failed. This may affect the generation of other folders and files.", folder.Name)
	utilities.Check(os.Mkdir(projectName+"/"+folder.SourceDir+"/"+folder.Name, 0750), true, message)
}

func (folder *FolderStatement) GetName() string {
	return folder.Name
}

func (folder *FolderStatement) GetType() string {
	return "FOLDER"
}
