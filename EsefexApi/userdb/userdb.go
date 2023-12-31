package userdb

import "github.com/pkg/errors"

var ErrUserNotFound = errors.New("User not found")

type IUserDB interface {
	GetUser(userID string) (*User, error)
	GetUserByToken(token Token) (*User, error)
	GetAllUsers() ([]*User, error)
	SetUser(user User) error
	DeleteUser(userID string) error
	NewToken(userID string) (Token, error)
}
