package fileuserdb

import (
	"esefexapi/config"
	"esefexapi/userdb"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFileUserDB(t *testing.T) {
	config.InjectConfig(&config.Config{
		Database: config.Database{
			UserdbLocation: "./testdata/users.json",
		},
	})

	udb, err := NewFileUserDB()
	assert.Nil(t, err)

	defer func() {
		_ = udb.Close()
		_ = os.Remove("testdata/users.json")
		_ = os.Remove("testdata")
	}()

	assert.Equal(t, 0, len(udb.Users))

	user1 := userdb.User{
		ID: "1",
		Tokens: []userdb.Token{
			"1",
			"2",
		},
	}

	user2 := userdb.User{
		ID: "2",
		Tokens: []userdb.Token{
			"3",
			"4",
		},
	}

	err = udb.SetUser(user1)
	assert.Nil(t, err)

	err = udb.SetUser(user2)
	assert.Nil(t, err)

	assert.Equal(t, 2, len(udb.Users))

	user1.Tokens = []userdb.Token{
		"1",
		"2",
		"3",
	}

	err = udb.SetUser(user1)
	assert.Nil(t, err)

	assert.Equal(t, 2, len(udb.Users))

	user, err := udb.GetUser("1")
	assert.Nil(t, err)

	assert.Equal(t, &user1, user.Unwrap())

	users, err := udb.GetAllUsers()
	assert.Nil(t, err)

	assert.Equal(t, 2, len(users))

	user, err = udb.GetUserByToken("1")
	assert.Nil(t, err)
	assert.Equal(t, &user1, user.Unwrap())

	err = udb.DeleteUser("1")
	assert.Nil(t, err)

	assert.Equal(t, 1, len(udb.Users))

	Ouser, err := udb.GetUser("1")
	assert.Nil(t, err)
	assert.True(t, Ouser.IsNone())

	err = udb.DeleteUser("2")
	assert.Nil(t, err)

	assert.Equal(t, 0, len(udb.Users))
}
