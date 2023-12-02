package main

import (
	"esefexbot/audioprocessing"
	"os"
)

func main() {
	// load data from file
	f, err := os.Open("test.s16le")
	if err != nil {
		panic(err)
	}
	defer f.Close()

	stat, err := f.Stat()
	if err != nil {
		panic(err)
	}

	// read data to byte slice
	bytes := make([]byte, stat.Size())
	_, err = f.Read(bytes)
	if err != nil {
		panic(err)
	}

	// convert byte slice to int16 slice
	var shorts []int16
	shorts = audioprocessing.AsPCMs16le(bytes)

}
