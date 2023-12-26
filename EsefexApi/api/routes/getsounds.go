package routes

import (
	"encoding/json"
	"esefexapi/sounddb"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

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
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Write(jsonResponse)

	log.Println("got /sounds request")
}
