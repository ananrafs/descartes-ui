package config

type (
	Config struct {
		Http     HttpConfig     `yaml:"http"`
		Frontend FrontendConfig `yaml:"frontend"`
	}

	HttpConfig struct {
		Port string `yaml:"port"`
	}

	FrontendConfig struct {
		Dist string `yaml:"dist"`
	}
)
