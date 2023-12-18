package audioprocessing

import (
	"encoding/binary"
	"io"
	"os"
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
			return err
		} else {
			break
		}
	}

	s.Load(bytes)
	return nil
}

func S16leFromFile(p string) *S16leCacheReader {
	f, err := os.Open(p)
	if err != nil {
		panic(err)
	}

	reader := &S16leCacheReader{}
	reader.LoadFromReader(f)

	return reader
}
