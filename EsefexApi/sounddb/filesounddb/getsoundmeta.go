package filesounddb

import (
	"encoding/json"
	"esefexapi/sounddb"
	"fmt"
	"io"
	"os"
)

// GetSoundMeta implements sounddb.SoundDB.
func (f *FileDB) GetSoundMeta(uid sounddb.SoundURI) (sounddb.SoundMeta, error) {
	path := fmt.Sprintf("%s/%s/%s_meta.json", f.location, uid.GuildID, uid.SoundID)
	metaFile, err := os.Open(path)

	if err != nil {
		return sounddb.SoundMeta{}, err
	}

	var sound sounddb.SoundMeta

	byteValue, _ := io.ReadAll(metaFile)
	json.Unmarshal(byteValue, &sound)
	metaFile.Close()

	return sound, nil
}
