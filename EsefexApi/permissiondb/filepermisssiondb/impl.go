package filepermisssiondb

import (
	"esefexapi/permissions"
	"esefexapi/types"
)

// GetChannel implements permissiondb.PermissionDB.
func (f *FilePermissionDB) GetChannel(guildID types.GuildID, channelID types.ChannelID) (permissions.Permissions, error) {
	f.rw.RLock()
	defer f.rw.RUnlock()

	return f.ensureGuild(guildID).GetChannel(channelID), nil
}

// GetRole implements permissiondb.PermissionDB.
func (f *FilePermissionDB) GetRole(guildID types.GuildID, roleID types.RoleID) (permissions.Permissions, error) {
	f.rw.RLock()
	defer f.rw.RUnlock()

	return f.ensureGuild(guildID).GetRole(roleID), nil
}

// GetUser implements permissiondb.PermissionDB.
func (f *FilePermissionDB) GetUser(guildID types.GuildID, userID types.UserID) (permissions.Permissions, error) {
	f.rw.RLock()
	defer f.rw.RUnlock()

	return f.ensureGuild(guildID).GetUser(userID), nil
}

// UpdateChannel implements permissiondb.PermissionDB.
func (f *FilePermissionDB) UpdateChannel(guildID types.GuildID, channelID types.ChannelID, p permissions.Permissions) error {
	f.rw.Lock()
	defer f.rw.Unlock()

	f.ensureGuild(guildID).UpdateChannel(channelID, p)
	go f.save()
	return nil
}

// UpdateRole implements permissiondb.PermissionDB.
func (f *FilePermissionDB) UpdateRole(guildID types.GuildID, roleID types.RoleID, p permissions.Permissions) error {
	f.rw.Lock()
	defer f.rw.Unlock()

	f.ensureGuild(guildID).UpdateRole(roleID, p)
	go f.save()
	return nil
}

// UpdateUser implements permissiondb.PermissionDB.
func (f *FilePermissionDB) UpdateUser(guildID types.GuildID, userID types.UserID, p permissions.Permissions) error {
	f.rw.Lock()
	defer f.rw.Unlock()

	f.ensureGuild(guildID).UpdateUser(userID, p)
	go f.save()
	return nil
}

func (f *FilePermissionDB) GetUsers(guildID types.GuildID) ([]types.UserID, error) {
	f.rw.RLock()
	defer f.rw.RUnlock()

	var users []types.UserID
	for k := range f.ensureGuild(guildID).Users {
		users = append(users, k)
	}

	return users, nil
}

func (f *FilePermissionDB) GetRoles(guildID types.GuildID) ([]types.RoleID, error) {
	f.rw.RLock()
	defer f.rw.RUnlock()

	var roles []types.RoleID
	for k := range f.ensureGuild(guildID).Roles {
		roles = append(roles, k)
	}

	return roles, nil
}

func (f *FilePermissionDB) GetChannels(guildID types.GuildID) ([]types.ChannelID, error) {
	f.rw.RLock()
	defer f.rw.RUnlock()

	var channels []types.ChannelID
	for k := range f.ensureGuild(guildID).Channels {
		channels = append(channels, k)
	}

	return channels, nil
}
