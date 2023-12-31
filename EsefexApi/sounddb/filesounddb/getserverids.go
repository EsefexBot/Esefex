package filesounddb

import (
	"log"
	"os"

	"github.com/pkg/errors"
)

// GetServerIDs implements sounddb.SoundDB.
func (f *FileDB) GetServerIDs() ([]string, error) {
	files, err := os.ReadDir(f.location)
	if err != nil {
		log.Printf("Error reading directory: %+v", err)
		return nil, errors.Wrap(err, "Error reading directory")
	}

	ids := make([]string, 0)

	for _, file := range files {
		if file.IsDir() {
			ids = append(ids, file.Name())
		}
	}

	return ids, nil
}
