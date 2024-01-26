package dcgoutil

import (
	// "log"

	"esefexapi/opt"
	"esefexapi/types"

	"github.com/bwmarrin/discordgo"

	// "github.com/davecgh/go-spew/spew"
	"github.com/pkg/errors"
)

// util functions for discordgo

// checks if the bot is in a voice channel in a guild
func BotInVC(ds *discordgo.Session, guildID, channelID string) (bool, error) {
	vs, err := ds.State.VoiceState(guildID, ds.State.User.ID)
	if err == discordgo.ErrStateNotFound {
		return false, nil
	} else if err != nil {
		return false, errors.Wrap(err, "Error getting bot voice state")
	}
	return vs.ChannelID == channelID, nil
}

// gets the voice state of the bot in a guild
func GetBotVC(ds *discordgo.Session, guildID types.GuildID) (opt.Option[*discordgo.VoiceState], error) {
	vc, err := ds.State.VoiceState(guildID.String(), ds.State.User.ID)
	if err == discordgo.ErrStateNotFound {
		return opt.None[*discordgo.VoiceState](), nil
	} else if err != nil {
		return opt.None[*discordgo.VoiceState](), errors.Wrap(err, "Error getting bot voice state")
	}

	return opt.Some(vc), nil
}

// gets a list of users in a the channel the bot is in
func GetVCUsers(ds *discordgo.Session, guildID, channelID string) ([]*discordgo.VoiceState, error) {
	// Get the Guild object
	guild, err := ds.State.Guild(guildID)
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

// gets the voice state of a user in a guild
func UserGuildVC(ds *discordgo.Session, guildID types.GuildID, userID types.UserID) (opt.Option[*discordgo.VoiceState], error) {
	vs, err := ds.State.VoiceState(guildID.String(), userID.String())
	if err == discordgo.ErrStateNotFound {
		return opt.None[*discordgo.VoiceState](), nil
	} else if err != nil {
		return opt.None[*discordgo.VoiceState](), errors.Wrap(err, "Error getting voice state")
	}

	return opt.Some(vs), nil
}

// gets the voice state of a user in any guild
func UserVCAny(ds *discordgo.Session, userID types.UserID) (opt.Option[*discordgo.VoiceState], error) {
	for _, guild := range ds.State.Guilds {
		vs, err := ds.State.VoiceState(guild.ID, userID.String())
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

		botVCopt, err := GetBotVC(ds, types.GuildID(guild.ID))
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

// gets the guild a user is connected to (if any)
func UserVCGuild(ds *discordgo.Session, userID types.UserID) (opt.Option[*discordgo.Guild], error) {
	Ochan, err := UserVCAny(ds, userID)
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

// get the list of guilds a user is a member of
func UserGuilds(ds *discordgo.Session, userID types.UserID) ([]*discordgo.Guild, error) {
	guilds := []*discordgo.Guild{}

	for _, guild := range ds.State.Guilds {
		member, err := ds.State.Member(guild.ID, userID.String())
		if err == discordgo.ErrStateNotFound {
			continue
		} else if err != nil {
			return nil, errors.Wrap(err, "Error getting member")
		}

		if member.User.ID == userID.String() {
			guilds = append(guilds, guild)
		}
	}

	return guilds, nil
}

func UserHasPermissions(member *discordgo.Member, perms int64) (bool, error) {
	return member.Permissions&perms == perms, nil
}

func UserIsOwner(ds *discordgo.Session, guildID types.GuildID, userID types.UserID) (bool, error) {
	guild, err := ds.Guild(guildID.String())
	if err != nil {
		return false, errors.Wrap(err, "Error getting guild")
	}

	return guild.OwnerID == userID.String(), nil
}
