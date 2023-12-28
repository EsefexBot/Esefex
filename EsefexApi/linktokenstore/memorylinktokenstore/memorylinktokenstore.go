package memorylinktokenstore

import (
	"esefexapi/linktokenstore"
	"esefexapi/util"
	"time"
)

var _ linktokenstore.ILinkTokenStore = &MemoryLinkTokenStore{}

type MemoryLinkTokenStore struct {
	linkTokens  map[string]linktokenstore.LinkToken
	expireAfter time.Duration
}

// CreateToken implements linktokenstore.ILinkTokenStore.
func (m *MemoryLinkTokenStore) CreateToken(userID string) (linktokenstore.LinkToken, error) {
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

			m.SetToken(userID, token)
			return token, nil
		}
	}
}

func NewMemoryLinkTokenStore(expireAfter time.Duration) *MemoryLinkTokenStore {
	return &MemoryLinkTokenStore{
		linkTokens:  map[string]linktokenstore.LinkToken{},
		expireAfter: expireAfter,
	}
}

func (m *MemoryLinkTokenStore) GetToken(userID string) (linktokenstore.LinkToken, error) {
	return m.linkTokens[userID], nil
}

func (m *MemoryLinkTokenStore) GetUser(token string) (string, error) {
	for k, v := range m.linkTokens {
		if v.Token == token {
			return k, nil
		}
	}

	return "", linktokenstore.ErrTokenNotFound
}

func (m *MemoryLinkTokenStore) SetToken(userID string, token linktokenstore.LinkToken) error {
	m.linkTokens[userID] = token
	return nil
}

func (m *MemoryLinkTokenStore) DeleteToken(userID string) error {
	delete(m.linkTokens, userID)
	return nil
}

func (m *MemoryLinkTokenStore) ValidateToken(tokenStr string) (bool, error) {
	user, err := m.GetUser(tokenStr)
	if err != nil {
		return false, err
	}

	token, err := m.GetToken(user)
	if err != nil {
		return false, err
	}

	if token.Expiry.Before(time.Now()) {
		m.DeleteToken(user)
		return false, linktokenstore.ErrTokenExpired
	}

	return token.Token == tokenStr, nil
}
