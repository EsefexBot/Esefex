package dcgoutil

import (
	"log"

	"github.com/bwmarrin/discordgo"
)

// util functions for discordgo

func BotInVC(ds *discordgo.Session, serverID, channelID string) (bool, error) {
	vs, err := ds.State.VoiceState(serverID, ds.State.User.ID)
	if err != nil {
		log.Printf("Error getting bot voice state: %s\n", err)
		return false, err
	}
	if vs == nil {
		log.Println("Bot is not in a voice channel")
		return false, nil
	}

	// log.Printf("Bot is in channel %s\n", vs.ChannelID)
	// log.Printf("VoiceState: %+v\n", vs)
	// log.Printf("asked id: %s\n", channelID)

	return vs.ChannelID == channelID, nil
}

func GetBotVC(ds *discordgo.Session, serverID string) (*discordgo.VoiceState, error) {
	return ds.State.VoiceState(serverID, ds.State.User.ID)
}

// gets a list of users in a the channel the bot is in (excluding the bot)
func GetVCUsers(s *discordgo.Session, serverID, channelID string) ([]*discordgo.VoiceState, error) {
	// Get the Guild object
	guild, err := s.State.Guild(serverID)
	if err != nil {
		return nil, err
	}

	// Find the VoiceChannel object
	voiceChannel, err := s.State.Channel(channelID)
	if err != nil {
		return nil, err
	}

	// Get the list of VoiceStates for the VoiceChannel
	voiceStates := guild.VoiceStates
	channelUsers := []*discordgo.VoiceState{}

	for _, vs := range voiceStates {
		if vs.ChannelID == voiceChannel.ID {
			channelUsers = append(channelUsers, vs)
		}
	}

	return channelUsers, nil
}

func UserVC(s *discordgo.Session, serverID, userID string) (string, error) {
	vs, err := s.State.VoiceState(serverID, userID)
	if err != nil {
		return "", err
	}
	if vs == nil {
		return "", nil
	}

	return vs.ChannelID, nil
}
