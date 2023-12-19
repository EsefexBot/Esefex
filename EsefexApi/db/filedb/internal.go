package filedb

import (
	"esefexapi/db"
	"math/rand"
	"slices"
	"strconv"
)

func (f *FileDB) generateSoundID(serverId string) (string, error) {
	// generate random number with 16 digits
	min := 100000000
	max := 999999999

	for {
		id := strconv.FormatInt(int64(rand.Intn(max-min)+min), 10)

		exists, err := f.soundExists(db.SuidFromStrings(serverId, id))
		if err != nil {
			return "", err
		}

		if !exists {
			return id, nil
		}
	}
}

func (f *FileDB) soundExists(uid db.SoundUID) (bool, error) {
	uids, err := f.GetSoundUIDs(uid.ServerID)
	if err != nil {
		return false, err
	}

	return slices.Contains(uids, uid), nil
}
