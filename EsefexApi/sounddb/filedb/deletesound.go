package filedb

import (
	"esefexapi/sounddb"
	"fmt"
	"log"
	"os"
)

// DeleteSound implements sounddb.SoundDB.
func (f *FileDB) DeleteSound(uid sounddb.SoundUID) error {
	path := fmt.Sprintf("sounds/%s/%s_meta.json", uid.ServerID, uid.SoundID)
	err := os.Remove(path)
	if err != nil {
		log.Fatal(err)
		return err
	}

	path = fmt.Sprintf("sounds/%s/%s_sound.mp3", uid.ServerID, uid.SoundID)
	err = os.Remove(path)
	if err != nil {
		log.Fatal(err)
		return err
	}

	return nil
}
