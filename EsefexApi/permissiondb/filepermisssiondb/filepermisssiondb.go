package filepermisssiondb

import (
	"encoding/json"
	"esefexapi/permissiondb"
	"esefexapi/permissions"
	"esefexapi/util"
	"log"
	"os"
	"sync"

	"github.com/pkg/errors"
)

var _ permissiondb.PermissionDB = &FilePermissionDB{}

type FilePermissionDB struct {
	file  *os.File
	rw    *sync.RWMutex
	stack permissions.PermissionStack
}

func NewFilePermissionDB(path string) (*FilePermissionDB, error) {
	log.Printf("Creating FileDB at %s", path)
	file, err := util.EnsureFile(path)
	if err != nil {
		return nil, errors.Wrap(err, "Error ensuring file")
	}

	fpdb := &FilePermissionDB{
		file:  file,
		rw:    &sync.RWMutex{},
		stack: permissions.NewPermissionStack(),
	}
	err = fpdb.load()
	if err != nil {
		return nil, errors.Wrap(err, "Error loading permission stack")
	}

	return fpdb, nil
}

// load loads the permission stack from the file.
func (f *FilePermissionDB) load() error {
	f.rw.Lock()
	defer f.rw.Unlock()

	// reset file
	_, err := f.file.Seek(0, 0)
	if err != nil {
		return errors.Wrap(err, "Error seeking to start of file")
	}

	// read file
	var perms permissions.PermissionStack
	err = json.NewDecoder(f.file).Decode(&perms)
	if err != nil {
		log.Printf("Error decoding file, creating empty permission stack: (%v)", err)
		f.stack = permissions.NewPermissionStack()
	} else {
		f.stack = perms
	}

	return nil
}

func (f *FilePermissionDB) Close() error {
	f.rw.Lock()
	defer f.rw.Unlock()
	log.Println("Closing file permissiondb")

	err := f.save()
	if err != nil {
		return errors.Wrap(err, "Error saving permissiondb")
	}
	return f.file.Close()
}

// save saves the permission stack to the file.
// It assumes that the lock is already held.
func (f FilePermissionDB) save() error {
	// reset file
	_, err := f.file.Seek(0, 0)
	if err != nil {
		return errors.Wrap(err, "Error seeking to start of file")
	}
	err = f.file.Truncate(0)
	if err != nil {
		return errors.Wrap(err, "Error truncating file")
	}

	err = json.NewEncoder(f.file).Encode(f.stack)
	if err != nil {
		return errors.Wrap(err, "Error encoding file")
	}

	return nil
}
