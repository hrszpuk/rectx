package templates

import (
	"os"
	"rectx/utilities"
)

type FileStatement struct {
	Name      string
	SourceDir string
	Content   string
}

func NewFileStatement(name string, sourceDir string, content string) *FileStatement {
	return &FileStatement{
		Name:      name,
		SourceDir: sourceDir,
		Content:   content,
	}
}

func (file *FileStatement) Generate() {
	f, err := os.Create(file.SourceDir + file.Name)
	utilities.Check(err)

	_, err = f.WriteString(file.Content)
	utilities.Check(err)
}
