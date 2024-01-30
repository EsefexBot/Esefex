package filesounddb

import (
	"esefexapi/sounddb"
	"fmt"
	"os"

	"github.com/pkg/errors"
)

// DeleteSound implements sounddb.SoundDB.
func (f *FileDB) DeleteSound(uid sounddb.SoundUID) error {
	path := fmt.Sprintf("%s/%s/%s_meta.json", f.location, uid.GuildID, uid.SoundName.GetSoundID())
	err := os.Remove(path)
	if err != nil {
		return errors.Wrap(err, "Error removing meta file")
	}

	path = fmt.Sprintf("%s/%s/%s_sound.s16le", f.location, uid.GuildID, uid.SoundName.GetSoundID())
	err = os.Remove(path)
	if err != nil {
		return errors.Wrap(err, "Error removing sound file")
	}

	return nil
}
