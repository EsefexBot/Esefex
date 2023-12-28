package filesounddb

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
		log.Println(err)
		return err
	}

	path = fmt.Sprintf("%s/%s/%s_sound.s16le", f.location, uid.ServerID, uid.SoundID)
	err = os.Remove(path)
	if err != nil {
		log.Println(err)
		return err
	}

	return nil
}
