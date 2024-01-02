package audioprocessing

import (
	"encoding/binary"
	"io"

	"github.com/pkg/errors"
)

type S16leMixReader struct {
	sources []*io.Reader
}

func NewS16leMixReader() *S16leMixReader {
	return &S16leMixReader{}
}

func (s *S16leMixReader) Read(p []byte) (n int, err error) {
	if s.sources == nil {
		return 0, io.EOF
	}

	if len(p)%2 != 0 {
		return 0, io.ErrShortBuffer
	}

	for i := 0; i < len(p); i += 2 {
		if len(s.sources) == 0 {
			return i, nil
		}

		var shorts []int16
		var remainingSources []*io.Reader

		for j, source := range s.sources {
			var short int16

			err = binary.Read(*source, binary.LittleEndian, &short)

			if err == nil {
				shorts = append(shorts, short)
				remainingSources = append(remainingSources, s.sources[j])
			} else if err != io.EOF {
				return 0, errors.Wrap(err, "Error reading from source")
			}
		}

		mix := MixPCMs16leClip(shorts)

		p[i] = byte(mix)
		p[i+1] = byte(mix >> 8)

		s.sources = remainingSources
	}

	return len(p), nil
}

func (s *S16leMixReader) RemoveSources(rs []io.Reader) {
	for _, r := range rs {
		for i, source := range s.sources {
			if *source == r {
				s.sources = append(s.sources[:i], s.sources[i+1:]...)
				break
			}
		}
	}
}

func (s *S16leMixReader) AddSource(source io.Reader) {
	s.sources = append(s.sources, &source)
}

func (s *S16leMixReader) SourceCount() int {
	return len(s.sources)
}

func (s *S16leMixReader) Empty() bool {
	return len(s.sources) == 0
}
