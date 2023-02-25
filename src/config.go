package main

import (
	"io/ioutil"
	"os"

	"gopkg.in/yaml.v2"
)

type Config struct {
	BaseUrl string `yaml:"baseUrl"`
}

func getConfig() Config {
	configPath := os.Getenv("REST_CLIENT_CONFIG")

	yamlFile, err := ioutil.ReadFile(configPath)
	if err != nil {
		panic(err)
	}

	// Unmarshal YAML data into a generic interface value
	var config Config
	err = yaml.Unmarshal(yamlFile, &config)
	if err != nil {
		panic(err)
	}

	return config
}
