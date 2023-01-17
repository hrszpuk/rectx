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

func (folder *FolderStatement) Generate() {
	utilities.Check(os.Mkdir(folder.SourceDir+folder.Name, 0750))
}
