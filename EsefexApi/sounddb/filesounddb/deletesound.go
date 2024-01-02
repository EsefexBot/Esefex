package filesounddb

import (
	"esefexapi/sounddb"
	"fmt"
	"log"
	"os"

	"github.com/pkg/errors"
)

// DeleteSound implements sounddb.SoundDB.
func (f *FileDB) DeleteSound(uid sounddb.SoundUID) error {
	path := fmt.Sprintf("%s/%s/%s_meta.json", f.location, uid.ServerID, uid.SoundID)
	err := os.Remove(path)
	if err != nil {
		log.Printf("Error removing meta file: %+v", err)
		return errors.Wrap(err, "Error removing meta file")
	}

	path = fmt.Sprintf("%s/%s/%s_sound.s16le", f.location, uid.ServerID, uid.SoundID)
	err = os.Remove(path)
	if err != nil {
		log.Printf("Error removing sound file: %+v", err)
		return errors.Wrap(err, "Error removing sound file")
	}

	return nil
}
