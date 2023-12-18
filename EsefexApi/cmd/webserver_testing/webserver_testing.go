package main

import (
	"esefexapi/api"
	c "esefexapi/appcontext"
	"log"
)

func main() {
	log.SetFlags(log.LstdFlags | log.Lshortfile)

	println("Starting webserver...")

	context := c.Context{
		CustomProtocol: "esefex",
		ApiPort:        "8080",
	}

	api.Run(&context)
}
