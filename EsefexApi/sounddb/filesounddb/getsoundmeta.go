package filesounddb

import (
	"encoding/json"
	"esefexapi/sounddb"
	"fmt"
	"io"
	"os"
)

// GetSoundMeta implements sounddb.SoundDB.
func (f *FileDB) GetSoundMeta(uid sounddb.SoundUID) (sounddb.SoundMeta, error) {
	path := fmt.Sprintf("%s/%s/%s_meta.json", f.location, uid.GuildID, uid.SoundName.GetSoundID())
	metaFile, err := os.Open(path)

	if err != nil {
		return sounddb.SoundMeta{}, err
	}

	var sound sounddb.SoundMeta

	byteValue, _ := io.ReadAll(metaFile)
	err = json.Unmarshal(byteValue, &sound)
	if err != nil {
		return sounddb.SoundMeta{}, err
	}
	metaFile.Close()

	return sound, nil
}
