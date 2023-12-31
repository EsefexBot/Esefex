package routes

import (
	"esefexapi/sounddb"
	"fmt"
	"io"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

// api/playsound/<user_id>/<server_id>/<sound_id>
func (h *RouteHandlers) PostPlaySoundInsecure(w http.ResponseWriter, r *http.Request) {
	log.Printf("got /playsound request\n")

	vars := mux.Vars(r)
	user_id := vars["user_id"]
	server_id := vars["server_id"]
	sound_id := vars["sound_id"]

	err := h.a.PlaySoundInsecure(sounddb.SuidFromStrings(server_id, sound_id), server_id, user_id)
	if err != nil {
		errorMsg := fmt.Sprintf("Error playing sound: %+v", err)
		log.Println(errorMsg)
		http.Error(w, errorMsg, http.StatusInternalServerError)
		return
	}

	w.Header().Set("Access-Control-Allow-Origin", "*")
	io.WriteString(w, "Play sound!\n")
}
