package commands

import (
	"fmt"

	"github.com/bwmarrin/discordgo"
)

var (
	SessionCommand = &discordgo.ApplicationCommand{
		Name:        "session",
		Description: "Get join link for session",
	}
)

func Session(s *discordgo.Session, i *discordgo.InteractionCreate) {
	protocol := "esefex"
	route := "joinsession"

	url := fmt.Sprintf("%s://%s/%s", protocol, route, i.GuildID)

	s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
		Type: discordgo.InteractionResponseChannelMessageWithSource,
		Data: &discordgo.InteractionResponseData{
			Content: url,
		},
	})
}
