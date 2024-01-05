package fileuserdb

import (
	"encoding/json"
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

func NewFileUserDB(filePath string) (*FileUserDB, error) {
	// get file handle
	os.MkdirAll(path.Dir(filePath), os.ModePerm)

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

	f.Save()
	return f.file.Close()
}

func (f *FileUserDB) Save() error {
	f.fileLock.Lock()
	defer f.fileLock.Unlock()

	// reset file
	f.file.Seek(0, 0)
	f.file.Truncate(0)

	usrArr := make([]userdb.User, 0, len(f.Users))
	for _, user := range f.Users {
		usrArr = append(usrArr, user)
	}

	// write file
	return json.NewEncoder(f.file).Encode(usrArr)
}
