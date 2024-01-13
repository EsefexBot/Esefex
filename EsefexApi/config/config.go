package config

import (
	"io"
	"os"

	"github.com/pelletier/go-toml"
	"github.com/pkg/errors"
)

type Config struct {
	VerificationExpiry float32  `toml:"verification_expiry"`
	HttpApi            HttpApi  `toml:"http_api"`
	Database           Database `toml:"database"`
	Bot                Bot      `toml:"bot"`
}

type HttpApi struct {
	Port           int    `toml:"port"`
	CustomProtocol string `toml:"custom_protocol"`
}

type Database struct {
	SounddbLocation      string `toml:"sounddb_location"`
	UserdbLocation       string `toml:"userdb_location"`
	Permissiondblocation string `toml:"permissiondb_location"`
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
