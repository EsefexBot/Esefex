package filepermisssiondb

import (
	"encoding/json"
	"esefexapi/permissiondb"
	"esefexapi/permissions"
	"esefexapi/types"
	"os"
	"sync"

	"github.com/pkg/errors"
)

var _ permissiondb.PermissionDB = &FilePermissionDB{}

type FilePermissionDB struct {
	filePath string
	rw       *sync.RWMutex
	stack    *permissions.PermissionStack
}

func NewFilePermissionDB(path string) (*FilePermissionDB, error) {
	fpdb := &FilePermissionDB{
		filePath: path,
		rw:       &sync.RWMutex{},
		stack:    permissions.NewPermissionStack(),
	}
	err := fpdb.load()
	return fpdb, err
}

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

// load loads the permission stack from the file.
func (f *FilePermissionDB) load() error {
	file, err := os.Open(f.filePath)
	if err != nil {
		return errors.Wrap(err, "Error opening file")
	}
	defer file.Close()

	err = json.NewDecoder(file).Decode(f.stack)
	if err != nil {
		return errors.Wrap(err, "Error decoding permission stack")
	}

	return nil
}

// save saves the permission stack to the file.
func (f *FilePermissionDB) save() error {
	file, err := os.OpenFile(f.filePath, os.O_RDWR|os.O_CREATE, os.ModePerm)
	if err != nil {
		return errors.Wrap(err, "Error opening file")
	}
	defer file.Close()

	err = json.NewEncoder(file).Encode(f.stack)
	if err != nil {
		return errors.Wrap(err, "Error encoding permission stack")
	}

	return nil
}
