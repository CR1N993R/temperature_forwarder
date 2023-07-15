package client

import (
	"gopkg.in/yaml.v3"
	"os"
	"path/filepath"
)

type Clients struct {
	Clients []string `yaml:"clients"`
}

func readSpaceConfig() []byte {
	fBytes, err := os.ReadFile(filepath.Clean("clients.yml")) // ??? s.o.
	if err != nil {
		panic("can not read config file")
	}
	return fBytes
}

func LoadConfigFromFile() []string {
	content := readSpaceConfig()
	var clients Clients
	err := yaml.Unmarshal(content, &clients)
	if err != nil {
		panic("Error during space processing")
	}
	return clients.Clients
}
