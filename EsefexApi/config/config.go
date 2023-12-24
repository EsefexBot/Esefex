package config

import (
	"io"
	"log"
	"os"

	"github.com/pelletier/go-toml"
)

var instance *Config

type Config struct {
	Test         string       `toml:"test"`
	HttpApi      HttpApi      `toml:"http_api"`
	FileDatabase FileDatabase `toml:"file_database"`
	Bot          Bot          `toml:"bot"`
}

type HttpApi struct {
	Port           int    `toml:"port"`
	Domain         string `toml:"domain"`
	CustomProtocol string `toml:"custom_protocol"`
}

type FileDatabase struct {
	Location string `toml:"location"`
}

type Bot struct {
}

func LoadConfig(path string) (*Config, error) {
	// load config from file
	f, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	configStr, err := io.ReadAll(f)
	if err != nil {
		return nil, err
	}

	log.Println("Loaded config from file")
	log.Println(string(configStr))

	var config Config
	err = toml.Unmarshal(configStr, &config)
	if err != nil {
		return nil, err
	}

	return &config, nil
}
