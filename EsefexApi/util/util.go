package util

import (
	"fmt"
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
