package config

import (
	"bytes"
	"github.com/BurntSushi/toml"
	"os"
	"rectx/utilities"
)

type Config struct {

	// User stuff
	author string `toml:"user.author"`
	email  string `toml:"user.email"`

	// License stuff
	defaultLicense        string `toml:"license.default"`
	licenseLocation       string `toml:"license.location"`
	licenseDownloadSource string `toml:"license.download.source"`

	// ReCT stuff
	rgocLocation       string `toml:"compiler.rgoc.location"`
	rctcLocation       string `toml:"compiler.rctc.location"`
	compilerPreference string `toml:"compiler.preference"`
	packagesLocation   string `toml:"compiler.packages.location"`

	// Template stuff
	defaultTemplate        string   `toml:"template.default"`
	templateLocation       string   `toml:"template.location"`
	standardTemplates      []string `toml:"template.standards"`
	templateDownloadSource string   `toml:"template.download.source"`

	// Config stuff
	configDownloadSource string `toml:"config.download.source"`
}

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
	utilities.Check(os.Mkdir(GetUserHome()+"/.rectx", os.ModePerm))
	GenerateDefaultConfigFile()
}

// ValidateConfigFile Check if config exists and if not generate it
func ValidateConfigFile() {
	home := GetUserHome()

	/// Validation
	// Here we're just checking if the correct files/folders exist for ~/.rectx

	if _, err := os.Stat(home + "/.rectx"); os.IsNotExist(err) {
		// if ~/.rectx doesn't exist then we need to regenerate *everything*
		GenerateNewConfigDirectory()
		return
	}

	if _, err := os.Stat(home + "/.rectx/config.toml"); os.IsNotExist(err) {
		// Download default config file from source and put it in config.toml
		GenerateDefaultConfigFile()
	}

	if _, err := os.Stat(home + "/.rectx/templates"); os.IsNotExist(err) {
		// if ~/.rectx/templates generation is handled by the templates module
	}

	if _, err := os.Stat(home + "/.rectx/licenses"); os.IsNotExist(err) {
		// if ~/.rectx/templates generation is handled by the templates module
		GenerateLicenses()
	}
}

func GenerateDefaultConfigFile() {
	home := GetUserHome()
	utilities.DownloadFile("https://hrszpuk.github.io/rectx/defaultConfig.toml", home+"/.rectx/config.toml")
}

func GenerateLicenses() {
	utilities.Check(os.Mkdir(GetUserHome()+"/.rectx/licenses", os.ModePerm))

	licenses := []string{
		"Apache_License_2.0",
		"Boost_Software_License",
		"GNU_AGPLv3",
		"GNU_GPL3",
		"GNU_LGPLv3",
		"MIT_License",
		"Mozilla_Public_License_2.0",
	}
	for _, license := range licenses {
		utilities.DownloadFile(
			"https://hrszpuk.github.io/rectx/licenses/"+license,
			GetUserHome()+"/.rectx/licenses/"+license,
		)
	}
}

func GetUserHome() string {
	home, err := os.UserHomeDir()
	utilities.Check(err)
	return home
}
