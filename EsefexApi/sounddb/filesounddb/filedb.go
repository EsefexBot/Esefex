package filesounddb

import (
	"esefexapi/config"
	"esefexapi/sounddb"
	"log"
	"os"

	"github.com/pkg/errors"
)

var _ sounddb.ISoundDB = &FileDB{}

// FileDB implements SoundDB
type FileDB struct {
	location string
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
