package config

import (
	"bytes"
	"os"
	"os/user"
	"rectx/utilities"

	"github.com/BurntSushi/toml"
)

// Generates a new config file with the default field values.
// When loading a config, you will create a new default config then overwrite it using the Load(path) method.
func NewConfig() *Config {

	var conf Config
	User, err := user.Current()
	utilities.Check(err)

	conf.User.Author = User.Username

	conf.Compiler.Preference = "rgoc"

	conf.Template.DownloadSource = utilities.GetRectxDownloadSource() + "/templates"
	conf.Template.Location = utilities.GetRectxPath() + "/templates"
	conf.Template.Default = "default.rectx.template"
	conf.Template.Standards = []string{"default", "short", "short_with_build"}

	conf.Licenses.DownloadSource = utilities.GetRectxDownloadSource() + "/licenses"
	conf.Licenses.Location = utilities.GetRectxPath() + "/licenses"

	return &conf
}

// Loads the values from the config file path provided and overwrites the current field values.
func (conf *Config) Load(path string) *Config {

	_, err := toml.DecodeFile(path, &conf)
	utilities.Check(err)
	return conf
}

// Dumps the values of the config struct into a config file.
func (conf *Config) Dump(path string) *Config {

	var f *os.File
	if _, err := os.Stat(path); os.IsExist(err) {
		f, err = os.OpenFile(path, os.O_WRONLY, os.ModeType)
		utilities.Check(err)
	} else {
		f, err = os.Create(path)
		utilities.Check(err)
	}

	defer f.Close()

	buffer := new(bytes.Buffer)
	err := toml.NewEncoder(buffer).Encode(conf)
	utilities.Check(err)

	f.Write(buffer.Bytes())

	return conf
}

//Generates an entirely new config directory with all the bells and whistles (licenses, templates, etc)
func GenerateNewConfigDirectory() {
	utilities.Check(os.Mkdir(utilities.GetRectxPath(), os.ModePerm))
	NewConfig().Dump(utilities.GetRectxPath() + "/config.toml")
	GenerateLicenses()
	GenerateTemplates()
}

// Ensures all parts of the config exist and regenerates them if they do not.
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
		NewConfig().Dump(utilities.GetRectxPath() + "/config.toml")
	}

	if _, err := os.Stat(home + "/templates"); os.IsNotExist(err) {
		GenerateTemplates()
	} else {
		ValidateTemplates()
	}

	if _, err := os.Stat(home + "/licenses"); os.IsNotExist(err) {
		GenerateLicenses()
	} else {
		ValidateLicenses()
	}
}
