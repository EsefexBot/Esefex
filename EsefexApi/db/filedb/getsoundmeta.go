package filedb

import (
	"encoding/json"
	"esefexapi/db"
	"fmt"
	"io"
	"log"
	"os"
)

func (f *FileDB) GetSoundMeta(uid db.SoundUID) (db.SoundMeta, error) {
	path := fmt.Sprintf("sounds/%s/%s_meta.json", uid.ServerID, uid.SoundID)
	metaFile, err := os.Open(path)

	if err != nil {
		log.Fatal(err)
	}

	var sound db.SoundMeta

	byteValue, _ := io.ReadAll(metaFile)
	json.Unmarshal(byteValue, &sound)
	metaFile.Close()

	return sound, nil
}
