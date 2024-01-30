package filesounddb

import (
	"esefexapi/sounddb"
	"esefexapi/util"
	"fmt"
)

func (f *FileDB) SoundExists(uid sounddb.SoundUID) (bool, error) {
	path := fmt.Sprintf("%s/%s/%s_meta.json", f.location, uid.GuildID, uid.SoundName.GetSoundID())
	exists, err := util.PathExists(path)
	if err != nil {
		return false, err
	}

	return exists, nil
}
