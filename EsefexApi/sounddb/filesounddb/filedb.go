package filesounddb

import (
	"esefexapi/config"
	"esefexapi/sounddb"
	"esefexapi/types"
	"fmt"
	"log"
	"os"

	"github.com/pkg/errors"
)

var _ sounddb.ISoundDB = &FileDB{}

// FileDB implements SoundDB
type FileDB struct {
	location string
}

// Gets the sound Name from the sound ID
// GetSoundNameByID implements sounddb.ISoundDB.
func (f *FileDB) GetSoundNameByID(guildID types.GuildID, ID types.SoundID) (types.SoundName, error) {
	soundUIDs, err := f.GetSoundUIDs(guildID)
	if err != nil {
		return "", errors.Wrap(err, "Error getting sound UIDs")
	}

	for _, soundUID := range soundUIDs {
		meta, err := f.GetSoundMeta(soundUID)
		if err != nil {
			return "", errors.Wrap(err, "Error getting sound meta")
		}

		if meta.SoundID == ID {
			return types.SoundName(meta.Name), nil
		}
	}

	return "", errors.Wrap(fmt.Errorf("Sound not found"), "Error getting sound name by ID")
}

// NewFileDB returns a new FileDB
func NewFileDB() (*FileDB, error) {
	location := config.Get().Database.SounddbLocation

	log.Printf("Creating FileDB at %s", location)
	err := os.MkdirAll(location, os.ModePerm)
	if err != nil {
		return nil, errors.Wrap(err, "Error creating directory")
	}

	return &FileDB{
		location: location,
	}, nil
}
