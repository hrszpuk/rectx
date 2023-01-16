package config

type (
	Config struct {
		Project      ProjectConfig
		BuildProfile BuildConfig
		RunProfile   RunConfig
	}

	ProjectConfig struct {
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

func CreateDefaultConfig() *Config {
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

	projectConfig := ProjectConfig{
		Name:    "",
		Authors: nil,
		Version: "",
	}

	config := Config{
		Project:      projectConfig,
		BuildProfile: buildConfig,
		RunProfile:   runConfig,
	}

	return &config
}
