package util

import (
	"fmt"
	"os"
	"regexp"
	"strings"

	"github.com/bwmarrin/discordgo"
	"github.com/samber/lo"
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

func GetEmojiURL(emoji string) string {
	runes := []rune(emoji)

	name := strings.Join(lo.Map(runes, func(r rune, index int) string { return fmt.Sprintf("%x", r) }), "-")

	url := fmt.Sprintf("https://raw.githubusercontent.com/twitter/twemoji/master/assets/svg/%s.svg", name)

	return url
}
