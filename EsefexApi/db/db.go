package db

import (
	"esefexapi/bot/commands/cmdhashstore"
	"esefexapi/linktokenstore"
	"esefexapi/permissiondb"
	"esefexapi/sounddb"
	"esefexapi/userdb"
)

type Databases struct {
	SoundDB        sounddb.ISoundDB
	UserDB         userdb.IUserDB
	LinkTokenStore linktokenstore.ILinkTokenStore
	PermissionDB   permissiondb.IPermissionDB
	CmdHashStore   cmdhashstore.ICommandHashStore
}

// func CreateDatabases(cfg *config.Config, ds *discordgo.Session) (*Databases, error) {
// 	dbs := &Databases{}

// 	sdb, err := filesounddb.NewFileDB(cfg.Database.SounddbLocation)
// 	if err != nil {
// 		return nil, err
// 	}
// 	sdbc, err := dbcache.NewSoundDBCache(sdb)
// 	if err != nil {
// 		return nil, err
// 	}
// 	dbs.SoundDB = sdbc

// 	udb, err := fileuserdb.NewFileUserDB(cfg.Database.UserdbLocation)
// 	if err != nil {
// 		return nil, err
// 	}
// 	dbs.UserDB = udb

// 	pdb, err := filepermisssiondb.NewFilePermissionDB(cfg.Database.Permissiondblocation, ds)
// 	if err != nil {
// 		return nil, err
// 	}
// 	dbs.PermissionDB = pdb

// 	verT := time.Duration(cfg.VerificationExpiry * float32(time.Minute))
// 	ldb := memorylinktokenstore.NewMemoryLinkTokenStore(verT)
// 	dbs.LinkTokenStore = ldb

// 	hdb := cmdhashstore.NewFileCmdHashStore(cfg.Database.CmdHashStoreLocation)
// 	dbs.CmdHashStore = hdb

// 	return dbs, nil
// }
