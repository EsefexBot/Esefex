package filedb

import (
	"encoding/binary"
	"encoding/json"
	"esefexapi/sounddb"
	"fmt"
	"log"
	"os"
)

// AddSound implements sounddb.SoundDB.
func (f *FileDB) AddSound(serverID string, name string, iconUrl string, pcm []int16) (sounddb.SoundUID, error) {
	sid, err := f.generateSoundID(serverID)
	if err != nil {
		return sounddb.SoundUID{}, err
	}

	sound := sounddb.SoundMeta{
		SoundID:  sid,
		ServerID: serverID,
		Name:     name,
		Icon:     iconUrl,
	}

	// Make sure the db folder exists
	path := fmt.Sprintf("%s/%s", f.location, serverID)
	os.MkdirAll(path, os.ModePerm)

	// write meta file
	path = fmt.Sprintf("%s/%s/%s_meta.json", f.location, serverID, sound.SoundID)
	metaFile, err := os.Create(path)
	if err != nil {
		log.Fatal(err)
		return sounddb.SoundUID{}, err
	}

	metaJson, err := json.Marshal(sound)
	if err != nil {
		log.Fatal(err)
		return sounddb.SoundUID{}, err
	}

	metaFile.Write(metaJson)
	metaFile.Close()

	// write sound file

	path = fmt.Sprintf("%s/%s/%s_sound.s16le", f.location, serverID, sound.SoundID)

	soundFile, err := os.Create(path)
	if err != nil {
		log.Fatal(err)
		return sounddb.SoundUID{}, err
	}

	err = binary.Write(soundFile, binary.LittleEndian, pcm)
	if err != nil {
		log.Fatal(err)
		return sounddb.SoundUID{}, err
	}

	return sound.GetUID(), nil
}
