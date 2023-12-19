package filedb

import (
	"esefexapi/db"
	"esefexapi/util"
	"fmt"
	"log"
	"os"
	"strings"
)

func (f *FileDB) GetSoundUIDs(serverID string) ([]db.SoundUID, error) {
	path := fmt.Sprintf("sounds/%s", serverID)
	if !util.PathExists(path) {
		return make([]db.SoundUID, 0), nil
	}

	files, err := os.ReadDir(path)
	if err != nil {
		log.Fatal(err)
	}

	uids := make([]db.SoundUID, 0)

	for _, file := range files {
		name := file.Name()
		nameSplits := strings.Split(name, "_")

		if len(nameSplits) == 2 && nameSplits[1] == "meta.json" {
			uids = append(uids, db.SuidFromStrings(serverID, nameSplits[0]))
		}
	}

	return uids, nil
}
