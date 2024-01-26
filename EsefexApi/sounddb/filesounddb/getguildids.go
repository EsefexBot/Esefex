package filesounddb

import (
	"esefexapi/types"
	"log"
	"os"

	"github.com/pkg/errors"
)

// GetGuildIDs implements sounddb.SoundDB.
func (f *FileDB) GetGuildIDs() ([]types.GuildID, error) {
	files, err := os.ReadDir(f.location)
	if err != nil {
		log.Printf("Error reading directory: %+v", err)
		return nil, errors.Wrap(err, "Error reading directory")
	}

	ids := make([]types.GuildID, 0)

	for _, file := range files {
		if file.IsDir() {
			ids = append(ids, types.GuildID(file.Name()))
		}
	}

	return ids, nil
}
