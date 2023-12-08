package util

import (
	"fmt"
	"github.com/bwmarrin/discordgo"
	"io"
	"net/http"
	"os"
	"regexp"
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

func DownloadSound(url string, path string) error {
	out, err := os.Create(path)
	if err != nil {
		return err
	}
	defer out.Close()

	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 || resp.Header.Get("Content-Type") != "audio/mpeg" {
		return fmt.Errorf("status code error: %v", resp.StatusCode)
	}

	_, err = io.Copy(out, resp.Body)
	if err != nil {
		return err
	}

	return nil
}
