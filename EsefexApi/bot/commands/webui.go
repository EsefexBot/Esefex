package commands

import (
	"fmt"

	"github.com/bwmarrin/discordgo"
)

var WebUICommand = &discordgo.ApplicationCommand{
	Name:        "webui",
	Description: "The link to the web interface.",
}

func (c *CommandHandlers) WebUI(s *discordgo.Session, i *discordgo.InteractionCreate) (*discordgo.InteractionResponse, error) {
	return &discordgo.InteractionResponse{
		Type: discordgo.InteractionResponseChannelMessageWithSource,
		Data: &discordgo.InteractionResponseData{
			Embeds: []*discordgo.MessageEmbed{
				{
					Title: "Web Interface",
					URL:   fmt.Sprintf("%s/static/simpleui", c.domain),
				},
			},
			Components: []discordgo.MessageComponent{
				discordgo.ActionsRow{
					Components: []discordgo.MessageComponent{
						discordgo.Button{
							Style: discordgo.LinkButton,
							Label: "Open Web Interface",
							URL:   fmt.Sprintf("%s/static/simpleui", c.domain),
							// TODO: For some reason, you need to set the emoji to something, otherwise the request will fail
							// This is a bug in the discordgo library
							// I should probably make a PR to fix this
							Emoji: discordgo.ComponentEmoji{
								Name: "📝",
							},
						},
					},
				},
			},
		},
	}, nil
}
