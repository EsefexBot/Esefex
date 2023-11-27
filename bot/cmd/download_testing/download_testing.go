package main

import (
	"esefexbot/filedb"
	"log"
	"net/http"
)

func main() {
	url := "https://cdn.discordapp.com/attachments/777344211828604950/1178474445673869404/y2mate.com_-_heheheha_Bass_boosted.mp3"

	req, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
	}

	len := req.ContentLength
	log.Println(len)

	buf := make([]byte, len)
	req.Body.Read(buf)

	filedb.AddSound("testserver", "test", ":sus:", buf)
}
