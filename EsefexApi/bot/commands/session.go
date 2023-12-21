package commands

import (
	"esefexapi/bot/actions"
	"fmt"
	"log"

	"github.com/bwmarrin/discordgo"
)

var (
	SessionCommand = &discordgo.ApplicationCommand{
		Name:        "session",
		Description: "Get join link for session",
	}
)

func (c *CommandHandlers) Session(s *discordgo.Session, i *discordgo.InteractionCreate) {
	g, err := s.State.Guild(i.GuildID)
	if err != nil {
		log.Printf("Cannot get guild: %v", err)
	}

	var userChannel string
	userConnected := false
	for _, vs := range g.VoiceStates {
		if vs.UserID == i.Member.User.ID {
			userConnected = true
			userChannel = vs.ChannelID
		}
	}
	if !userConnected {
		s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
			Type: discordgo.InteractionResponseChannelMessageWithSource,
			Data: &discordgo.InteractionResponseData{
				Content: "You must be connected to a voice channel to get the session link.",
			},
		})
		return
	}

	protocol := "esefex"
	route := "joinsession"

	url := fmt.Sprintf("%s://%s/%s", protocol, route, i.GuildID)

	s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
		Type: discordgo.InteractionResponseChannelMessageWithSource,
		Data: &discordgo.InteractionResponseData{
			Content: url,
		},
	})

	actions.JoinChannelVoice(s, i.GuildID, userChannel)
}
