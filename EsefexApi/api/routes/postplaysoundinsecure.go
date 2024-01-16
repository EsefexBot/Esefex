package routes

import (
	"esefexapi/sounddb"
	"esefexapi/types"
	"fmt"
	"io"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

// api/playsound/<user_id>/<guild_id>/<sound_id>
func (h *RouteHandlers) PostPlaySoundInsecure() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Printf("got /playsound request\n")

		vars := mux.Vars(r)
		user_id := types.UserID(vars["user_id"])
		guild_id := types.GuildID(vars["guild_id"])
		sound_id := types.SoundID(vars["sound_id"])

		err := h.a.PlaySoundInsecure(sounddb.New(guild_id, sound_id), guild_id, user_id)
		if err != nil {
			errorMsg := fmt.Sprintf("Error playing sound: %+v", err)
			log.Println(errorMsg)
			http.Error(w, errorMsg, http.StatusInternalServerError)
			return
		}

		_, err = io.WriteString(w, "Play sound!\n")
		if err != nil {
			log.Println(err)
		}
	})
}
