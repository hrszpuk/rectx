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
		Optimisation    string
		SourceDirectory string
		BuildDirectory  string
		Commands        []string
		ExecutableName  string
	}

	RunConfig struct {
		Compiler       string
		Optimisation   string
		ExecutablePath string
		Commands       []string
	}
)
