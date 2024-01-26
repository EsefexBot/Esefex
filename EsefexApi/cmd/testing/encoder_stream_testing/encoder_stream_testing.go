package main

import (
	"esefexapi/audioprocessing"
	"esefexapi/util"
	"log"
)

func main() {
	src, err := audioprocessing.NewS16leCacheReaderFromFile("testsounds/test1.s16le")
	if err != nil {
		panic(err)
	}

	enc, err := audioprocessing.NewOpusCliEncoder(src)
	if err != nil {
		panic(err)
	}

	for {
		b, err := enc.EncodeNext()
		if err != nil {
			panic(err)
		}

		util.PrintBytesCustom(b, 32)
		log.Printf("framelen: %d\n", len(b))
	}
}
