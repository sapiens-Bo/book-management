// Package config ...
package config

import (
	"fmt"
	"log"
	"os"

	"github.com/ilyakaznacheev/cleanenv"
	"gopkg.in/yaml.v3"
)

const envVar = "CONFIG_BMAPP_PATH"

// Config is struct for config file
type Config struct {
	Path string `yaml:"path" path-default:"/$HOME/Books/"`
	App  string `yaml:"app" app-default:"zathura"`
	Last string `yaml:"last"`
}

// MustLoad func loaded and return Config type
func MustLoad() *Config {
	configPath := os.Getenv(envVar)
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

// SetLast func ...
func (c *Config) SetLast(name string) {
	data, err := os.ReadFile(os.Getenv(envVar))
	if err != nil {
		fmt.Printf("set last error: %s", err)
		os.Exit(1)
	}

	err = yaml.Unmarshal(data, c)
	if err != nil {
		fmt.Printf("parse error: %s", err)
		os.Exit(1)
	}

	c.Last = name

	newData, err := yaml.Marshal(c)
	if err != nil {
		fmt.Printf("marshal error: %s", err)
		os.Exit(1)
	}

	err = os.WriteFile(os.Getenv(envVar), newData, 0644)
	if err != nil {
		fmt.Printf("write file error: %s", err)
		os.Exit(1)
	}
}
