package filesounddb

import (
	"esefexapi/sounddb"
	"slices"
)

func (f *FileDB) SoundExists(uid sounddb.SoundUID) (bool, error) {
	uids, err := f.GetSoundUIDs(uid.ServerID)
	if err != nil {
		return false, err
	}

	return slices.Contains(uids, uid), nil
}
