package fileuserdb

import (
	"esefexapi/userdb"
	"esefexapi/util"
	"log"
	"slices"

	"github.com/pkg/errors"
)

// GetUser implements userdb.UserDB.
func (f *FileUserDB) GetUser(id string) (*userdb.User, error) {
	for _, user := range f.Users {
		if user.ID == id {
			return &user, nil
		}
	}
	return nil, userdb.ErrUserNotFound
}

// AddUser implements userdb.UserDB.
func (f *FileUserDB) SetUser(user userdb.User) error {
	f.Users[user.ID] = user

	go f.Save()

	return nil
}

// DeleteUser implements userdb.UserDB.
func (f *FileUserDB) DeleteUser(id string) error {
	delete(f.Users, id)

	go f.Save()

	return nil
}

// GetAllUsers implements userdb.UserDB.
func (f *FileUserDB) GetAllUsers() ([]*userdb.User, error) {
	users := make([]*userdb.User, 0, len(f.Users))
	for _, user := range f.Users {
		users = append(users, &user)
	}
	return users, nil
}

// GetUserByToken implements userdb.UserDB.
func (f *FileUserDB) GetUserByToken(token userdb.Token) (*userdb.User, error) {
	for _, user := range f.Users {
		if slices.Contains(user.Tokens, token) {
			return &user, nil
		}
	}
	return nil, userdb.ErrUserNotFound
}

func (f *FileUserDB) NewToken(userID string) (userdb.Token, error) {
	token := util.RandomString(util.TokenCharset, 32)

	user, err := f.getOrCreateUser(userID)
	if err != nil {
		return "", errors.Wrap(err, "Error getting or creating user")
	}

	user.Tokens = append(user.Tokens, userdb.Token(token))
	f.SetUser(*user)

	log.Printf("New token for user %s: %s\n", userID, token)
	log.Printf("%v", f)

	go f.Save()

	return userdb.Token(token), nil
}

func (f *FileUserDB) getOrCreateUser(userID string) (*userdb.User, error) {
	user, err := f.GetUser(userID)
	if err == userdb.ErrUserNotFound {
		f.SetUser(userdb.User{
			ID:     userID,
			Tokens: []userdb.Token{},
		})
		user, err = f.GetUser(userID)
	}
	if err != nil {
		return nil, errors.Wrap(err, "Error getting user")
	}
	return user, nil
}
