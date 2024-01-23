package userdb

import (
	"esefexapi/opt"
	// "github.com/pkg/errors"
)

// var ErrUserNotFound = errors.New("User not found")

type IUserDB interface {
	GetUser(userID string) (opt.Option[*User], error)
	SetUser(user User) error
	DeleteUser(userID string) error
	GetAllUsers() ([]*User, error)
	GetUserByToken(token Token) (opt.Option[*User], error)
	NewToken(userID string) (Token, error)
}
