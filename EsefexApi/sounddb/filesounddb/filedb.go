package filesounddb

import (
	"esefexapi/sounddb"
	"log"
	"os"
)

var _ sounddb.ISoundDB = &FileDB{}

// FileDB implements SoundDB
type FileDB struct {
	location string
}

// NewFileDB returns a new FileDB
func NewFileDB(location string) (*FileDB, error) {
	log.Printf("Creating FileDB at %s", location)
	err := os.MkdirAll(location, os.ModePerm)
	if err != nil {
		return nil, err
	}

	return &FileDB{
		location: location,
	}, nil
}
