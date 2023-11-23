package filedb

import (
	"encoding/json"
	"fmt"
	"io"
	"math/rand"
	"os"
	"slices"
	"strconv"
	"strings"
)

type SoundMeta struct {
	Id       string `json:"id"`
	ServerId string `json:"serverId"`
	Name     string `json:"name"`
	Icon     string `json:"icon"`
}

func GetSound(serverId string, soundId string) SoundMeta {
	path := fmt.Sprintf("sounds/%s/%s_meta.json", serverId, soundId)
	metaFile, err := os.Open(path)

	if err != nil {
		fmt.Println(err)
	}

	var sound SoundMeta

	byteValue, _ := io.ReadAll(metaFile)
	json.Unmarshal(byteValue, &sound)
	metaFile.Close()

	return sound
}

// this isnt pretty but it works for now (and probably forever)
func GenerateSoundID(serverId string) string {
	// generate random number with 16 digits
	min := 100000000
	max := 999999999

	for true {
		id := strconv.FormatInt(int64(rand.Intn(max-min)+min), 10)

		if !SoundExists(serverId, id) {
			return id
		}
	}

	panic("Could not generate sound id, all ids are taken!!!! Thats like almost a billion sounds! How long have you been running this server?!?!?!")
}

func SoundExists(serverId string, soundId string) bool {
	return slices.Contains(GetSoundIDs(serverId), soundId)
}

func GetSounds(serverId string) []SoundMeta {
	ids := GetSoundIDs(serverId)
	sounds := make([]SoundMeta, 0)

	for _, id := range ids {
		sounds = append(sounds, GetSound(serverId, id))
	}

	return sounds
}

func GetSoundIDs(serverId string) []string {
	path := fmt.Sprintf("sounds/%s", serverId)
	files, err := os.ReadDir(path)

	if err != nil {
		fmt.Println(err)
	}

	ids := make([]string, 0)

	for _, file := range files {
		name := file.Name()
		nameSplits := strings.Split(name, "_")

		if len(nameSplits) == 2 && nameSplits[1] == "meta.json" {
			ids = append(ids, nameSplits[0])
		}
	}

	return ids
}

func AddSound(serverId string, name string, image string, file []byte) string {
	sound := SoundMeta{
		Id:       GenerateSoundID(serverId),
		ServerId: serverId,
		Name:     name,
		Icon:     image,
	}

	// Make sure the sounds folder exists
	path := fmt.Sprintf("sounds/%s", serverId)
	os.MkdirAll(path, os.ModePerm)

	// write meta file
	path = fmt.Sprintf("sounds/%s/%s_meta.json", serverId, sound.Id)
	metaFile, err := os.Create(path)

	if err != nil {
		fmt.Println(err)
	}

	metaJson, err := json.Marshal(sound)

	if err != nil {
		fmt.Println(err)
	}

	metaFile.Write(metaJson)
	metaFile.Close()

	// write sound file

	path = fmt.Sprintf("sounds/%s/%s_sound.mp3", serverId, sound.Id)

	soundFile, err := os.Create(path)

	if err != nil {
		fmt.Println(err)
	}

	soundFile.Write(file)
	soundFile.Close()

	return sound.Id
}

func DeleteSound(serverId string, sound_id string) {
	path := fmt.Sprintf("sounds/%s/%s_meta.json", serverId, sound_id)
	err := os.Remove(path)
	if err != nil {
		fmt.Println(err)
	}

	path = fmt.Sprintf("sounds/%s/%s_sound.mp3", serverId, sound_id)
	err = os.Remove(path)
	if err != nil {
		fmt.Println(err)
	}
}
