package filesounddb

import (
	"encoding/binary"
	"encoding/json"
	"esefexapi/sounddb"
	"fmt"
	"log"
	"os"

	"github.com/pkg/errors"
)

// AddSound implements sounddb.SoundDB.
func (f *FileDB) AddSound(serverID string, name string, icon sounddb.Icon, pcm []int16) (sounddb.SoundUID, error) {
	sid, err := f.generateSoundID(serverID)
	if err != nil {
		return sounddb.SoundUID{}, errors.Wrap(err, "Error generating sound ID")
	}

	sound := sounddb.SoundMeta{
		SoundID:  sid,
		ServerID: serverID,
		Name:     name,
		Icon:     icon,
	}

	// Make sure the db folder exists
	path := fmt.Sprintf("%s/%s", f.location, serverID)
	os.MkdirAll(path, os.ModePerm)

	// write meta file
	path = fmt.Sprintf("%s/%s/%s_meta.json", f.location, serverID, sound.SoundID)
	metaFile, err := os.Create(path)
	if err != nil {
		log.Printf("Error creating meta file: %+v", err)
		return sounddb.SoundUID{}, errors.Wrap(err, "Error creating meta file")
	}

	metaJson, err := json.Marshal(sound)
	if err != nil {
		log.Printf("Error marshalling meta: %+v", err)
		return sounddb.SoundUID{}, errors.Wrap(err, "Error marshalling meta")
	}

	metaFile.Write(metaJson)
	metaFile.Close()

	// write sound file

	path = fmt.Sprintf("%s/%s/%s_sound.s16le", f.location, serverID, sound.SoundID)

	soundFile, err := os.Create(path)
	if err != nil {
		log.Printf("Error creating sound file: %+v", err)
		return sounddb.SoundUID{}, errors.Wrap(err, "Error creating sound file")
	}

	err = binary.Write(soundFile, binary.LittleEndian, pcm)
	if err != nil {
		log.Printf("Error writing sound file: %+v", err)
		return sounddb.SoundUID{}, errors.Wrap(err, "Error writing sound file")
	}

	return sound.GetUID(), nil
}
