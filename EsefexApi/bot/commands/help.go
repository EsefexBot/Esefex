package commands

import (
	"esefexapi/bot/commands/helpmsg"

	"github.com/bwmarrin/discordgo"
	"github.com/pkg/errors"
)

var HelpCommand = &discordgo.ApplicationCommand{
	Name:        "help",
	Description: "Get help about the bot.",
	Options: []*discordgo.ApplicationCommandOption{
		{
			Name:        "help-category",
			Description: "The category to get help about. Leave empty to get general help",
			Type:        discordgo.ApplicationCommandOptionString,
			Required:    false,
			Choices: []*discordgo.ApplicationCommandOptionChoice{
				{
					Name:  "UI",
					Value: "UI",
				},
				{
					Name:  "Commands:Bot",
					Value: "Commands:Bot",
				},
				{
					Name:  "Commands:Sound",
					Value: "Commands:Sound",
				},
				{
					Name:  "Commands:User",
					Value: "Commands:User",
				},
				{
					Name:  "Commands:Permission",
					Value: "Commands:Permission",
				},
			},
		},
	},
}

func (c *CommandHandlers) Help(s *discordgo.Session, i *discordgo.InteractionCreate) (*discordgo.InteractionResponse, error) {
	category := ""
	if len(i.ApplicationCommandData().Options) > 0 {
		category = i.ApplicationCommandData().Options[0].StringValue()
	} else {
		category = ""
	}

	msg, err := helpmsg.GetHelpMessage(category)
	if err != nil {
		return nil, errors.Wrap(err, "failed to get help message")
	}

	return &discordgo.InteractionResponse{
		Type: discordgo.InteractionResponseChannelMessageWithSource,
		Data: &discordgo.InteractionResponseData{
			Content: msg,
		},
	}, nil
}
