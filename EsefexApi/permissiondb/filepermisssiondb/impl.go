package filepermisssiondb

import (
	"esefexapi/permissions"
	"esefexapi/types"
)

// GetChannel implements permissiondb.PermissionDB.
func (f *FilePermissionDB) GetChannel(channelID types.ChannelID) (permissions.Permissions, error) {
	f.rw.RLock()
	defer f.rw.RUnlock()

	return f.stack.GetChannel(channelID), nil
}

// GetRole implements permissiondb.PermissionDB.
func (f *FilePermissionDB) GetRole(roleID types.RoleID) (permissions.Permissions, error) {
	f.rw.RLock()
	defer f.rw.RUnlock()

	return f.stack.GetRole(roleID), nil
}

// GetUser implements permissiondb.PermissionDB.
func (f *FilePermissionDB) GetUser(userID types.UserID) (permissions.Permissions, error) {
	f.rw.RLock()
	defer f.rw.RUnlock()

	return f.stack.GetUser(userID), nil
}

// UpdateChannel implements permissiondb.PermissionDB.
func (f *FilePermissionDB) UpdateChannel(channelID types.ChannelID, p permissions.Permissions) error {
	f.rw.Lock()
	defer f.rw.Unlock()

	f.stack.UpdateChannel(channelID, p)
	go f.save()
	return nil
}

// UpdateRole implements permissiondb.PermissionDB.
func (f *FilePermissionDB) UpdateRole(roleID types.RoleID, p permissions.Permissions) error {
	f.rw.Lock()
	defer f.rw.Unlock()

	f.stack.UpdateRole(roleID, p)
	go f.save()
	return nil
}

// UpdateUser implements permissiondb.PermissionDB.
func (f *FilePermissionDB) UpdateUser(userID types.UserID, p permissions.Permissions) error {
	f.rw.Lock()
	defer f.rw.Unlock()

	f.stack.UpdateUser(userID, p)
	go f.save()
	return nil
}

func (f *FilePermissionDB) GetUsers() ([]types.UserID, error) {
	f.rw.RLock()
	defer f.rw.RUnlock()

	var users []types.UserID
	for k := range f.stack.User {
		users = append(users, k)
	}

	return users, nil
}

func (f *FilePermissionDB) GetRoles() ([]types.RoleID, error) {
	f.rw.RLock()
	defer f.rw.RUnlock()

	var roles []types.RoleID
	for k := range f.stack.Role {
		roles = append(roles, k)
	}

	return roles, nil
}

func (f *FilePermissionDB) GetChannels() ([]types.ChannelID, error) {
	f.rw.RLock()
	defer f.rw.RUnlock()

	var channels []types.ChannelID
	for k := range f.stack.Channel {
		channels = append(channels, k)
	}

	return channels, nil
}
