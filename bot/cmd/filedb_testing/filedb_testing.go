package main

import (
	"esefexbot/filedb"
	"os"
	"time"
)

func main() {
	// sounds := filedb.GetSounds("34234234")

	// for _, sound := range sounds {
	// 	println(sound.Id)
	// 	println(sound.ServerId)
	// 	println(sound.Name)
	// 	println(sound.Icon)
	// }

	serverId := "7657"

	f2uPath := "upload_sound.mp3"

	f2u, err := os.ReadFile(f2uPath)
	if err != nil {
		println(err)
	}

	id := filedb.AddSound(serverId, "mogus", "https://image.com/123", f2u)

	time.Sleep(time.Duration(2500) * time.Millisecond)

	println(filedb.SoundExists(serverId, id))

	// time.Sleep(time.Duration(2500) * time.Millisecond)

	// filedb.DeleteSound(serverId, id)

	// time.Sleep(time.Duration(2500) * time.Millisecond)

	// println(filedb.SoundExists(serverId, id))
}
