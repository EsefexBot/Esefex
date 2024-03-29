package audioprocessing

import (
	"encoding/binary"
	"io"
	"os"

	"github.com/pkg/errors"
)

type S16leCacheReader struct {
	bytes  []byte
	cursor int
}

func (s *S16leCacheReader) Read(p []byte) (n int, err error) {
	if s.cursor >= len(s.bytes) {
		return 0, io.EOF
	}

	n = copy(p, s.bytes[s.cursor:])
	s.cursor += n
	return n, nil
}

func (s *S16leCacheReader) Load(bytes []byte) {
	s.bytes = bytes
	s.cursor = 0
}

func (s *S16leCacheReader) LoadFromReader(reader io.Reader) (err error) {
	var bytes []byte

	for {
		var short int16

		err = binary.Read(reader, binary.LittleEndian, &short)

		if err == nil {
			bytes = append(bytes, byte(short))
			bytes = append(bytes, byte(short>>8))
		} else if err != io.EOF {
			return errors.Wrap(err, "Error reading from reader")
		} else {
			break
		}
	}

	s.Load(bytes)
	return nil
}

func NewS16leCacheReader() *S16leCacheReader {
	return &S16leCacheReader{}
}

func NewS16leCacheReaderFromBytes(b []byte) *S16leCacheReader {
	reader := &S16leCacheReader{}
	reader.Load(b)

	return reader
}

func NewS16leCacheReaderFromPCM(pcm []int16) *S16leCacheReader {
	reader := &S16leCacheReader{}

	for _, short := range pcm {
		reader.bytes = append(reader.bytes, byte(short))
		reader.bytes = append(reader.bytes, byte(short>>8))
	}

	return reader
}

func NewS16leCacheReaderFromFile(p string) (*S16leCacheReader, error) {
	f, err := os.Open(p)
	if err != nil {
		panic(err)
	}

	reader := &S16leCacheReader{}
	err = reader.LoadFromReader(f)
	if err != nil {
		return nil, errors.Wrap(err, "Error loading from reader")
	}

	return reader, nil
}
