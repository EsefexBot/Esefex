package main

import (
	"esefexapi/audioprocessing"
	"esefexapi/util"
	"log"
)

func main() {
	src := audioprocessing.S16leFromFile("testsounds/test1.s16le")
	enc, err := audioprocessing.NewOpusEncoder(src)
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
