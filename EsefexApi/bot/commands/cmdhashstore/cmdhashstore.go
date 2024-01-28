package cmdhashstore

import (
	"esefexapi/util"
	"io"
	"os"

	"github.com/pkg/errors"
)

type ICommandHashStore interface {
	GetCommandHash() (string, error)
	SetCommandHash(hash string) error
}

type FileCmdHashStore struct {
	FilePath string
}

func NewFileCmdHashStore(filePath string) *FileCmdHashStore {
	return &FileCmdHashStore{
		FilePath: filePath,
	}
}

func (f *FileCmdHashStore) GetCommandHash() (string, error) {
	file, err := util.EnsureFile(f.FilePath)
	if err != nil {
		return "", err
	}
	defer file.Close()

	hash, err := io.ReadAll(file)
	if err != nil {
		return "", err
	}

	return string(hash), nil
}

func (f *FileCmdHashStore) SetCommandHash(hash string) error {
	file, err := os.Create(f.FilePath)
	if err != nil {
		return errors.Wrap(err, "error creating file")
	}
	defer file.Close()

	_, err = file.WriteString(hash)
	if err != nil {
		return errors.Wrap(err, "error writing to file")
	}

	return nil
}
