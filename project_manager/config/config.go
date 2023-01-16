package config

import (
	"bytes"
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
	utilities.Check(err)
}

func (config *ProjectConfig) Dump(path string) {
	f, err := os.Open(path)
	utilities.Check(err)
	defer f.Close()

	buffer := new(bytes.Buffer)
	err = toml.NewEncoder(buffer).Encode(config)
	utilities.Check(err)

	f.Write(buffer.Bytes())
}
