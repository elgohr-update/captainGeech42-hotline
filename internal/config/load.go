package config

import (
	"io/ioutil"

	yaml "gopkg.in/yaml.v2"
)

func LoadConfig(path string) *Config {
	config := &Config{}

	raw, err := ioutil.ReadFile(path)
	if err != nil {
		panic(err)
	}

	err = yaml.Unmarshal(raw, &config)
	if err != nil {
		panic(err)
	}

	return config
}
