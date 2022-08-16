package config

type AppConfig struct {
	Connection struct {
		Driver string `yaml:"driver"`
		Dsn    string `yaml:"dsn"`
	}
	TemplatePath        string `yaml:"templatePath"`
	GeneratedModelsPath string `yaml:"generatedModelsPath"`
}
