package dcgoutil

import (
	"github.com/bwmarrin/discordgo"
	"github.com/pkg/errors"
)

var BotNotInVC = errors.New("Bot is not in a voice channel")
var UserNotInVC = errors.New("User is not in a voice channel")

// util functions for discordgo

func BotInVC(ds *discordgo.Session, serverID, channelID string) (bool, error) {
	vs, err := ds.State.VoiceState(serverID, ds.State.User.ID)
	if err != nil {
		return false, errors.Wrap(err, "Error getting bot voice state")
	}
	if vs == nil {
		return false, nil
	}

	return vs.ChannelID == channelID, nil
}

func GetBotVC(ds *discordgo.Session, serverID string) (*discordgo.VoiceState, error) {
	vc, err := ds.State.VoiceState(serverID, ds.State.User.ID)
	if err != nil {
		return nil, errors.Wrap(err, "Error getting bot voice state")
	}

	return vc, nil
}

// gets a list of users in a the channel the bot is in (excluding the bot)
func GetVCUsers(s *discordgo.Session, serverID, channelID string) ([]*discordgo.VoiceState, error) {
	// Get the Guild object
	guild, err := s.State.Guild(serverID)
	if err != nil {
		return nil, errors.Wrap(err, "Error getting guild")
	}

	// Find the VoiceChannel object
	voiceChannel, err := s.State.Channel(channelID)
	if err != nil {
		return nil, errors.Wrap(err, "Error getting voice channel")
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

func UserServerVC(s *discordgo.Session, serverID, userID string) (string, error) {
	vs, err := s.State.VoiceState(serverID, userID)
	if err != nil {
		return "", errors.Wrap(err, "Error getting voice state")
	}
	if vs == nil {
		return "", UserNotInVC
	}

	return vs.ChannelID, nil
}

func UserVC(s *discordgo.Session, userID string) (*discordgo.VoiceState, error) {
	for _, guild := range s.State.Guilds {
		vs, err := s.State.VoiceState(guild.ID, userID)
		if err != nil {
			return nil, errors.Wrap(err, "Error getting voice state")
		}
		if vs == nil {
			continue
		}

		return vs, nil
	}

	return nil, BotNotInVC
}

func UserInBotVC(s *discordgo.Session, userID string) (bool, error) {
	for _, guild := range s.State.Guilds {
		vs, err := s.State.VoiceState(guild.ID, userID)
		if err != nil {
			return false, errors.Wrap(err, "Error getting voice state")
		}
		if vs == nil {
			continue
		}

		botVC, err := GetBotVC(s, guild.ID)
		if err != nil {
			return false, errors.Wrap(err, "Error getting bot voice state")
		}
		if botVC == nil {
			continue
		}

		return botVC.ChannelID == vs.ChannelID, nil
	}

	return false, nil
}
