package filesounddb

import (
	"esefexapi/sounddb"
	"slices"

	"github.com/pkg/errors"
)

func (f *FileDB) SoundExists(uid sounddb.SoundUID) (bool, error) {
	uids, err := f.GetSoundUIDs(uid.ServerID)
	if err != nil {
		return false, errors.Wrap(err, "Error getting sound uids")
	}

	return slices.Contains(uids, uid), nil
}
