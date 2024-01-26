package permissiondb

import (
	"esefexapi/permissions"
	"esefexapi/types"
)

type PermissionDB interface {
	GetUser(guild types.GuildID, userID types.UserID) (permissions.Permissions, error)
	GetRole(guild types.GuildID, roleID types.RoleID) (permissions.Permissions, error)
	GetChannel(guild types.GuildID, channelID types.ChannelID) (permissions.Permissions, error)
	UpdateUser(guild types.GuildID, userID types.UserID, p permissions.Permissions) error
	UpdateRole(guild types.GuildID, roleID types.RoleID, p permissions.Permissions) error
	UpdateChannel(guild types.GuildID, channelID types.ChannelID, p permissions.Permissions) error
	GetUsers(guild types.GuildID) ([]types.UserID, error)
	GetRoles(guild types.GuildID) ([]types.RoleID, error)
	GetChannels(guild types.GuildID) ([]types.ChannelID, error)
	Query(guild types.GuildID, user types.UserID) (permissions.Permissions, error)
}
