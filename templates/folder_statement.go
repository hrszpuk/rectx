package templates

import (
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
	utilities.Check(os.Mkdir(projectName+"/"+folder.SourceDir+"/"+folder.Name, 0750))
}

func (folder *FolderStatement) GetName() string {
	return folder.Name
}

func (folder *FolderStatement) GetType() string {
	return "FOLDER"
}
