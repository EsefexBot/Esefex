package commands

import (
	"fmt"

	"github.com/bwmarrin/discordgo"
	"github.com/davecgh/go-spew/spew"
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
	spew.Dump(i.ApplicationCommandData().Options)

	return nil, errors.Wrap(fmt.Errorf("not implemented"), "permission")
}
