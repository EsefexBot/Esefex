package routes

import (
	"encoding/json"
	"esefexapi/sounddb"
	"esefexapi/types"
	"esefexapi/util/dcgoutil"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

// api/sounds/<guild_id>
func (h *RouteHandlers) GetSounds() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		guildId := types.GuildID(vars["guild_id"])
		userID := r.Context().Value("user").(types.UserID)

		// check that the user is in the guild
		inGuild, err := dcgoutil.UserInGuild(h.ds, guildId, userID)
		if err != nil {
			errorMsg := fmt.Sprintf("Error checking if user is in guild: %+v", err)

			log.Println(errorMsg)
			http.Error(w, errorMsg, http.StatusInternalServerError)
			return
		}

		if !inGuild {
			errorMsg := fmt.Sprintf("User is not in guild %s", guildId)

			log.Println(errorMsg)
			http.Error(w, errorMsg, http.StatusForbidden)
			return
		}

		uids, err := h.dbs.SoundDB.GetSoundUIDs(guildId)
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
		_, err = w.Write(jsonResponse)
		if err != nil {
			log.Println(err)
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
	})
}
