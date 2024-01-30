package filesounddb

import (
	"encoding/binary"
	"encoding/json"
	"esefexapi/sounddb"
	"esefexapi/types"
	"fmt"
	"log"
	"os"

	"github.com/pkg/errors"
)

// AddSound implements sounddb.SoundDB.
func (f *FileDB) AddSound(guildID types.GuildID, name types.SoundName, icon sounddb.Icon, pcm []int16) (sounddb.SoundUID, error) {
	// check if sound already exists
	suid := sounddb.SoundUID{
		GuildID:   guildID,
		SoundName: name,
	}

	exists, err := f.SoundExists(suid)
	if err != nil {
		return sounddb.SoundUID{}, errors.Wrap(err, "Error checking if sound exists")
	}

	if exists {
		return sounddb.SoundUID{}, errors.Wrap(fmt.Errorf("Sound already exists, pick a different name or delete the existing sound"), "Sound already exists")
	}

	sound := sounddb.SoundMeta{
		SoundID: name.GetSoundID(),
		GuildID: guildID,
		Name:    name,
		Icon:    icon,
	}

	// Make sure the db folder exists
	path := fmt.Sprintf("%s/%s", f.location, guildID)
	err = os.MkdirAll(path, os.ModePerm)
	if err != nil {
		log.Printf("Error creating guild folder: %+v", err)
		return sounddb.SoundUID{}, errors.Wrap(err, "Error creating guild folder")
	}

	// write meta file
	path = fmt.Sprintf("%s/%s/%s_meta.json", f.location, guildID, sound.SoundID)
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

	_, err = metaFile.Write(metaJson)
	if err != nil {
		log.Printf("Error writing meta file: %+v", err)
		return sounddb.SoundUID{}, errors.Wrap(err, "Error writing meta file")
	}
	metaFile.Close()

	// write sound file

	path = fmt.Sprintf("%s/%s/%s_sound.s16le", f.location, guildID, sound.SoundID)

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
