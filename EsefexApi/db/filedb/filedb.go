package filedb

import (
	"esefexapi/db"
)

var _ db.SoundDB = &FileDB{}

// FileDB implements SoundDB
type FileDB struct{}

// func GetSoundMetas(serverId string) []db.SoundMeta {
// 	ids := getSoundIDs(serverId)
// 	sounds := make([]db.SoundMeta, 0)

// 	for _, id := range ids {
// 		sounds = append(sounds, GetSoundMeta(db.SuidFromStrings(serverId, id)))
// 	}

// 	return sounds
// }

// func getAllSoundUids() []db.SoundUid {
// 	servers := GetAllServerIds()
// 	suids := make([]db.SoundUid, 0)

// 	for _, server := range servers {
// 		ids := getSoundIDs(server)
// 		for _, id := range ids {
// 			suids = append(suids, db.SuidFromStrings(server, id))
// 		}
// 	}

// 	return suids
// }
