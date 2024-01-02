package fileuserdb

import (
	"esefexapi/opt"
	"esefexapi/userdb"
	"esefexapi/util"
	"log"
	"slices"

	"github.com/pkg/errors"
)

// GetUser implements userdb.UserDB.
func (f *FileUserDB) GetUser(id string) (opt.Option[*userdb.User], error) {
	for _, user := range f.Users {
		if user.ID == id {
			return opt.Some(&user), nil
		}
	}
	return opt.None[*userdb.User](), nil
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
func (f *FileUserDB) GetUserByToken(token userdb.Token) (opt.Option[*userdb.User], error) {
	for _, user := range f.Users {
		if slices.Contains(user.Tokens, token) {
			return opt.Some(&user), nil
		}
	}

	return opt.None[*userdb.User](), nil
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
	Ouser, err := f.GetUser(userID)
	if Ouser.IsNone() {
		f.SetUser(userdb.User{
			ID:     userID,
			Tokens: []userdb.Token{},
		})
		Ouser, err = f.GetUser(userID)
	}
	if err != nil {
		return nil, errors.Wrap(err, "Error getting user")
	}
	return Ouser.Unwrap(), nil
}
