package config

import (
	"io"
	"os"

	"github.com/pelletier/go-toml"
	"github.com/pkg/errors"
)

var instance *Config

type Config struct {
	HttpApi     HttpApi           `toml:"http_api"`
	FileSoundDB FileSoundDatabase `toml:"file_sound_database"`
	FileUserDB  FileUserDatabase  `toml:"file_user_database"`
	Bot         Bot               `toml:"bot"`
}

type HttpApi struct {
	Port           int    `toml:"port"`
	CustomProtocol string `toml:"custom_protocol"`
}

type FileSoundDatabase struct {
	Location string `toml:"location"`
}

type FileUserDatabase struct {
	Location string `toml:"location"`
}

type Bot struct {
	UseTimeouts bool    `toml:"use_timeouts"`
	Timeout     float32 `toml:"timeout"`
}

func LoadConfig(path string) (*Config, error) {
	// load config from file
	f, err := os.Open(path)
	if err != nil {
		return nil, errors.Wrap(err, "Error opening config file")
	}
	defer f.Close()

	configStr, err := io.ReadAll(f)
	if err != nil {
		return nil, errors.Wrap(err, "Error reading config file")
	}

	// log.Println("Loaded config from file")
	// log.Println(string(configStr))

	var config Config
	err = toml.Unmarshal(configStr, &config)
	if err != nil {
		return nil, errors.Wrap(err, "Error unmarshalling config")
	}

	return &config, nil
}
