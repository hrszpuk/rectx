package config

import (
	"os"
	"rectx/utilities"
)

type Config struct {
	// ReCT stuff
	rgocLocation       string
	rctcLocation       string
	compilerPreference string
	packagesLocation   string

	// Template stuff
	defaultTemplate        string
	templateLocation       string
	standardTemplates      []string
	templateDownloadSource string

	// User stuff
	author string
	email  string

	// License stuff
	defaultLicense        string
	licenseLocation       string
	licenseDownloadSource string

	// Config stuff
	configDownloadSource string
}

// GenerateNewConfigDirectory generates a new config file, templates folder, license folder, etc
func GenerateNewConfigDirectory() {
	utilities.Check(os.Mkdir(GetUserHome()+"/.rectx", os.ModePerm))
}

// ValidateConfigFile Check if config exists and if not generate it
func ValidateConfigFile() {
	home := GetUserHome()

	/// Validation
	// Check ~/.rectx exists
	if _, err := os.Stat(home + "/.rectx"); os.IsNotExist(err) {
		utilities.Check(os.Mkdir(home+"/.rectx", os.ModePerm))
	}

	// Check ~/.rectx/config.toml exists
	if _, err := os.Stat(home + "/.rectx/config.toml"); os.IsNotExist(err) {
		// Download default config file from source and put it in config.toml
		GenerateDefaultConfigFile()
	}
}

func GenerateDefaultConfigFile() {
	home := GetUserHome()
	utilities.DownloadFile("https://hrszpuk.github.io/rectx/defaultConfig.toml", home+"/.rectx/config.toml")
}

func GetUserHome() string {
	home, err := os.UserHomeDir()
	utilities.Check(err)
	return home
}
