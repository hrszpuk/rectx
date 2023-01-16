package config

import (
	"bytes"
	"github.com/BurntSushi/toml"
	"os"
	"rectx/utilities"
)

func LoadConfig(path string) *Config {
	var config Config
	_, err := toml.DecodeFile(path, &config)
	utilities.Check(err)
	return &config
}

func DumpConfig(path string, config *Config) {
	f, err := os.Open(path)
	utilities.Check(err)
	defer f.Close()

	buffer := new(bytes.Buffer)
	err = toml.NewEncoder(buffer).Encode(config)
	utilities.Check(err)

	f.Write(buffer.Bytes())
}

// GenerateNewConfigDirectory generates a new config file, templates folder, license folder, etc
func GenerateNewConfigDirectory() {
	utilities.Check(os.Mkdir(utilities.GetRectxPath(), os.ModePerm))
	GenerateDefaultConfigFile()
	GenerateLicenses()
	GenerateTemplates()
}

// ValidateConfig Check if config exists and if not generate it
func ValidateConfig() {
	home := utilities.GetRectxPath()

	/// Validation
	// Here we're just checking if the correct files/folders exist for ~/.rectx

	if _, err := os.Stat(home); os.IsNotExist(err) {
		// if ~/.rectx doesn't exist then we need to regenerate *everything*
		GenerateNewConfigDirectory()
		return
	}

	if _, err := os.Stat(home + "/config.toml"); os.IsNotExist(err) {
		// Download default config file from source and put it in config.toml
		GenerateDefaultConfigFile()
	}

	if _, err := os.Stat(home + "/templates"); os.IsNotExist(err) {
		// if ~/.rectx/templates generation is handled by the templates module
		GenerateTemplates()
	}

	if _, err := os.Stat(home + "/licenses"); os.IsNotExist(err) {
		// if ~/.rectx/templates generation is handled by the templates module
		GenerateLicenses()
	}
}

func GenerateDefaultConfigFile() {
	home := utilities.GetRectxPath()
	utilities.DownloadFile("https://hrszpuk.github.io/rectx/defaultConfig.toml", home+"/config.toml")
}
