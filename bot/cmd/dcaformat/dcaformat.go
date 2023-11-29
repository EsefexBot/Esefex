package main

import (
	"esefexbot/util"
	"log"
)

func main() {
	s, err := util.LoadDcaSound("debug.dca")
	if err != nil {
		panic(err)
	}

	log.Printf("Sending %v bytes", len(s))
	log.Printf("Width: %v", len(s[0]))
}
