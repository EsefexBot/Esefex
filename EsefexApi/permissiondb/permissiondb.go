package permissiondb

import (
	"esefexapi/permissions"
	"esefexapi/types"
)

type PermissionDB interface {
	GetUser(userID types.UserID) (permissions.Permissions, error)
	GetRole(roleID types.RoleID) (permissions.Permissions, error)
	GetChannel(channelID types.ChannelID) (permissions.Permissions, error)
	UpdateUser(userID types.UserID, p permissions.Permissions) error
	UpdateRole(roleID types.RoleID, p permissions.Permissions) error
	UpdateChannel(channelID types.ChannelID, p permissions.Permissions) error
	GetUsers() ([]types.UserID, error)
	GetRoles() ([]types.RoleID, error)
	GetChannels() ([]types.ChannelID, error)
	Query(user types.UserID, guild types.GuildID) (permissions.Permissions, error)
}
