package filedb

import (
	"encoding/json"
	"esefexapi/util"
	"fmt"
	"io"
	"log"
	"math/rand"
	"os"
	"slices"
	"strconv"
	"strings"
)

type SoundUid struct {
	ServerId string
	SoundId  string
}

type SoundMeta struct {
	SoundId  string `json:"id"`
	ServerId string `json:"serverId"`
	Name     string `json:"name"`
	Icon     string `json:"icon"`
}

func SuidFromStrings(serverId string, soundId string) SoundUid {
	return SoundUid{
		ServerId: serverId,
		SoundId:  soundId,
	}
}

func GetSoundMeta(suid SoundUid) SoundMeta {
	path := fmt.Sprintf("sounds/%s/%s_meta.json", suid.ServerId, suid.SoundId)
	metaFile, err := os.Open(path)

	if err != nil {
		log.Fatal(err)
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

	for {
		id := strconv.FormatInt(int64(rand.Intn(max-min)+min), 10)

		if !SoundExists(serverId, id) {
			return id
		}
	}
}

func SoundExists(serverId string, soundId string) bool {
	return slices.Contains(GetSoundIDs(serverId), soundId)
}

func GetSoundMetas(serverId string) []SoundMeta {
	ids := GetSoundIDs(serverId)
	sounds := make([]SoundMeta, 0)

	for _, id := range ids {
		sounds = append(sounds, GetSoundMeta(SuidFromStrings(serverId, id)))
	}

	return sounds
}

func GetSoundIDs(serverId string) []string {
	path := fmt.Sprintf("sounds/%s", serverId)
	if !util.PathExists(path) {
		return make([]string, 0)
	}

	files, err := os.ReadDir(path)
	if err != nil {
		log.Fatal(err)
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

func GetAllServerIds() []string {
	files, err := os.ReadDir("sounds")
	if err != nil {
		log.Fatal(err)
	}

	ids := make([]string, 0)

	for _, file := range files {
		if file.IsDir() {
			ids = append(ids, file.Name())
		}
	}

	return ids
}

func GetAllSoundUids() []SoundUid {
	servers := GetAllServerIds()
	suids := make([]SoundUid, 0)

	for _, server := range servers {
		ids := GetSoundIDs(server)
		for _, id := range ids {
			suids = append(suids, SuidFromStrings(server, id))
		}
	}

	return suids
}

func AddSound(serverId string, name string, image string, soundUrl string) string {
	sound := SoundMeta{
		SoundId:  GenerateSoundID(serverId),
		ServerId: serverId,
		Name:     name,
		Icon:     image,
	}

	// Make sure the sounds folder exists
	path := fmt.Sprintf("sounds/%s", serverId)
	os.MkdirAll(path, os.ModePerm)

	// write meta file
	path = fmt.Sprintf("sounds/%s/%s_meta.json", serverId, sound.SoundId)
	metaFile, err := os.Create(path)
	if err != nil {
		log.Fatal(err)
	}

	metaJson, err := json.Marshal(sound)
	if err != nil {
		log.Fatal(err)
	}

	metaFile.Write(metaJson)
	metaFile.Close()

	// write sound file

	path = fmt.Sprintf("sounds/%s/%s_sound.mp3", serverId, sound.SoundId)

	err = util.DownloadSound(soundUrl, path)
	if err != nil {
		log.Fatal(err)
	}

	return sound.SoundId
}

func DeleteSound(serverId string, sound_id string) {
	path := fmt.Sprintf("sounds/%s/%s_meta.json", serverId, sound_id)
	err := os.Remove(path)
	if err != nil {
		log.Fatal(err)
	}

	path = fmt.Sprintf("sounds/%s/%s_sound.mp3", serverId, sound_id)
	err = os.Remove(path)
	if err != nil {
		log.Fatal(err)
	}
}

func LoadSoundBytes(suid SoundUid) ([]byte, error) {
	path := fmt.Sprintf("sounds/%s/%s_sound.s16le", suid.ServerId, suid.SoundId)
	f, err := os.Open(path)
	if err != nil {
		return nil, err
	}

	bytes, err := io.ReadAll(f)
	if err != nil {
		return nil, err
	}

	return bytes, nil
}
