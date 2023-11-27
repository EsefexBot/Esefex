package util

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"regexp"

	"github.com/bwmarrin/discordgo"
)

func PathExists(path string) bool {
	_, err := os.Stat(path)
	return !os.IsNotExist(err)
}

func ExtractIconUrl(icon *discordgo.ApplicationCommandInteractionDataOption) string {
	r, err := regexp.Compile(`<:.+:\d+>`)
	if err != nil {
		panic(err)
	}

	m := r.FindString(fmt.Sprint(icon.Value))

	rn, err := regexp.Compile(`\d+`)
	if err != nil {
		panic(err)
	}

	id := rn.FindString(m)

	return fmt.Sprintf("https://cdn.discordapp.com/emojis/%v.webp", id)
}

func GetSoundURL(guildID, name string) string {
	return fmt.Sprintf("https://cdn.discordapp.com/attachments/%v/%v.mp3", guildID, name)
}

func DownloadSound(url string) ([]byte, error) {
	req, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
	}

	if req.Header.Get("Content-Type") != "audio/mpeg" {
		return nil, fmt.Errorf("invalid content type: %v", req.Header.Get("Content-Type"))
	}

	len := req.ContentLength
	buf := make([]byte, len)
	req.Body.Read(buf)

	return buf, nil
}
