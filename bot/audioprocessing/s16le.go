package audioprocessing

import "io"

type s16leReader struct {
	bytes  []byte
	cursor int
}

func (s *s16leReader) Read(p []byte) (n int, err error) {
	if s.cursor >= len(s.bytes) {
		return 0, io.EOF
	}

	n = copy(p, s.bytes[s.cursor:])
	s.cursor += n
	return n, nil
}
