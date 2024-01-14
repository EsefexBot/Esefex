package filepermisssiondb

import (
	"encoding/json"
	"esefexapi/permissiondb"
	"esefexapi/permissions"
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
// TODO: Make this work
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
