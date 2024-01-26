package main

import (
	"esefexapi/audioprocessing"
	"io"
	"os"
)

func main() {
	sounds := []string{"test1.s16le", "test2.s16le", "test3.s16le", "goofy_ahh.s16le", "oh_mah_gawd.s16le", "ohio_ahh.s16le"}

	mixReader := audioprocessing.S16leMixReader{}

	for _, sound := range sounds {
		file, err := os.Open(sound)
		if err != nil {
			panic(err)
		}

		cacheReader := audioprocessing.S16leCacheReader{}
		err = cacheReader.LoadFromReader(file)
		mixReader.AddSource(&cacheReader)
	}

	fileOut, err := os.Create("out.s16le")
	if err != nil {
		panic(err)
	}

	_, err = io.Copy(fileOut, &mixReader)
	if err != nil {
		panic(err)
	}
}
