package filesounddb

import (
	"encoding/json"
	"esefexapi/sounddb"
	"esefexapi/types"
	"esefexapi/util"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/pkg/errors"
)

// GetSoundUIDs implements sounddb.SoundDB.
func (f *FileDB) GetSoundUIDs(guildID types.GuildID) ([]sounddb.SoundUID, error) {
	path := fmt.Sprintf("%s/%s", f.location, guildID)

	pathExists, err := util.PathExists(path)
	if err != nil {
		return nil, errors.Wrap(err, "Error checking if path exists")
	}

	if !pathExists {
		return make([]sounddb.SoundUID, 0), nil
	}

	files, err := os.ReadDir(path)
	if err != nil {
		log.Printf("Error reading directory: %+v", err)
		return nil, errors.Wrap(err, "Error reading directory")
	}

	uids := make([]sounddb.SoundUID, 0)

	for _, file := range files {
		name := file.Name()
		// if its a meta file parse to json
		if strings.HasSuffix(name, "_meta.json") {
			// read the file
			file, err := os.ReadFile(fmt.Sprintf("%s/%s", path, name))
			if err != nil {
				return nil, errors.Wrap(err, "Error reading meta file")
			}

			soundMeta := sounddb.SoundMeta{}

			err = json.Unmarshal(file, &soundMeta)
			if err != nil {
				return nil, errors.Wrap(err, "Error reading meta file")
			}

			uids = append(uids, sounddb.SoundUID{
				GuildID:   guildID,
				SoundName: soundMeta.Name,
			})
		}
	}

	return uids, nil
}
