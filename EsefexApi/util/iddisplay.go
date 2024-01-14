package util

import (
	"esefexapi/types"

	"github.com/bwmarrin/discordgo"
	"github.com/pkg/errors"
)

func UserIDName(ds *discordgo.Session, u types.UserID) (string, error) {
	user, err := ds.User(u.String())
	if err != nil {
		return "", errors.Wrap(err, "Error getting user")
	}
	return user.Username, nil
}

func RoleIDName(ds *discordgo.Session, g types.GuildID, r types.RoleID) (string, error) {
	if r == "everyone" {
		return "everyone", nil
	}

	roles, err := ds.GuildRoles(g.String())
	if err != nil {
		return "", errors.Wrap(err, "Error getting roles")
	}

	for _, role := range roles {
		if role.ID == r.String() {
			return role.Name, nil
		}
	}

	return "", errors.Wrap(err, "Error getting role")
}

func ChannelIDName(ds *discordgo.Session, g types.GuildID, c types.ChannelID) (string, error) {
	channels, err := ds.GuildChannels(g.String())
	if err != nil {
		return "", errors.Wrap(err, "Error getting channels")
	}

	for _, channel := range channels {
		if channel.ID == c.String() {
			return channel.Name, nil
		}
	}

	return "", errors.Wrap(err, "Error getting channel")
}

func UserIDMention(ds *discordgo.Session, u types.UserID) (string, error) {
	user, err := ds.User(u.String())
	if err != nil {
		return "", errors.Wrap(err, "Error getting user")
	}
	return user.Mention(), nil
}

func RoleIDMention(ds *discordgo.Session, g types.GuildID, r types.RoleID) (string, error) {
	if r == "everyone" {
		return "@everyone", nil
	}

	roles, err := ds.GuildRoles(g.String())
	if err != nil {
		return "", errors.Wrap(err, "Error getting roles")
	}

	for _, role := range roles {
		if role.ID == r.String() {
			return role.Mention(), nil
		}
	}

	return "", errors.Wrap(err, "Error getting role")
}

func ChannelIDMention(ds *discordgo.Session, g types.GuildID, c types.ChannelID) (string, error) {
	channels, err := ds.GuildChannels(g.String())
	if err != nil {
		return "", errors.Wrap(err, "Error getting channels")
	}

	for _, channel := range channels {
		if channel.ID == c.String() {
			return channel.Mention(), nil
		}
	}

	return "", errors.Wrap(err, "Error getting channel")
}
