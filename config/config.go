package config

type AppConfig struct {
	Connection struct {
		Driver string `yaml:"driver"`
		Dsn    string `yaml:"dsn"`
	}
}
