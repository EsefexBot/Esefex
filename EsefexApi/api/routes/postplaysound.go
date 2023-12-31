package routes

import (
	"esefexapi/userdb"
	"fmt"
	"io"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

// api/playsound/<sound_id>
func (h *RouteHandlers) PostPlaySound(w http.ResponseWriter, r *http.Request) {
	log.Printf("got /playsound request\n")

	vars := mux.Vars(r)
	sound_id := vars["sound_id"]

	user_token := r.Header.Get("User-Token")
	user, err := h.dbs.UserDB.GetUserByToken(userdb.Token(user_token))
	if err != nil {
		errorMsg := fmt.Sprintf("Error getting user by token: %+v", err)

		log.Println(errorMsg)
		http.Error(w, errorMsg, http.StatusUnauthorized)
		return
	}

	err = h.a.PlaySound(sound_id, user.ID)
	if err != nil {
		errorMsg := fmt.Sprintf("Error playing sound: %+v", err)

		log.Println(errorMsg)
		http.Error(w, errorMsg, http.StatusInternalServerError)
		return
	}

	w.Header().Set("Access-Control-Allow-Origin", "*")
	io.WriteString(w, "Play sound!\n")
}
