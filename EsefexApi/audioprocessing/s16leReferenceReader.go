package audioprocessing

import (
	"io"
	"log"
)

type S16leReferenceReader struct {
	data   *[]int16
	cursor int
}

func (s *S16leReferenceReader) Read(p []byte) (n int, err error) {
	if s.cursor >= len(*s.data)*2 {
		log.Println("EOF")
		return 0, io.EOF
	}

	n = 0

	for i := range p {
		b, err := s.getByte(s.cursor)
		if err != nil {
			break
		}

		p[i] = b
		s.cursor++
		n++
	}

	return n, nil
}

func (s *S16leReferenceReader) getByte(i int) (byte, error) {
	if i >= len(*s.data)*2 {
		return 0, io.EOF
	}

	if i%2 == 0 {
		return byte((*s.data)[i/2] & 0xff), nil
	} else {
		return byte((*s.data)[i/2] >> 8), nil
	}
}

func (s *S16leReferenceReader) Load(data *[]int16) {
	s.data = data
	s.cursor = 0
}

func NewS16leReferenceReader() *S16leReferenceReader {
	return &S16leReferenceReader{}
}

func NewS16leReferenceReaderFromRef(pcm *[]int16) *S16leReferenceReader {
	reader := &S16leReferenceReader{}
	reader.Load(pcm)

	return reader
}
