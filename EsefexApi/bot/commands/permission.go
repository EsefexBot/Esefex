package commands

import (
	"esefexapi/permissions"
	"esefexapi/types"
	"esefexapi/util/dcgoutil"
	"esefexapi/util/refl"
	"fmt"
	"log"

	"github.com/bwmarrin/discordgo"
	"github.com/pkg/errors"
)

var PermissionCommand = &discordgo.ApplicationCommand{
	Name:        "permission",
	Description: "All commands related to permissions.",
	Options: []*discordgo.ApplicationCommandOption{
		{
			Name:        "set",
			Description: "Set a permission for a user, a channel or a role.",
			Type:        discordgo.ApplicationCommandOptionSubCommand,
			Options: []*discordgo.ApplicationCommandOption{
				{
					Name:        "user-role-channel",
					Description: "The user, role or channel to set the permission for.",
					Type:        discordgo.ApplicationCommandOptionString,
					Required:    true,
				},
				{
					Name:        "permission",
					Description: "The permission to set.",
					Type:        discordgo.ApplicationCommandOptionString,
					Required:    true,
					Choices:     getPathOptions(),
				},
				{
					Name:        "value",
					Description: "The value to set the permission to. (Allow,Deny or Unset)",
					Type:        discordgo.ApplicationCommandOptionString,
					Required:    true,
					Choices: []*discordgo.ApplicationCommandOptionChoice{
						{
							Name:  "Allow",
							Value: "Allow",
						},
						{
							Name:  "Deny",
							Value: "Deny",
						},
						{
							Name:  "Unset",
							Value: "Unset",
						},
					},
				},
			},
		},
		{
			Name:        "get",
			Description: "Get the value of a permission for a user, a channel or a role.",
			Type:        discordgo.ApplicationCommandOptionSubCommand,
			Options: []*discordgo.ApplicationCommandOption{
				{
					Name:        "user-role-channel",
					Description: "The user, role or channel to get the permission for.",
					Type:        discordgo.ApplicationCommandOptionString,
					Required:    true,
				},
				{
					Name:        "permission",
					Description: "The permission to get.",
					Type:        discordgo.ApplicationCommandOptionString,
					Required:    true,
					Choices:     getPathOptions(),
				},
			},
		},
		{
			Name:        "clear",
			Description: "Clear all permissions for a user, a channel or a role.",
			Type:        discordgo.ApplicationCommandOptionSubCommand,
			Options: []*discordgo.ApplicationCommandOption{
				{
					Name:        "user-role-channel",
					Description: "The user, role or channel to clear the permissions for.",
					Type:        discordgo.ApplicationCommandOptionString,
					Required:    true,
				},
			},
		},
		{
			Name:        "list",
			Description: "List all permissions for a user, a channel or a role.",
			Type:        discordgo.ApplicationCommandOptionSubCommand,
			Options: []*discordgo.ApplicationCommandOption{
				{
					Name:        "user-role-channel",
					Description: "The user, role or channel to list the permissions for.",
					Type:        discordgo.ApplicationCommandOptionString,
					Required:    true,
				},
			},
		},
		{
			Name:        "list-all",
			Description: "List all permissions for all users, channels and roles.",
			Type:        discordgo.ApplicationCommandOptionSubCommand,
		},
	},
}

func (c *CommandHandlers) Permission(s *discordgo.Session, i *discordgo.InteractionCreate) (*discordgo.InteractionResponse, error) {
	switch i.ApplicationCommandData().Options[0].Name {
	case "set":
		return c.PermissionSet(s, i)
	case "get":
		return c.PermissionGet(s, i)
	case "clear":
		return c.PermissionClear(s, i)
	case "list":
		return c.PermissionList(s, i)
	case "list-all":
		return c.PermissionListAll(s, i)
	default:
		return nil, errors.Wrap(fmt.Errorf("Unknown subcommand %s", i.ApplicationCommandData().Options[0].Name), "Error handling user command")
	}
}

// TODO: Fix the race condition that might occur here
func (c *CommandHandlers) PermissionSet(s *discordgo.Session, i *discordgo.InteractionCreate) (*discordgo.InteractionResponse, error) {
	id := fmt.Sprintf("%v", i.ApplicationCommandData().Options[0].Options[0].Value)
	ty, err := extractTypeFromString(s, types.GuildID(i.GuildID), id)
	if err != nil {
		return nil, errors.Wrap(err, "Error extracting type from string")
	}

	p, err := getPermissions(s, c.dbs, types.GuildID(i.GuildID), id)
	if err != nil {
		return nil, errors.Wrap(err, "Error getting permissions")
	}

	ps := permissions.PSFromString(fmt.Sprintf("%v", i.ApplicationCommandData().Options[0].Options[2].Value))

	ppath := fmt.Sprintf("%v", i.ApplicationCommandData().Options[0].Options[1].Value)

	err = refl.SetNestedFieldValue(&p, ppath, ps)
	if err != nil {
		return nil, errors.Wrap(err, "Error setting nested field value")
	}

	guildID := types.GuildID(i.GuildID)

	log.Printf("Permissiondb is nul? %v", c.dbs.PermissionDB == nil)

	switch ty.PermissionType {
	case permissions.User:
		err = c.dbs.PermissionDB.UpdateUser(guildID, types.UserID(ty.ID), p)
	case permissions.Role:
		err = c.dbs.PermissionDB.UpdateRole(guildID, types.RoleID(ty.ID), p)
	case permissions.Channel:
		err = c.dbs.PermissionDB.UpdateChannel(guildID, types.ChannelID(ty.ID), p)
	}

	if err != nil {
		return nil, errors.Wrap(err, "Error setting permissions")
	}

	return &discordgo.InteractionResponse{
		Type: discordgo.InteractionResponseChannelMessageWithSource,
		Data: &discordgo.InteractionResponseData{
			Content: fmt.Sprintf("Set %s for %s to %s", ppath, id, ps.String()),
		},
	}, nil
}

// TODO: Better alignment for the list all command (maybe use a table?)
func (c *CommandHandlers) PermissionListAll(s *discordgo.Session, i *discordgo.InteractionCreate) (*discordgo.InteractionResponse, error) {

	guildID := types.GuildID(i.GuildID)
	tabl, err := c.dbs.PermissionDB.GetGuild(guildID).FmtStack(s, guildID)
	if err != nil {
		return nil, errors.Wrap(err, "Error formatting permissions")
	}

	resp := fmt.Sprintf("Permissions for all users, channels and roles:```%s```", tabl)

	return &discordgo.InteractionResponse{
		Type: discordgo.InteractionResponseChannelMessageWithSource,
		Data: &discordgo.InteractionResponseData{
			Content: resp,
		},
	}, nil
}

func (c *CommandHandlers) PermissionGet(s *discordgo.Session, i *discordgo.InteractionCreate) (*discordgo.InteractionResponse, error) {
	id := fmt.Sprintf("%v", i.ApplicationCommandData().Options[0].Options[0].Value)

	p, err := getPermissions(s, c.dbs, types.GuildID(i.GuildID), id)
	if err != nil {
		return nil, errors.Wrap(err, "Error getting permissions")
	}

	ppath := fmt.Sprintf("%v", i.ApplicationCommandData().Options[0].Options[1].Value)

	ps, err := getPermission(p, ppath)
	if err != nil {
		return nil, errors.Wrap(err, "Error getting permission")
	}

	ty, err := extractTypeFromString(s, types.GuildID(i.GuildID), id)
	if err != nil {
		return nil, errors.Wrap(err, "Error extracting type from string")
	}

	var name string
	switch ty.PermissionType {
	case permissions.User:
		name, err = dcgoutil.UserIDName(s, types.UserID(ty.ID))
	case permissions.Role:
		name, err = dcgoutil.RoleIDName(s, types.GuildID(i.GuildID), types.RoleID(ty.ID))
	case permissions.Channel:
		name, err = dcgoutil.ChannelIDMention(s, types.GuildID(i.GuildID), types.ChannelID(ty.ID))
	}
	if err != nil {
		return nil, errors.Wrap(err, "Error getting name")
	}

	// TODO: add name display here
	resp := fmt.Sprintf("%s for %s: %s", ppath, name, ps.String())
	return &discordgo.InteractionResponse{
		Type: discordgo.InteractionResponseChannelMessageWithSource,
		Data: &discordgo.InteractionResponseData{
			Content: resp,
		},
	}, nil
}

// TODO: Fix this command (it is not clearing permissions)
func (c *CommandHandlers) PermissionClear(s *discordgo.Session, i *discordgo.InteractionCreate) (*discordgo.InteractionResponse, error) {
	id := fmt.Sprintf("%v", i.ApplicationCommandData().Options[0].Options[0].Value)
	ty, err := extractTypeFromString(s, types.GuildID(i.GuildID), id)
	if err != nil {
		return nil, errors.Wrap(err, "Error extracting type from string")
	}

	var name string
	switch ty.PermissionType {
	case permissions.User:
		name, err = dcgoutil.UserIDName(s, types.UserID(ty.ID))
	case permissions.Role:
		name, err = dcgoutil.RoleIDName(s, types.GuildID(i.GuildID), types.RoleID(ty.ID))
	case permissions.Channel:
		name, err = dcgoutil.ChannelIDMention(s, types.GuildID(i.GuildID), types.ChannelID(ty.ID))
	}
	if err != nil {
		return nil, errors.Wrap(err, "Error getting name")
	}

	guildID := types.GuildID(i.GuildID)

	switch ty.PermissionType {
	case permissions.User:
		err = c.dbs.PermissionDB.UpdateUser(guildID, types.UserID(ty.ID), permissions.NewUnset())
	case permissions.Role:
		err = c.dbs.PermissionDB.UpdateRole(guildID, types.RoleID(ty.ID), permissions.NewUnset())
	case permissions.Channel:
		err = c.dbs.PermissionDB.UpdateChannel(guildID, types.ChannelID(ty.ID), permissions.NewUnset())
	}
	if err != nil {
		return nil, errors.Wrap(err, "Error clearing permissions")
	}

	return &discordgo.InteractionResponse{
		Type: discordgo.InteractionResponseChannelMessageWithSource,
		Data: &discordgo.InteractionResponseData{
			Content: fmt.Sprintf("Cleared permissions for %s %s", ty.PermissionType, name),
		},
	}, nil
}

func (c *CommandHandlers) PermissionList(s *discordgo.Session, i *discordgo.InteractionCreate) (*discordgo.InteractionResponse, error) {
	id := fmt.Sprintf("%v", i.ApplicationCommandData().Options[0].Options[0].Value)

	p, err := getPermissions(s, c.dbs, types.GuildID(i.GuildID), id)
	if err != nil {
		return nil, errors.Wrap(err, "Error getting permissions")
	}

	pstr, err := formatPermissions(p)
	if err != nil {
		return nil, errors.Wrap(err, "Error formatting permissions")
	}

	ty, err := extractTypeFromString(s, types.GuildID(i.GuildID), id)
	if err != nil {
		return nil, errors.Wrap(err, "Error extracting type from string")
	}

	var name string
	switch ty.PermissionType {
	case permissions.User:
		name, err = dcgoutil.UserIDName(s, types.UserID(ty.ID))
	case permissions.Role:
		name, err = dcgoutil.RoleIDName(s, types.GuildID(i.GuildID), types.RoleID(ty.ID))
	case permissions.Channel:
		name, err = dcgoutil.ChannelIDMention(s, types.GuildID(i.GuildID), types.ChannelID(ty.ID))
	}
	if err != nil {
		return nil, errors.Wrap(err, "Error getting name")
	}

	resp := fmt.Sprintf("Permissions for %s:\n", name)
	resp += pstr

	return &discordgo.InteractionResponse{
		Type: discordgo.InteractionResponseChannelMessageWithSource,
		Data: &discordgo.InteractionResponseData{
			Content: resp,
		},
	}, nil
}
