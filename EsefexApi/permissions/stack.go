package permissions

import (
	"esefexapi/opt"
	"esefexapi/types"
	"slices"
)

type PermissionStack struct {
	Users    map[types.UserID]Permissions
	Roles    map[types.RoleID]Permissions
	Channels map[types.ChannelID]Permissions
}

func NewPermissionStack() PermissionStack {
	return PermissionStack{
		Users:    make(map[types.UserID]Permissions),
		Roles:    make(map[types.RoleID]Permissions),
		Channels: make(map[types.ChannelID]Permissions),
	}
}

func (ps *PermissionStack) GetUser(userID types.UserID) Permissions {
	if u, ok := ps.Users[userID]; ok {
		return u
	}
	return NewUnset()
}

func (ps *PermissionStack) GetRole(roleID types.RoleID) Permissions {
	if r, ok := ps.Roles[roleID]; ok {
		return r
	}
	return NewUnset()
}

func (ps *PermissionStack) GetChannel(channelID types.ChannelID) Permissions {
	if c, ok := ps.Channels[channelID]; ok {
		return c
	}
	return NewUnset()
}

func (ps *PermissionStack) SetUser(user types.UserID, p Permissions) {
	ps.Users[user] = p
}

func (ps *PermissionStack) SetRole(role types.RoleID, p Permissions) {
	ps.Roles[role] = p
}

func (ps *PermissionStack) SetChannel(channel types.ChannelID, p Permissions) {
	ps.Channels[channel] = p
}

func (ps *PermissionStack) UnsetUser(user types.UserID) {
	delete(ps.Users, user)
}

func (ps *PermissionStack) UnsetRole(role types.RoleID) {
	delete(ps.Roles, role)
}

func (ps *PermissionStack) UnsetChannel(channel types.ChannelID) {
	delete(ps.Channels, channel)
}

func (ps *PermissionStack) UpdateUser(user types.UserID, p Permissions) {
	if _, ok := ps.Users[user]; !ok {
		ps.Users[user] = NewUnset()
	}

	ps.Users[user] = ps.Users[user].MergeParent(p)

	ps.clean()
}

func (ps *PermissionStack) UpdateRole(role types.RoleID, p Permissions) {
	if _, ok := ps.Roles[role]; !ok {
		ps.Roles[role] = NewUnset()
	}

	ps.Roles[role] = ps.Roles[role].MergeParent(p)

	ps.clean()
}

func (ps *PermissionStack) UpdateChannel(channel types.ChannelID, p Permissions) {
	if _, ok := ps.Channels[channel]; !ok {
		ps.Channels[channel] = NewUnset()
	}

	ps.Channels[channel] = ps.Channels[channel].MergeParent(p)

	ps.clean()
}

// clean removes all permissions that are just unset.
func (ps *PermissionStack) clean() {
	for user, p := range ps.Users {
		if p == NewUnset() {
			delete(ps.Users, user)
		}
	}

	for role, p := range ps.Roles {
		if p == NewUnset() {
			delete(ps.Roles, role)
		}
	}

	for channel, p := range ps.Channels {
		if p == NewUnset() {
			delete(ps.Channels, channel)
		}
	}
}

// Query returns the permission state for a given user, role, and channel by merging them together.
// The order of precedence is user > channel > role.
// This means that if a user has a permission set, it will override the channel and role permissions.
// If a channel has a permission set, it will override the role permissions.
// roles is a list of roles that the user has in order of precedence.
func (ps *PermissionStack) Query(user types.UserID, roles []types.RoleID, channel opt.Option[types.ChannelID]) Permissions {
	userPS := ps.Users[user]

	slices.Reverse(roles)
	rolesPS := NewUnset()
	for _, role := range roles {
		r := ps.Roles[role]
		rolesPS = rolesPS.MergeParent(r)
	}

	var channelPS Permissions = NewUnset()
	if channel.IsSome() {
		channelPS = ps.Channels[channel.Unwrap()]
	}

	return NewUnset().MergeParent(rolesPS).MergeParent(channelPS).MergeParent(userPS)

}
