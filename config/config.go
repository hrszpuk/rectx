package config

import (
	"bytes"
	"os"
	"os/user"
	"rectx/utilities"

	"github.com/BurntSushi/toml"
)

// Generates a new ProjectConfig file with the default field values.
// When loading a ProjectConfig, you will create a new default ProjectConfig then overwrite it using the Load(path) method.
func NewConfig() *Config {

	var conf Config
	User, err := user.Current()
	if err != nil {
		utilities.Check(err, false, "Unable to fetch current system user... Proceeding with default. (non-fatal)")
		conf.User.Author = ""
	} else {
		conf.User.Author = User.Username
	}

	conf.Compiler.Preference = "rgoc"

	conf.Template.DownloadSource = utilities.GetRectxDownloadSource() + "/templates"
	conf.Template.Location = utilities.GetRectxPath() + "/templates"
	conf.Template.Default = "default.rectx.template"
	conf.Template.Standards = []string{"default", "short", "short_with_build"}

	conf.Licenses.DownloadSource = utilities.GetRectxDownloadSource() + "/licenses"
	conf.Licenses.Location = utilities.GetRectxPath() + "/licenses"

	return &conf
}

// Loads the values from the ProjectConfig file path provided and overwrites the current field values.
func (conf *Config) Load(path string) *Config {

	_, err := toml.DecodeFile(path, &conf)
	utilities.Check(err, true, "Attempt to decode ProjectConfig failed during a load!")
	return conf
}

// Dumps the values of the ProjectConfig struct into a ProjectConfig file.
func (conf *Config) Dump(path string) *Config {

	var f *os.File
	if _, err := os.Stat(path); os.IsExist(err) {
		f, err = os.OpenFile(path, os.O_WRONLY, os.ModeType)
		utilities.Check(err, true, "Attempt to open ProjectConfig for dump failed!")
	} else {
		f, err = os.Create(path)
		utilities.Check(err, true, "Attempt to recover broken ProjectConfig during dump failed!")
	}

	defer f.Close()

	buffer := new(bytes.Buffer)
	err := toml.NewEncoder(buffer).Encode(conf)
	utilities.Check(err, true, "Attempt to encode ProjectConfig into writeable bytes failed!")

	f.Write(buffer.Bytes())

	return conf
}

// Generates an entirely new ProjectConfig directory with all the bells and whistles (licenses, templates, etc)
func GenerateNewConfigDirectory() {
	if err := os.Mkdir(utilities.GetRectxPath(), os.ModePerm); os.IsPermission(err) {
		utilities.Check(err, true, "Attempt to create ProjectConfig directory failed due to permissions!")
	} else {
		utilities.Check(err, true, "Attempt to create ProjectConfig directory failed for unknown an unknown reason!")
	}
	NewConfig().Dump(utilities.GetRectxPath() + "/ProjectConfig.toml")
	GenerateLicenses()
	GenerateTemplates()
}

// Ensures all parts of the ProjectConfig exist and regenerates them if they do not.
func ValidateConfig() {
	home := utilities.GetRectxPath()

	/// Validation
	// Here we're just checking if the correct files/folders exist for ~/.rectx

	if _, err := os.Stat(home); os.IsNotExist(err) {
		// if ~/.rectx doesn't exist then we need to regenerate *everything*
		GenerateNewConfigDirectory()
		return
	}

	if _, err := os.Stat(home + "/ProjectConfig.toml"); os.IsNotExist(err) {
		NewConfig().Dump(utilities.GetRectxPath() + "/ProjectConfig.toml")
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
