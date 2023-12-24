package filedb

import (
	"esefexapi/sounddb"
	"fmt"
	"log"
	"os"
)

// DeleteSound implements sounddb.SoundDB.
func (f *FileDB) DeleteSound(uid sounddb.SoundUID) error {
	path := fmt.Sprintf("%s/%s/%s_meta.json", f.location, uid.ServerID, uid.SoundID)
	err := os.Remove(path)
	if err != nil {
		log.Fatal(err)
		return err
	}

	path = fmt.Sprintf("%s/%s/%s_sound.mp3", f.location, uid.ServerID, uid.SoundID)
	err = os.Remove(path)
	if err != nil {
		log.Fatal(err)
		return err
	}

	return nil
}
