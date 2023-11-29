package main

import (
	"esefexbot/api"
	c "esefexbot/appcontext"
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
