package filedb

import (
	"encoding/json"
	"esefexapi/sounddb"
	"fmt"
	"io"
	"log"
	"os"
)

// GetSoundMeta implements sounddb.SoundDB.
func (f *FileDB) GetSoundMeta(uid sounddb.SoundUID) (sounddb.SoundMeta, error) {
	path := fmt.Sprintf("sounds/%s/%s_meta.json", uid.ServerID, uid.SoundID)
	metaFile, err := os.Open(path)

	if err != nil {
		log.Fatal(err)
	}

	var sound sounddb.SoundMeta

	byteValue, _ := io.ReadAll(metaFile)
	json.Unmarshal(byteValue, &sound)
	metaFile.Close()

	return sound, nil
}
