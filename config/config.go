package config

import (
	"github.com/jinzhu/configor"
	"log"
)

var (
	config *Config
)

func Init(path string) *Config {
	if config != nil {
		return config
	}

	cfg := &Config{}
	if err := configor.Load(cfg, path); err != nil {
		log.Fatalf("failed to load config file — %v\n", err)
	}
	if err := config.Validate(); err != nil {
		log.Fatalf("config is not valid – %v", err)
	}

	config = cfg
	return config
}

func Get() *Config {
	return config
}
