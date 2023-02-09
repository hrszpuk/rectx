package ProjectConfig

import "os"

// Checks if the .rectx config directory exists
func ValidateProjectConfigDirectory() {
	if _, err := os.Stat(".rectx"); os.IsNotExist(err) {
		GenerateProjectConfigDirectory()
	}
}

// Generates a .rectx config
func GenerateProjectConfigDirectory() {
	os.Mkdir(".rectx", os.ModeDir)

}
