package filesounddb

import (
	"esefexapi/sounddb"
	"math/rand"
	"strconv"

	"github.com/pkg/errors"
)

func (f *FileDB) generateSoundID(serverId string) (string, error) {
	// generate random number with 16 digits
	min := 100000000
	max := 999999999

	for {
		id := strconv.FormatInt(int64(rand.Intn(max-min)+min), 10)

		exists, err := f.SoundExists(sounddb.SuidFromStrings(serverId, id))
		if err != nil {
			return "", errors.Wrap(err, "Error checking if sound exists")
		}

		if !exists {
			return id, nil
		}
	}
}
