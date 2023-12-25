package commands

import (
	"esefexapi/bot/actions"
	"fmt"

	"github.com/bwmarrin/discordgo"
)

var (
	SessionCommand = &discordgo.ApplicationCommand{
		Name:        "session",
		Description: "Get join link for session",
	}
)

func (c *CommandHandlers) Session(s *discordgo.Session, i *discordgo.InteractionCreate) (*discordgo.InteractionResponse, error) {
	g, err := s.State.Guild(i.GuildID)
	if err != nil {
		return nil, err
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
		return &discordgo.InteractionResponse{
			Type: discordgo.InteractionResponseChannelMessageWithSource,
			Data: &discordgo.InteractionResponseData{
				Content: "You must be connected to a voice channel to get the session link.",
			}}, nil
	}

	route := "joinsession"

	// https://esefex.com/joinsession/1234567890
	url := fmt.Sprintf("https://%s/%s/%s", c.domain, route, i.GuildID)

	actions.JoinChannelVoice(s, i.GuildID, userChannel)

	return &discordgo.InteractionResponse{
		Type: discordgo.InteractionResponseChannelMessageWithSource,
		Data: &discordgo.InteractionResponseData{
			Content: url,
		},
	}, nil
}
