package memorylinktokenstore

import (
	"esefexapi/linktokenstore"
	"esefexapi/types"
	"esefexapi/util"
	"time"

	"github.com/pkg/errors"
)

var _ linktokenstore.ILinkTokenStore = &MemoryLinkTokenStore{}

type MemoryLinkTokenStore struct {
	linkTokens  map[types.UserID]linktokenstore.LinkToken
	expireAfter time.Duration
}

// CreateToken implements linktokenstore.ILinkTokenStore.
func (m *MemoryLinkTokenStore) CreateToken(userID types.UserID) (linktokenstore.LinkToken, error) {
	for {
		token := util.RandomString(util.TokenCharset, 32)
		// check if token exists
		// if not, return token
		// if yes, try again
		_, err := m.GetUser(token)
		if err != nil {
			token := linktokenstore.LinkToken{
				Token:  token,
				Expiry: time.Now().Add(time.Hour * 24),
			}

			err = m.SetToken(userID, token)
			if err != nil {
				return linktokenstore.LinkToken{}, errors.Wrap(err, "Error setting token")
			}
			return token, nil
		}
	}
}

func NewMemoryLinkTokenStore(expireAfter time.Duration) *MemoryLinkTokenStore {
	return &MemoryLinkTokenStore{
		linkTokens:  map[types.UserID]linktokenstore.LinkToken{},
		expireAfter: expireAfter,
	}
}

func (m *MemoryLinkTokenStore) GetToken(userID types.UserID) (linktokenstore.LinkToken, error) {
	return m.linkTokens[userID], nil
}

func (m *MemoryLinkTokenStore) GetUser(token string) (types.UserID, error) {
	for k, v := range m.linkTokens {
		if v.Token == token {
			return k, nil
		}
	}

	return "", linktokenstore.ErrTokenNotFound
}

func (m *MemoryLinkTokenStore) SetToken(userID types.UserID, token linktokenstore.LinkToken) error {
	m.linkTokens[userID] = token
	return nil
}

func (m *MemoryLinkTokenStore) DeleteToken(userID types.UserID) error {
	delete(m.linkTokens, userID)
	return nil
}

func (m *MemoryLinkTokenStore) ValidateToken(tokenStr string) (bool, error) {
	user, err := m.GetUser(tokenStr)
	if err != nil {
		return false, errors.Wrap(err, "Error getting user from token")
	}

	token, err := m.GetToken(user)
	if err != nil {
		return false, errors.Wrap(err, "Error getting token")
	}

	if token.Expiry.Before(time.Now()) {
		err = m.DeleteToken(user)
		if err != nil {
			return false, errors.Wrap(err, "Error deleting token")
		}
		return false, linktokenstore.ErrTokenExpired
	}

	return token.Token == tokenStr, nil
}
