package dcgoutil

import (
	// "log"

	"esefexapi/opt"

	"github.com/bwmarrin/discordgo"
	// "github.com/davecgh/go-spew/spew"
	"github.com/pkg/errors"
)

// util functions for discordgo

// checks if the bot is in a voice channel in a server
func BotInVC(ds *discordgo.Session, serverID, channelID string) (bool, error) {
	vs, err := ds.State.VoiceState(serverID, ds.State.User.ID)
	if err == discordgo.ErrStateNotFound {
		return false, nil
	} else if err != nil {
		return false, errors.Wrap(err, "Error getting bot voice state")
	}
	return vs.ChannelID == channelID, nil
}

// gets the voice state of the bot in a server
func GetBotVC(ds *discordgo.Session, serverID string) (opt.Option[*discordgo.VoiceState], error) {
	vc, err := ds.State.VoiceState(serverID, ds.State.User.ID)
	if err == discordgo.ErrStateNotFound {
		return opt.None[*discordgo.VoiceState](), nil
	} else if err != nil {
		return opt.None[*discordgo.VoiceState](), errors.Wrap(err, "Error getting bot voice state")
	}

	return opt.Some(vc), nil
}

// gets a list of users in a the channel the bot is in
func GetVCUsers(ds *discordgo.Session, serverID, channelID string) ([]*discordgo.VoiceState, error) {
	// Get the Guild object
	guild, err := ds.State.Guild(serverID)
	if err != nil {
		return nil, errors.Wrap(err, "Error getting guild")
	}

	// Find the VoiceChannel object
	voiceChannel, err := ds.State.Channel(channelID)
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

// gets the voice state of a user in a server
func UserServerVC(ds *discordgo.Session, serverID, userID string) (opt.Option[*discordgo.VoiceState], error) {
	vs, err := ds.State.VoiceState(serverID, userID)
	if err == discordgo.ErrStateNotFound {
		return opt.None[*discordgo.VoiceState](), nil
	} else if err != nil {
		return opt.None[*discordgo.VoiceState](), errors.Wrap(err, "Error getting voice state")
	}

	return opt.Some(vs), nil
}

// gets the voice state of a user in any server
func UserVC(ds *discordgo.Session, userID string) (opt.Option[*discordgo.VoiceState], error) {
	for _, guild := range ds.State.Guilds {
		vs, err := ds.State.VoiceState(guild.ID, userID)
		if err == discordgo.ErrStateNotFound {
			continue
		} else if err != nil {
			return opt.None[*discordgo.VoiceState](), errors.Wrap(err, "Error getting voice state")
		}

		return opt.Some(vs), nil
	}

	return opt.None[*discordgo.VoiceState](), nil
}

// checks if a user is in the same voice channel as the bot
func UserInBotVC(ds *discordgo.Session, userID string) (bool, error) {
	for _, guild := range ds.State.Guilds {
		vs, err := ds.State.VoiceState(guild.ID, userID)
		if err == discordgo.ErrStateNotFound {
			continue
		} else if err != nil {
			return false, errors.Wrap(err, "Error getting voice state")
		}

		botVCopt, err := GetBotVC(ds, guild.ID)
		if err != nil {
			return false, errors.Wrap(err, "Error getting bot voice state")
		} else if botVCopt.IsNone() {
			continue
		}
		botVC := botVCopt.Unwrap()

		return botVC.ChannelID == vs.ChannelID, nil
	}

	return false, nil
}

// gets the server a user is connected to (if any)
func UserServer(ds *discordgo.Session, userID string) (opt.Option[*discordgo.Guild], error) {
	Ochan, err := UserVC(ds, userID)
	if err != nil {
		return opt.None[*discordgo.Guild](), errors.Wrap(err, "Error getting user voice state")
	} else if Ochan.IsNone() {
		return opt.None[*discordgo.Guild](), nil
	}

	vs := Ochan.Unwrap()

	guild, err := ds.State.Guild(vs.GuildID)
	if err != nil {
		return opt.None[*discordgo.Guild](), errors.Wrap(err, "Error getting guild")
	}

	return opt.Some(guild), nil
}
