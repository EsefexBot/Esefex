package main

import (
	// "esefexbot/filedb"
	"log"
	"net/http"
	"os"
)

func main() {
	url := "https://cdn.discordapp.com/attachments/777344211828604950/1178752698552700978/test.txt"

	req, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(req.Header)

	len := req.ContentLength
	log.Println(len)

	buf := make([]byte, len)
	req.Body.Read(buf)

	os.WriteFile("test.txt", buf, os.ModePerm)

	// filedb.AddSound("testserver", "test", ":sus:", buf)
}
