package filesounddb

import (
	"encoding/binary"
	"esefexapi/sounddb"
	"fmt"
	"io"
	"os"

	"github.com/pkg/errors"
)

// GetSoundPcm implements sounddb.SoundDB.
func (f *FileDB) GetSoundPcm(uid sounddb.SoundURI) (*[]int16, error) {
	path := fmt.Sprintf("%s/%s/%s_sound.s16le", f.location, uid.GuildID, uid.SoundID)
	sf, err := os.Open(path)
	if err != nil {
		return nil, errors.Wrap(err, "Error opening sound file")
	}

	sfs, err := sf.Stat()
	if err != nil {
		return nil, errors.Wrap(err, "Error getting sound file stats")
	}

	pcm := make([]int16, sfs.Size()/2)

	err = binary.Read(sf, binary.LittleEndian, &pcm)
	if err != nil && err != io.EOF {
		return nil, errors.Wrap(err, "Error reading sound file")
	}

	return &pcm, nil
}
