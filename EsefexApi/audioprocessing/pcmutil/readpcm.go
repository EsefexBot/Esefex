package pcmutil

import (
	"encoding/binary"
	"io"
)

// ReadPCM reads PCM data from r and stores it in buf.
// buf must be a slice of int16s.
// ReadPCM returns the number of bytes read and an error if any.
// ReadPCM assumes that the data is in little endian.
// if the io.Reader gives an EOF before the buffer is filled, ReadPCM returns io.EOF and the number of bytes read.
// unlike binary.Read, ReadPCM will still write to buf even if the io.Reader gives an EOF before the buffer is filled.
func ReadPCM(r io.Reader, buf *[]int16) (int, error) {
	for i := range *buf {
		err := binary.Read(r, binary.LittleEndian, &(*buf)[i])
		if err != nil {
			return i, err
		}
	}

	return len(*buf), nil
}
