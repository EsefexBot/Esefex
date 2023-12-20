package routes

import (
	"encoding/json"
	"esefexapi/audioplayer"
	"esefexapi/sounddb"
	"fmt"
	"io"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type RouteHandlers struct {
	a      audioplayer.IAudioPlayer
	db     sounddb.ISoundDB
	cProto string
}

func NewRouteHandler(db sounddb.ISoundDB, a audioplayer.IAudioPlayer, cProto string) *RouteHandlers {
	return &RouteHandlers{
		a:      a,
		db:     db,
		cProto: cProto,
	}
}

// api/sounds/<server_id>
func (routes *RouteHandlers) GetSounds(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	server_id := vars["server_id"]

	uids, err := routes.db.GetSoundUIDs(server_id)
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(fmt.Sprintf("Error: %s", err)))
		return
	}

	sounds := make([]sounddb.SoundMeta, len(uids))
	for i, uid := range uids {
		m, err := routes.db.GetSoundMeta(uid)
		if err != nil {
			log.Println(err)
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(fmt.Sprintf("Error: %s", err)))
			return
		}

		sounds[i] = m
	}

	jsonResponse, err := json.Marshal(sounds)
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(fmt.Sprintf("Error: %s", err)))
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(jsonResponse)

	log.Println("got /sounds request")
}

// api/playsound/<user_id>/<server_id>/<sound_id>
func (routes *RouteHandlers) PlaySound(w http.ResponseWriter, r *http.Request) {
	log.Printf("got /playsound request\n")

	vars := mux.Vars(r)
	user_id := vars["user_id"]
	server_id := vars["server_id"]
	sound_id := vars["sound_id"]

	err := routes.a.PlaySound(sounddb.SuidFromStrings(server_id, sound_id), user_id)
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(fmt.Sprintf("Error: %s", err)))
		return
	}

	io.WriteString(w, "Play sound!\n")
}

// joinsession/<server_id>
func (routes *RouteHandlers) JoinSession(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	server_id := vars["server_id"]

	redirectUrl := fmt.Sprintf("%s://joinsession/%s", routes.cProto, server_id)
	response := fmt.Sprintf(`<meta http-equiv="refresh" content="0; URL=%s" />`, redirectUrl)

	w.Header().Set("Content-Type", "text/html")
	fmt.Fprint(w, response)

	log.Printf("got /joinsession request\n")
}

// /dump
func (routes *RouteHandlers) Dump(w http.ResponseWriter, r *http.Request) {
	log.Printf("%+v\n", r)

	io.WriteString(w, "Dump!\n")
}

// /
func (routes *RouteHandlers) Index(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "Index!\n")
}
