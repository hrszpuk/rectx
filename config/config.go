package config

import (
	"bytes"
	"os"
	"os/user"
	"rectx/utilities"

	"github.com/BurntSushi/toml"
)

func NewConfig() *Config {
	var conf Config
	User, err := user.Current()
	utilities.Check(err)

	conf.User.Author = User.Username
	conf.User.Email = ""
	conf.Compiler.Preference = "rgoc"

	conf.Template.DownloadSource = utilities.GetRectxDownloadSource() + "/templates"
	conf.Template.Location = utilities.GetRectxPath() + "/templates"
	conf.Template.Default = "default.rectx.template"
	conf.Template.Standards = []string{"default", "short", "short_with_build"}

	conf.Licenses.DownloadSource = utilities.GetRectxDownloadSource() + "/licenses"
	conf.Licenses.Location = utilities.GetRectxPath() + "/licenses"

	return &conf
}

func (self *Config) Load(path string) *Config {
	_, err := toml.DecodeFile(path, &self)
	utilities.Check(err)
	return self
}

func (self *Config) Dump(path string) *Config {
	var f *os.File
	if _, err := os.Stat(path); os.IsExist(err) {
		f, err = os.OpenFile(path, os.O_WRONLY, os.ModeType)
	} else {
		f, err = os.Create(path)
		utilities.Check(err)
	}

	defer f.Close()

	buffer := new(bytes.Buffer)
	err := toml.NewEncoder(buffer).Encode(self)
	utilities.Check(err)

	f.Write(buffer.Bytes())

	return self
}

// GenerateNewConfigDirectory generates a new config file, templates folder, license folder, etc
func GenerateNewConfigDirectory() {
	utilities.Check(os.Mkdir(utilities.GetRectxPath(), os.ModePerm))
	NewConfig().Dump(utilities.GetRectxPath() + "/config.toml")
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
		// Generate default config file and put it in config.toml
		NewConfig().Dump(utilities.GetRectxPath() + "/config.toml")
	}

	if _, err := os.Stat(home + "/templates"); os.IsNotExist(err) {
		// if ~/.rectx/templates generation is handled by the templates module
		GenerateTemplates()
	} else {
		ValidateTemplates()
	}

	if _, err := os.Stat(home + "/licenses"); os.IsNotExist(err) {
		// if ~/.rectx/templates generation is handled by the templates module
		GenerateLicenses()
	} else {
		ValidateLicenses()
	}
}
