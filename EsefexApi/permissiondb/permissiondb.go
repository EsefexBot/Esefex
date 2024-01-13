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
}
