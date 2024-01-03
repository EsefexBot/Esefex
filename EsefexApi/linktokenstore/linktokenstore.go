package linktokenstore

import (
	"fmt"
	"time"
)

var ErrTokenNotFound = fmt.Errorf("Token not found")
var ErrTokenExpired = fmt.Errorf("Token expired")

type ILinkTokenStore interface {
	// Get a token for a user
	GetToken(userID string) (LinkToken, error)
	// Get a user for a token
	GetUser(tokenStr string) (string, error)
	// Set a token for a user
	SetToken(userID string, token LinkToken) error
	// Delete a token for a user
	DeleteToken(userID string) error
	// Create a new token for a user (the token must be unique)
	CreateToken(userID string) (LinkToken, error)
	// Validate a token
	ValidateToken(tokenStr string) (bool, error)
}

type LinkToken struct {
	Token  string
	Expiry time.Time
}
