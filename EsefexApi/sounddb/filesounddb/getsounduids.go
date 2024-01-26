package filesounddb

import (
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
func (f *FileDB) GetSoundUIDs(guildID types.GuildID) ([]sounddb.SoundURI, error) {
	path := fmt.Sprintf("%s/%s", f.location, guildID)
	if !util.PathExists(path) {
		return make([]sounddb.SoundURI, 0), nil
	}

	files, err := os.ReadDir(path)
	if err != nil {
		log.Printf("Error reading directory: %+v", err)
		return nil, errors.Wrap(err, "Error reading directory")
	}

	uids := make([]sounddb.SoundURI, 0)

	for _, file := range files {
		name := file.Name()
		nameSplits := strings.Split(name, "_")

		if len(nameSplits) == 2 && nameSplits[1] == "meta.json" {
			uids = append(uids, sounddb.SuidFromStrings(guildID.String(), nameSplits[0]))
		}
	}

	return uids, nil
}
