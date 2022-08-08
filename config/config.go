package config

type AppConfig struct {
	Connection struct {
		Dsn string `yaml:"dsn"`
	}
}
