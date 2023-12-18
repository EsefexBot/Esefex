package main

import (
	"log"

	"layeh.com/gopus"
)

func main() {
	_, err := gopus.NewEncoder(48000, 2, gopus.Voip)
	if err != nil {
		log.Println("Failure :(")
		panic(err)
	}

	log.Println("Success!")

	// a, err := enc.Encode([]int16{1, 2, 3, 4}, 960, 960*2*2)
	// if err != nil {
	// 	panic(err)
	// }

	// log.Println(a)
}
