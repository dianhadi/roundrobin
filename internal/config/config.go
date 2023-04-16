package config

import (
	"io/ioutil"

	"gopkg.in/yaml.v2"
)

type RoundRobin struct {
	Instances []string `yaml:"instances"`
}

type Handler struct {
	MaxRetry   int `yaml:"max-retry"`
	MaxTimeout int `yaml:"max-timeout"`
}

type Config struct {
	RoundRobin RoundRobin `yaml:"round-robin"`
	Handler    Handler    `yaml:"handler"`
}

func Init(filename string) (*Config, error) {
	filename = "files/config/" + filename
	data, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil, err
	}

	var cfg Config
	err = yaml.Unmarshal(data, &cfg)
	if err != nil {
		return nil, err
	}

	return &cfg, nil
}
