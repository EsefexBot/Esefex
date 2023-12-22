package util

import (
	"crypto/rand"
	"encoding/binary"
	"fmt"
	"io"
	"os"
	"path/filepath"
)

func LoadDcaSound(path string) ([][]byte, error) {
	var buffer [][]byte

	if filepath.Ext(path) != ".dca" {
		return nil, fmt.Errorf("file is not a .dca file")
	}

	file, err := os.Open(path)
	if err != nil {
		fmt.Println("Error opening dca file :", err)
		return nil, err
	}

	var opuslen int16

	for {
		// Read opus frame length from dca file.
		err = binary.Read(file, binary.LittleEndian, &opuslen)

		// If this is the end of the file, just return.
		if err == io.EOF || err == io.ErrUnexpectedEOF {
			err := file.Close()
			if err != nil {
				return nil, err
			}
			return buffer, nil
		}

		if err != nil {
			fmt.Println("Error reading from dca file :", err)
			return nil, err
		}

		// Read encoded pcm from dca file.
		InBuf := make([]byte, opuslen)
		err = binary.Read(file, binary.LittleEndian, &InBuf)

		// Should not be any end of file errors
		if err != nil {
			fmt.Println("Error reading from dca file :", err)
			return nil, err
		}

		// Append encoded pcm data to the buffer.
		buffer = append(buffer, InBuf)
	}
}

func HallucinateDcaData(len int, opuslen int) [][]byte {
	var sound [][]byte

	for i := 0; i < len; i++ {
		buf := make([]byte, opuslen)
		rand.Read(buf)
		sound = append(sound, buf)
	}

	return sound
}

func ConstantDcaData(len int, opuslen int, val byte) [][]byte {
	var sound [][]byte

	for i := 0; i < len; i++ {
		buf := make([]byte, opuslen)
		for j := 0; j < opuslen; j++ {
			buf[j] = val
		}
		sound = append(sound, buf)
	}

	return sound
}
