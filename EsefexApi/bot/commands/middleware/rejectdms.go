package middleware

import (
	"esefexapi/bot/commands/cmdhandler"
	"esefexapi/util/dcgoutil"

	"github.com/bwmarrin/discordgo"
)

// RejectDMs rejects DMs
// This is important because we don't want to allow users to use commands in DMs
// because it will cause nil pointer errors
func (m *CommandMiddleware) RejectDMs(next cmdhandler.CommandHandlerWithErr, perms ...string) cmdhandler.CommandHandlerWithErr {
	return func(s *discordgo.Session, i *discordgo.InteractionCreate) (*discordgo.InteractionResponse, error) {
		if i.GuildID == "" {
			return &discordgo.InteractionResponse{
				Type: discordgo.InteractionResponseChannelMessageWithSource,
				Data: &discordgo.InteractionResponseData{
					Embeds: []*discordgo.MessageEmbed{
						{
							Title:       "This command does not work in DMs",
							Description: "Please use this command in a server that I am in or invite me to your server",
							Color:       0xFF0000,
						},
					},
					Components: []discordgo.MessageComponent{
						discordgo.ActionsRow{
							Components: []discordgo.MessageComponent{
								discordgo.Button{
									Label: "Invite me to your server",
									Style: discordgo.LinkButton,
									URL:   dcgoutil.GetInviteLink(s.State.User.ID),
									Emoji: discordgo.ComponentEmoji{Name: "🤖"},
								},
							},
						},
					},
				},
			}, nil
		}

		return next(s, i)
	}
}
