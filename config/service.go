package config

import (
	"gopkg.in/yaml.v3"
	"os"
	"path/filepath"
)

type Config struct {
	Port          int            `yaml:"port"`
	LokiInstances []LokiInstance `yaml:"loki_instances"`
	Clients       []string       `yaml:"clients"`
}

type LokiInstance struct {
	URL   string `yaml:"url"`
	Token string `yaml:"token"`
}

func readConfigFile() []byte {
	fBytes, err := os.ReadFile(filepath.Clean("config.yml")) // ??? s.o.
	if err != nil {
		panic("Cannot read config file")
	}
	return fBytes
}

func LoadConfigFromFile() Config {
	content := readConfigFile()
	var clients Config
	err := yaml.Unmarshal(content, &clients)
	if err != nil {
		panic("Error during space processing")
	}
	return clients
}
