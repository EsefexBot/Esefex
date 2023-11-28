package main

import (
	"esefexbot/api"
	c "esefexbot/appcontext"
)

func main() {
	println("Starting webserver...")

	context := c.Context{
		CustomProtocol: "esefex",
		ApiPort:        "8080",
	}

	api.Run(&context)
}
