package filedb

import (
	"esefexapi/sounddb"
)

var _ sounddb.ISoundDB = &FileDB{}

// FileDB implements SoundDB
type FileDB struct{}

// NewFileDB returns a new FileDB
func NewFileDB() *FileDB {
	return &FileDB{}
}
