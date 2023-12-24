package filedb

import (
	"encoding/binary"
	"esefexapi/sounddb"
	"fmt"
	"io"
	"os"
)

// GetSoundPcm implements sounddb.SoundDB.
func (f *FileDB) GetSoundPcm(uid sounddb.SoundUID) ([]int16, error) {
	path := fmt.Sprintf("%s/%s/%s_sound.s16le", f.location, uid.ServerID, uid.SoundID)
	sf, err := os.Open(path)
	if err != nil {
		return nil, err
	}

	sfs, err := sf.Stat()
	if err != nil {
		return nil, err
	}

	pcm := make([]int16, sfs.Size()/2)

	err = binary.Read(sf, binary.LittleEndian, &pcm)
	if err != nil && err != io.EOF {
		return nil, err
	}

	return pcm, nil
}
