package routes

import (
	"encoding/json"
	"esefexbot/appcontext"
	"esefexbot/filedb"
	"esefexbot/msg"

	// "esefexbot/msg"
	"fmt"
	"io"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type Routes struct {
	C *appcontext.Context
}

// api/sounds/<server_id>
func (routes *Routes) GetSounds(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	server_id := vars["server_id"]

	sounds := filedb.GetSounds(server_id)

	jsonResponse, err := json.Marshal(sounds)
	if err != nil {
		log.Fatal(err)
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(jsonResponse)

	log.Println("got /sounds request")
}

// api/playsound/<server_id>/<sound_id>
func (routes *Routes) PlaySound(w http.ResponseWriter, r *http.Request) {
	log.Printf("got /playsound request\n")

	vars := mux.Vars(r)

	log.Println(routes.C.Channels.PlaySound)

	routes.C.Channels.PlaySound <- msg.PlaySound{
		GuildID: vars["server_id"],
		SoundID: vars["sound_id"],
	}

	io.WriteString(w, "Play sound!\n")
}

// joinsession/<server_id>
func (routes *Routes) JoinSession(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	server_id := vars["server_id"]

	redirectUrl := fmt.Sprintf("%s://joinsession/%s", routes.C.CustomProtocol, server_id)
	response := fmt.Sprintf(`<meta http-equiv="refresh" content="0; URL=%s" />`, redirectUrl)

	w.Header().Set("Content-Type", "text/html")
	fmt.Fprint(w, response)

	log.Printf("got /joinsession request\n")
}

func (routes *Routes) Dump(w http.ResponseWriter, r *http.Request) {
	log.Printf("%+v\n", routes.C)
	log.Printf("%+v\n", r)

	io.WriteString(w, "Dump!\n")
}
