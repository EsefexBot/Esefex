package commands

import (
	"esefexapi/permissions"
	"esefexapi/types"
	"fmt"

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
					// TODO: Dynamically get the choices from the permission type on startup using reflection.
					// Choices: []*discordgo.ApplicationCommandOptionChoice{}
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
					// TODO: Dynamically get the choices from the permission type on startup using reflection.
					// Choices: []*discordgo.ApplicationCommandOptionChoice{}
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
	default:
		return nil, errors.Wrap(fmt.Errorf("Unknown subcommand %s", i.ApplicationCommandData().Options[0].Name), "Error handling user command")
	}
}

func (c *CommandHandlers) PermissionSet(s *discordgo.Session, i *discordgo.InteractionCreate) (*discordgo.InteractionResponse, error) {
	return nil, errors.Wrap(fmt.Errorf("Not implemented"), "Error handling user command PermissionSet")
}

func (c *CommandHandlers) PermissionGet(s *discordgo.Session, i *discordgo.InteractionCreate) (*discordgo.InteractionResponse, error) {
	return nil, errors.Wrap(fmt.Errorf("Not implemented"), "Error handling user command PermissionGet")
}

func (c *CommandHandlers) PermissionClear(s *discordgo.Session, i *discordgo.InteractionCreate) (*discordgo.InteractionResponse, error) {
	return nil, errors.Wrap(fmt.Errorf("Not implemented"), "Error handling user command PermissionClear")
}

func (c *CommandHandlers) PermissionList(s *discordgo.Session, i *discordgo.InteractionCreate) (*discordgo.InteractionResponse, error) {
	id := fmt.Sprintf("%v", i.ApplicationCommandData().Options[0].Options[0].Value)
	ty, err := extractTypeFromString(s, types.GuildID(i.GuildID), id)
	if err != nil {
		return nil, errors.Wrap(err, "Error extracting type from string")
	}

	var p permissions.Permissions

	switch ty.PermissionType {
	case permissions.User:
		p, err = c.dbs.PremissionDB.GetUser(types.UserID(ty.ID))
	case permissions.Role:
		p, err = c.dbs.PremissionDB.GetRole(types.RoleID(ty.ID))
	case permissions.Channel:
		p, err = c.dbs.PremissionDB.GetChannel(types.ChannelID(ty.ID))
	}

	if err != nil {
		return nil, errors.Wrap(err, "Error getting permissions")
	}

	ps, err := formatPermissions(p)
	if err != nil {
		return nil, errors.Wrap(err, "Error formatting permissions")
	}

	resp := fmt.Sprintf("Permissions for %s %s:\n", ty.PermissionType, ty.ID)
	resp += ps

	return &discordgo.InteractionResponse{
		Type: discordgo.InteractionResponseChannelMessageWithSource,
		Data: &discordgo.InteractionResponseData{
			Content: resp,
		},
	}, nil
}
