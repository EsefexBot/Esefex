package filedb

import (
	"encoding/binary"
	"encoding/json"
	"esefexapi/db"
	"fmt"
	"log"
	"os"
)

func (f *FileDB) AddSound(serverID string, name string, iconUrl string, pcm []int16) (db.SoundUID, error) {
	sid, err := f.generateSoundID(serverID)
	if err != nil {
		return db.SoundUID{}, err
	}

	sound := db.SoundMeta{
		SoundID:  sid,
		ServerID: serverID,
		Name:     name,
		Icon:     iconUrl,
	}

	// Make sure the sounds folder exists
	path := fmt.Sprintf("sounds/%s", serverID)
	os.MkdirAll(path, os.ModePerm)

	// write meta file
	path = fmt.Sprintf("sounds/%s/%s_meta.json", serverID, sound.SoundID)
	metaFile, err := os.Create(path)
	if err != nil {
		log.Fatal(err)
		return db.SoundUID{}, err
	}

	metaJson, err := json.Marshal(sound)
	if err != nil {
		log.Fatal(err)
		return db.SoundUID{}, err
	}

	metaFile.Write(metaJson)
	metaFile.Close()

	// write sound file

	path = fmt.Sprintf("sounds/%s/%s_sound.s16le", serverID, sound.SoundID)

	soundFile, err := os.Create(path)
	if err != nil {
		log.Fatal(err)
		return db.SoundUID{}, err
	}

	err = binary.Write(soundFile, binary.LittleEndian, pcm)
	if err != nil {
		log.Fatal(err)
		return db.SoundUID{}, err
	}

	return sound.GetUID(), nil
}
