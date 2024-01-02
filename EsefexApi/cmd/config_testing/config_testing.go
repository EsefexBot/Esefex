package main

import (
	"esefexapi/config"
	"log"

	"github.com/pelletier/go-toml"
)

func main() {
	cfg, err := config.LoadConfig("config.toml")
	if err != nil {
		log.Fatal(err)
	}

	log.Printf("Config: %+v", cfg)

	cfg = &config.Config{
		HttpApi: config.HttpApi{
			Port:           8080,
			Domain:         "http://localhost",
			CustomProtocol: "esefex",
		},
		FileSoundDB: config.FileSoundDatabase{
			Location: "/tmp/esefexapi",
		},
		Bot: config.Bot{},
	}

	utoml, err := toml.Marshal(cfg)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(string(utoml))
}
