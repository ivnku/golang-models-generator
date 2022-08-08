package initApp

import (
	"fmt"
	"gopkg.in/yaml.v3"
	"io/ioutil"
	"models-generator/config"
)

func InitConfig(path string) (*config.AppConfig, error) {
	if path == "" {
		path = "./config.yml"
	}

	yamlFile, err := ioutil.ReadFile(path)
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

	fmt.Printf("Result: %v\n", appConfig)
	return appConfig, nil
}
