package util

import (
	"fmt"
	"os"

	"github.com/bwmarrin/discordgo"
)

func PathExists(path string) bool {
	_, err := os.Stat(path)
	return !os.IsNotExist(err)
}

func GetIconURL(icon *discordgo.ApplicationCommandInteractionDataOption) string {
	return fmt.Sprintf("https://cdn.discordapp.com/emojis/%v.webp", icon.Value)
}

func GetSoundURL(guildID, name string) string {
	return fmt.Sprintf("https://cdn.discordapp.com/attachments/%v/%v.mp3", guildID, name)
}
