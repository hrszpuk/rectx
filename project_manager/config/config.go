package config

import (
	"bytes"
	"fmt"
	"github.com/BurntSushi/toml"
	"os"
	"rectx/utilities"
)

type (
	ProjectConfig struct {
		Project      ProjectDetailsConfig
		BuildProfile BuildConfig
		RunProfile   RunConfig
	}

	ProjectDetailsConfig struct {
		Name    string
		Authors []string
		Version string
	}

	BuildConfig struct {
		Compiler        string
		Optimisation    int
		SourceDirectory string
		BuildDirectory  string
		Commands        []string
		ExecutableName  string
	}

	RunConfig struct {
		ChangesCauseRebuild bool
		AlwaysRebuild       bool
		RebuildOptimisation int
		Commands            []string
	}
)

func CreateDefaultConfig() *ProjectConfig {
	runConfig := RunConfig{
		ChangesCauseRebuild: true,
		AlwaysRebuild:       false,
		RebuildOptimisation: 3,
		Commands:            []string{},
	}

	buildConfig := BuildConfig{
		Compiler:        "rgoc",
		Optimisation:    3,
		SourceDirectory: "",
		BuildDirectory:  "",
		Commands:        []string{},
		ExecutableName:  "",
	}

	projectConfig := ProjectDetailsConfig{
		Name:    "",
		Authors: nil,
		Version: "",
	}

	config := ProjectConfig{
		Project:      projectConfig,
		BuildProfile: buildConfig,
		RunProfile:   runConfig,
	}

	return &config
}

func (config *ProjectConfig) Load(path string) {
	_, err := toml.DecodeFile(path, config)
	utilities.Check(err, true, "Attempted to load project config but failed.")
}

func (config *ProjectConfig) Dump(path string) {
	if _, err := os.Stat(path); os.IsNotExist(err) {
		_, err = os.Create(path)
		utilities.Check(err, true, "Attempted to recover non-existent project config but failed to recover.")
	}

	f, err := os.OpenFile(path, os.O_WRONLY, os.ModeType)
	utilities.Check(err, true, "Attempted to open file in write only mode but failed for an unknown reason.")

	defer utilities.Check(f.Close(), true, "Attempt to close project config file and failed... Somehow?")

	buffer := new(bytes.Buffer)
	err = toml.NewEncoder(buffer).Encode(config)
	utilities.Check(err, true, "Attempted to deserialize config data but failed.")

	_, err = f.WriteString(buffer.String())
	utilities.Check(err, true, fmt.Sprintf("Attempted to write data to \"%s\" but failed for an unknown reason!", path))
}
