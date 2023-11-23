package main

import (
	"os"
	"time"
	"webserver/filedb"
)

// "net/http"

// c "webserver/appcontext"
// r "webserver/routes"

// "github.com/gorilla/mux"

func main() {
	// context := c.NewContext()

	// router := mux.NewRouter()

	// router.HandleFunc("/api/sounds/{server_id}", c.Wrap(r.GetSounds, context))
	// router.HandleFunc("/api/playsound/{server_id}/{sound_id}", c.Wrap(r.PlaySound, context))

	// router.HandleFunc("/dump", c.Wrap(r.Dump, context))

	// // http.Handle("/", router)
	// http.ListenAndServe(":8080", router)

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

	time.Sleep(time.Duration(2500) * time.Millisecond)

	filedb.DeleteSound(serverId, id)

	time.Sleep(time.Duration(2500) * time.Millisecond)

	println(filedb.SoundExists(serverId, id))
}
