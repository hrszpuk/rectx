package licenses

import (
	"os"
	"rectx/metavariables"
	"rectx/utilities"
)

func GenerateLicense(license string, variables map[string]string) {
	bytes, err := os.ReadFile(utilities.GetRectxPath() + "/licenses/" + license)
	utilities.Check(err, true, "Failed to fetch license file specified!")
	content := string(bytes)

	content = metavariables.NewParser(content, variables).Parse()

	file, err := os.Create(variables["%PROJECT_NAME%"] + "/" + "LICENSE")
	utilities.Check(err, false, "Failed to create LICENSE file for project!")
	_, err = file.WriteString(content)
	utilities.Check(err, false, "Attempted to write content to LICENSE file but failed.")
	err = file.Close()
	utilities.Check(err, false, "Failed to close file!")
}
