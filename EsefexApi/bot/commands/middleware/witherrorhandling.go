package middleware

import (
	"esefexapi/bot/commands/cmdhandler"
	"fmt"
	"log"

	"github.com/bwmarrin/discordgo"
	"github.com/pkg/errors"
)

// This is a special middleware that wraps the command handler
// It will catch any errors that occur while executing the command
// and respond to the interaction with the error
// This middleware should be the last middleware in the chain (i.e. outermost)
func (m *CommandMiddleware) WithErrorHandling(next cmdhandler.CommandHandlerWithErr) func(s *discordgo.Session, i *discordgo.InteractionCreate) {
	return func(s *discordgo.Session, i *discordgo.InteractionCreate) {
		r, err := next(s, i)
		if err != nil {
			log.Printf("Cannot execute command: %+v", err)

			err = s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
				Type: discordgo.InteractionResponseChannelMessageWithSource,
				Data: &discordgo.InteractionResponseData{
					Embeds: []*discordgo.MessageEmbed{
						{
							Title:       "An error has occured while executing the command",
							Color:       0xff0000,
							Description: fmt.Sprintf("```%+v```", errors.Cause(err)),
						},
					},
				},
			})
			if err != nil {
				log.Printf("Cannot respond to interaction: %+v", err)
			}
		}

		if r != nil {
			err = s.InteractionRespond(i.Interaction, r)
			if err != nil {
				log.Printf("Cannot respond to interaction: %+v", err)

				err = s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
					Type: discordgo.InteractionResponseChannelMessageWithSource,
					Data: &discordgo.InteractionResponseData{
						Embeds: []*discordgo.MessageEmbed{
							{
								Title:       "Could not send response to interaction",
								Color:       0xff0000,
								Description: fmt.Sprintf("```%+v```", errors.Cause(err)),
							},
						},
					},
				})
				if err != nil {
					log.Printf("Cannot respond to interaction: %+v", err)
				}
			}
		}
	}
}
