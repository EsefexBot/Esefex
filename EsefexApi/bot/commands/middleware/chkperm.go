package middleware

import (
	"esefexapi/bot/commands/cmdhandler"
	"esefexapi/permissions"
	"esefexapi/types"
	"esefexapi/util/dcgoutil"
	"esefexapi/util/refl"
	"fmt"

	"github.com/bwmarrin/discordgo"
	"github.com/pkg/errors"
)

func (m *CommandMiddleware) CheckPerms(next cmdhandler.CommandHandlerWithErr, perms ...string) cmdhandler.CommandHandlerWithErr {
	return func(s *discordgo.Session, i *discordgo.InteractionCreate) (*discordgo.InteractionResponse, error) {
		// check if user is admin or owner
		hasPerms, err := dcgoutil.UserHasPermissions(i.Member, discordgo.PermissionAdministrator)
		if err != nil {
			return nil, errors.Wrap(err, "Error checking user permissions")
		}

		isOwner, err := dcgoutil.UserIsOwner(s, types.GuildID(i.GuildID), types.UserID(i.Member.User.ID))
		if err != nil {
			return nil, errors.Wrap(err, "Error checking user permissions")
		}

		if hasPerms || isOwner {
			return next(s, i)
		}

		// check if user has permissions
		p, err := m.dbs.PermissionDB.Query(types.GuildID(i.GuildID), types.UserID(i.Member.User.ID))
		if err != nil {
			return nil, errors.Wrap(err, "Error querying permissions")
		}

		for _, permission := range perms {
			ps, err := refl.GetNestedFieldValue(p, permission)
			if err != nil {
				return nil, errors.Wrap(err, "Error getting nested field value")
			}

			if !ps.(permissions.PermissionState).Allowed() {
				return &discordgo.InteractionResponse{
					Type: discordgo.InteractionResponseChannelMessageWithSource,
					Data: &discordgo.InteractionResponseData{
						Content: fmt.Sprintf("You do not have the required permissions to use this command (missing `%s`)", permission),
					},
				}, nil
			}
		}

		return next(s, i)
	}
}
