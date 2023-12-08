package main

import (
	"gopkg.in/hraban/opus.v2"
)

func main() {
	const sampleRate = 48000
	const channels = 2

	enc, err := opus.NewEncoder(sampleRate, channels, opus.AppVoIP)
	if err != nil {
		panic(err)
	}

}
