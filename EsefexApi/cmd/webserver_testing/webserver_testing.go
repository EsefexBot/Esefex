package main

import (
	"esefexapi/api"
	c "esefexapi/ctx"
	"log"
)

func main() {
	log.SetFlags(log.LstdFlags | log.Lshortfile)

	println("Starting webserver...")

	context := c.Ctx{
		CustomProtocol: "esefex",
		ApiPort:        "8080",
	}

	api.Run(&context)
}
