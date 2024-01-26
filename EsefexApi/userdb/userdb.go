package userdb

import (
	"esefexapi/opt"
	"esefexapi/types"
	// "github.com/pkg/errors"
)

// var ErrUserNotFound = fmt.Errorf("User not found")

type IUserDB interface {
	GetUser(userID types.UserID) (opt.Option[*User], error)
	SetUser(user User) error
	DeleteUser(userID types.UserID) error
	GetAllUsers() ([]*User, error)
	GetUserByToken(token Token) (opt.Option[*User], error)
	NewToken(userID types.UserID) (Token, error)
}
