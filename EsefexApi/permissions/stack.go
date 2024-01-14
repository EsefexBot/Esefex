package permissions

import (
	"esefexapi/opt"
	"esefexapi/types"
	"slices"
)

type PermissionStack struct {
	User    map[types.UserID]Permissions
	Role    map[types.RoleID]Permissions
	Channel map[types.ChannelID]Permissions
}

func NewPermissionStack() PermissionStack {
	return PermissionStack{
		User:    make(map[types.UserID]Permissions),
		Role:    make(map[types.RoleID]Permissions),
		Channel: make(map[types.ChannelID]Permissions),
	}
}

func (ps *PermissionStack) GetUser(userID types.UserID) Permissions {
	if u, ok := ps.User[userID]; ok {
		return u
	}
	return NewUnset()
}

func (ps *PermissionStack) GetRole(roleID types.RoleID) Permissions {
	if r, ok := ps.Role[roleID]; ok {
		return r
	}
	return NewUnset()
}

func (ps *PermissionStack) GetChannel(channelID types.ChannelID) Permissions {
	if c, ok := ps.Channel[channelID]; ok {
		return c
	}
	return NewUnset()
}

func (ps *PermissionStack) SetUser(user types.UserID, p Permissions) {
	ps.User[user] = p
}

func (ps *PermissionStack) SetRole(role types.RoleID, p Permissions) {
	ps.Role[role] = p
}

func (ps *PermissionStack) SetChannel(channel types.ChannelID, p Permissions) {
	ps.Channel[channel] = p
}

func (ps *PermissionStack) UnsetUser(user types.UserID) {
	delete(ps.User, user)
}

func (ps *PermissionStack) UnsetRole(role types.RoleID) {
	delete(ps.Role, role)
}

func (ps *PermissionStack) UnsetChannel(channel types.ChannelID) {
	delete(ps.Channel, channel)
}

func (ps *PermissionStack) UpdateUser(user types.UserID, p Permissions) {
	if _, ok := ps.User[user]; !ok {
		ps.User[user] = NewUnset()
	}

	ps.User[user] = ps.User[user].MergeParent(p)

	ps.clean()
}

func (ps *PermissionStack) UpdateRole(role types.RoleID, p Permissions) {
	if _, ok := ps.Role[role]; !ok {
		ps.Role[role] = NewUnset()
	}

	ps.Role[role] = ps.Role[role].MergeParent(p)

	ps.clean()
}

func (ps *PermissionStack) UpdateChannel(channel types.ChannelID, p Permissions) {
	if _, ok := ps.Channel[channel]; !ok {
		ps.Channel[channel] = NewUnset()
	}

	ps.Channel[channel] = ps.Channel[channel].MergeParent(p)

	ps.clean()
}

// clean removes all permissions that are just unset.
func (ps *PermissionStack) clean() {
	for user, p := range ps.User {
		if p == NewUnset() {
			delete(ps.User, user)
		}
	}

	for role, p := range ps.Role {
		if p == NewUnset() {
			delete(ps.Role, role)
		}
	}

	for channel, p := range ps.Channel {
		if p == NewUnset() {
			delete(ps.Channel, channel)
		}
	}
}

// Query returns the permission state for a given user, role, and channel by merging them together.
// The order of precedence is user > channel > role.
// This means that if a user has a permission set, it will override the channel and role permissions.
// If a channel has a permission set, it will override the role permissions.
// roles is a list of roles that the user has in order of precedence.
func (ps *PermissionStack) Query(user types.UserID, roles []types.RoleID, channel opt.Option[types.ChannelID]) Permissions {
	userPS := ps.User[user]

	slices.Reverse(roles)
	rolesPS := NewUnset()
	for _, role := range roles {
		r := ps.Role[role]
		rolesPS = rolesPS.MergeParent(r)
	}

	var channelPS Permissions = NewUnset()
	if channel.IsSome() {
		channelPS = ps.Channel[channel.Unwrap()]
	}

	return NewUnset().MergeParent(rolesPS).MergeParent(channelPS).MergeParent(userPS)

}
