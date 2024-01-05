package routes

import (
	"encoding/json"
	"esefexapi/sounddb"
	"esefexapi/types"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

// api/sounds/<guild_id>
func (h *RouteHandlers) GetSounds(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	guild_id := types.GuildID(vars["guild_id"])

	uids, err := h.dbs.SoundDB.GetSoundUIDs(guild_id)
	if err != nil {
		errorMsg := fmt.Sprintf("Error getting sound uids: %+v", err)

		log.Print(errorMsg)
		http.Error(w, errorMsg, http.StatusInternalServerError)
		return
	}

	sounds := make([]sounddb.SoundMeta, len(uids))
	for i, uid := range uids {
		m, err := h.dbs.SoundDB.GetSoundMeta(uid)
		if err != nil {
			errorMsg := fmt.Sprintf("Error getting sound meta: %+v", err)

			log.Println(errorMsg)
			http.Error(w, errorMsg, http.StatusInternalServerError)
			return
		}

		sounds[i] = m
	}

	jsonResponse, err := json.Marshal(sounds)
	if err != nil {
		errorMsg := fmt.Sprintf("Error marshalling json: %+v", err)

		log.Println(errorMsg)
		http.Error(w, errorMsg, http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(jsonResponse)

	log.Println("got /sounds request")
}
