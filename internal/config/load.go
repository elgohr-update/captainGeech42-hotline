package config

import (
	"io/ioutil"
	"os"

	yaml "gopkg.in/yaml.v2"
)

func LoadConfig(path string) *Config {
	// TODO make this try to load from an env var too
	config := &Config{}

	raw, err := ioutil.ReadFile(os.ExpandEnv(path))
	if err != nil {
		panic(err)
	}

	err = yaml.Unmarshal(raw, &config)
	if err != nil {
		panic(err)
	}

	return config
}
