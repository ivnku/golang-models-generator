package initApp

import (
	"fmt"
	"gopkg.in/yaml.v3"
	"models-generator/config"
	"os"
)

func InitConfig(path string) (*config.AppConfig, error) {
	if path == "" {
		path = "./config.yml"
	}

	yamlFile, err := os.ReadFile(path)
	if err != nil {
		err = fmt.Errorf("error parsing YAML file: %s\n", err)
		return nil, err
	}

	var appConfig *config.AppConfig
	err = yaml.Unmarshal(yamlFile, &appConfig)
	if err != nil {
		err = fmt.Errorf("error parsing YAML file: %s\n", err)
		return nil, err
	}

	return appConfig, nil
}
