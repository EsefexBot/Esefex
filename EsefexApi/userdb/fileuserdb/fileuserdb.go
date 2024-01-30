package fileuserdb

import (
	"encoding/json"
	"esefexapi/config"
	"esefexapi/types"
	"esefexapi/userdb"
	"log"
	"os"
	"path"
	"sync"

	"github.com/pkg/errors"
)

var _ userdb.IUserDB = &FileUserDB{}

type FileUserDB struct {
	Users    map[types.UserID]userdb.User
	file     *os.File
	fileLock sync.Mutex
}

func NewFileUserDB() (*FileUserDB, error) {
	filePath := config.Get().Database.UserdbLocation

	// get file handle
	err := os.MkdirAll(path.Dir(filePath), os.ModePerm)
	if err != nil {
		return nil, errors.Wrap(err, "Error creating directory")
	}

	file, err := os.OpenFile(filePath, os.O_RDWR|os.O_CREATE, os.ModePerm)
	if err != nil {
		return nil, errors.Wrap(err, "Error opening file")
	}

	stat, err := file.Stat()
	if err != nil {
		return nil, errors.Wrap(err, "Error getting file stats")
	}

	if stat.Size() == 0 {
		log.Println("Users file is empty, writing empty array")
		_, err = file.WriteString("[]")
		if err != nil {
			return nil, errors.Wrap(err, "Error writing to file")
		}
	}

	// read file
	var users []userdb.User
	_ = json.NewDecoder(file).Decode(&users)

	// create map
	userMap := make(map[types.UserID]userdb.User)
	for _, user := range users {
		userMap[user.ID] = user
	}

	return &FileUserDB{
		Users: userMap,
		file:  file,
	}, nil
}

func (f *FileUserDB) Close() error {
	log.Println("Closing userdb")

	err := f.save()
	if err != nil {
		return errors.Wrap(err, "Error saving userdb")
	}
	return f.file.Close()
}

func (f FileUserDB) save() error {
	// reset file
	_, err := f.file.Seek(0, 0)
	if err != nil {
		return errors.Wrap(err, "Error seeking to start of file")
	}
	err = f.file.Truncate(0)
	if err != nil {
		return errors.Wrap(err, "Error truncating file")
	}

	usrArr := make([]userdb.User, 0, len(f.Users))
	for _, user := range f.Users {
		usrArr = append(usrArr, user)
	}

	// write file
	return json.NewEncoder(f.file).Encode(usrArr)
}
