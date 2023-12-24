package filedb

import (
	"esefexapi/sounddb"
	"esefexapi/util"
	"fmt"
	"log"
	"os"
	"strings"
)

// GetSoundUIDs implements sounddb.SoundDB.
func (f *FileDB) GetSoundUIDs(serverID string) ([]sounddb.SoundUID, error) {
	path := fmt.Sprintf("%s/%s", f.location, serverID)
	if !util.PathExists(path) {
		return make([]sounddb.SoundUID, 0), nil
	}

	files, err := os.ReadDir(path)
	if err != nil {
		log.Fatal(err)
	}

	uids := make([]sounddb.SoundUID, 0)

	for _, file := range files {
		name := file.Name()
		nameSplits := strings.Split(name, "_")

		if len(nameSplits) == 2 && nameSplits[1] == "meta.json" {
			uids = append(uids, sounddb.SuidFromStrings(serverID, nameSplits[0]))
		}
	}

	return uids, nil
}
