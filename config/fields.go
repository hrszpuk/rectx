package config

type (
	Config struct {
		User                     UserConfig     `toml:"user"`
		Licenses                 LicenseConfig  `toml:"licences"`
		Compiler                 CompilerConfig `toml:"compiler"`
		Template                 TemplateConfig `toml:"templates"`
		DefaultConfigDownloadUrl string         `toml:"defaultConfigDownloadUrl"`
	}

	UserConfig struct {
		Author string
		Email  string
	}

	LicenseConfig struct {
		Default        string
		Location       string
		DownloadSource string
	}

	CompilerConfig struct {
		RgocLocation     string
		RctcLocation     string
		Preference       string
		PackagesLocation string
	}

	TemplateConfig struct {
		Default        string
		Location       string
		Standards      []string
		DownloadSource string
	}
)
