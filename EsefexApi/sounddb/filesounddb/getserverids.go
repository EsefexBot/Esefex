package filesounddb

import (
	"log"
	"os"
)

// GetServerIDs implements sounddb.SoundDB.
func (f *FileDB) GetServerIDs() ([]string, error) {
	files, err := os.ReadDir(f.location)
	if err != nil {
		log.Print(err)
		return nil, err
	}

	ids := make([]string, 0)

	for _, file := range files {
		if file.IsDir() {
			ids = append(ids, file.Name())
		}
	}

	return ids, nil
}
