// Package config ...
package config

import (
	"log"
	"os"

	"github.com/ilyakaznacheev/cleanenv"
)

// Config is struct for config file
type Config struct {
	Path string `yaml:"path" path-default:"~/Books"`
}

// MustLoad func loaded and return Config type
func MustLoad() *Config {
	configPath := os.Getenv("CONFIG_BMAPP_PATH")
	if configPath == "" {
		log.Fatal("CONFIG_PATH is not set")
	}

	//check if file exists
	if _, err := os.Stat(configPath); os.IsNotExist(err) {
		log.Fatalf("config file does not exists: %s", configPath)
	}

	var cfg Config
	if err := cleanenv.ReadConfig(configPath, &cfg); err != nil {
		log.Fatalf("cannot read config: %s", err)
	}

	return &cfg
}
