package db

import (
	"esefexapi/linktokenstore"
	"esefexapi/sounddb"
	"esefexapi/userdb"
)

type Databases struct {
	SoundDB        sounddb.ISoundDB
	UserDB         userdb.IUserDB
	LinkTokenStore linktokenstore.ILinkTokenStore
}
