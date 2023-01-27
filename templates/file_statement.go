package templates

import (
	"fmt"
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

func (file *FileStatement) Generate(projectName string) {
	f, err := os.Create(projectName + "/" + file.SourceDir + "/" + file.Name)
	message := fmt.Sprintf("Attempt to generate file \"%s\" during template-based project generation failed.", file.Name)
	utilities.Check(err, true, message)

	_, err = f.WriteString(file.Content)
	message = fmt.Sprintf("Attempt to write template-specified content to \"%s\" (file) failed.", file.Name)
	utilities.Check(err, true, message)
}

func (file *FileStatement) GetName() string {
	return file.Name
}

func (file *FileStatement) GetType() string {
	return "FILE"
}
