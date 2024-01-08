package permissions

import (
	"esefexapi/opt"
	"esefexapi/types"
	"slices"
)

type PermissionStack struct {
	user    map[types.UserID]Permissions
	role    map[types.RoleID]Permissions
	channel map[types.ChannelID]Permissions
}

func NewPermissionStack() *PermissionStack {
	return &PermissionStack{
		user:    make(map[types.UserID]Permissions),
		role:    make(map[types.RoleID]Permissions),
		channel: make(map[types.ChannelID]Permissions),
	}
}

func (ps *PermissionStack) GetUser(userID types.UserID) Permissions {
	if u, ok := ps.user[userID]; ok {
		return u
	}
	return NewUnset()
}

func (ps *PermissionStack) GetRole(roleID types.RoleID) Permissions {
	if r, ok := ps.role[roleID]; ok {
		return r
	}
	return NewUnset()
}

func (ps *PermissionStack) GetChannel(channelID types.ChannelID) Permissions {
	if c, ok := ps.channel[channelID]; ok {
		return c
	}
	return NewUnset()
}

func (ps *PermissionStack) SetUser(user types.UserID, p Permissions) {
	ps.user[user] = p
}

func (ps *PermissionStack) SetRole(role types.RoleID, p Permissions) {
	ps.role[role] = p
}

func (ps *PermissionStack) SetChannel(channel types.ChannelID, p Permissions) {
	ps.channel[channel] = p
}

func (ps *PermissionStack) UnsetUser(user types.UserID) {
	delete(ps.user, user)
}

func (ps *PermissionStack) UnsetRole(role types.RoleID) {
	delete(ps.role, role)
}

func (ps *PermissionStack) UnsetChannel(channel types.ChannelID) {
	delete(ps.channel, channel)
}

func (ps *PermissionStack) UpdateUser(user types.UserID, p Permissions) {
	if _, ok := ps.user[user]; !ok {
		ps.user[user] = NewUnset()
	}

	ps.user[user] = ps.user[user].MergeParent(p)
}

func (ps *PermissionStack) UpdateRole(role types.RoleID, p Permissions) {
	if _, ok := ps.role[role]; !ok {
		ps.role[role] = NewUnset()
	}

	ps.role[role] = ps.role[role].MergeParent(p)
}

func (ps *PermissionStack) UpdateChannel(channel types.ChannelID, p Permissions) {
	if _, ok := ps.channel[channel]; !ok {
		ps.channel[channel] = NewUnset()
	}

	ps.channel[channel] = ps.channel[channel].MergeParent(p)
}

// Query returns the permission state for a given user, role, and channel by merging them together.
// The order of precedence is user > channel > role.
// This means that if a user has a permission set, it will override the channel and role permissions.
// If a channel has a permission set, it will override the role permissions.
// roles is a list of roles that the user has in order of precedence.
func (ps *PermissionStack) Query(user types.UserID, roles []types.RoleID, channel opt.Option[types.ChannelID]) Permissions {
	userPS := ps.user[user]

	slices.Reverse(roles)
	rolesPS := NewUnset()
	for _, role := range roles {
		r := ps.role[role]
		rolesPS = rolesPS.MergeParent(r)
	}

	var channelPS Permissions = NewUnset()
	if channel.IsSome() {
		channelPS = ps.channel[channel.Unwrap()]
	}

	return NewUnset().MergeParent(rolesPS).MergeParent(channelPS).MergeParent(userPS)

}
