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
			Content: fmt.Sprintf("The web interface can be found at %s/static/simpleui", c.domain),
		},
	}, nil
}
