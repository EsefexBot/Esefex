package linktokenstore

import (
	"esefexapi/types"
	"fmt"
	"time"
)

var ErrTokenNotFound = fmt.Errorf("Token not found")
var ErrTokenExpired = fmt.Errorf("Token expired")

type ILinkTokenStore interface {
	// Get a token for a user
	GetToken(userID types.UserID) (LinkToken, error)
	// Get a user for a token
	GetUser(tokenStr string) (types.UserID, error)
	// Set a token for a user
	SetToken(userID types.UserID, token LinkToken) error
	// Delete a token for a user
	DeleteToken(userID types.UserID) error
	// Create a new token for a user (the token must be unique)
	CreateToken(userID types.UserID) (LinkToken, error)
	// Validate a token
	ValidateToken(tokenStr string) (bool, error)
}

type LinkToken struct {
	Token  string
	Expiry time.Time
}
