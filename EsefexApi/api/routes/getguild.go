package routes

import (
	"esefexapi/types"
	"esefexapi/util/dcgoutil"
	"log"
	"net/http"
)

// api/guild
// returns the guild a user is connected to
func (h *RouteHandlers) GetGuild() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		userID := r.Context().Value("user").(types.UserID)
		Ovs, err := dcgoutil.UserVCAny(h.ds, userID)
		if err != nil {
			errorMsg := "Error getting user Voice State"
			http.Error(w, errorMsg, http.StatusInternalServerError)
			return
		}
		if Ovs.IsNone() {
			http.Error(w, "User not connected to guild channel", http.StatusForbidden)
			return
		}

		guildID := Ovs.Unwrap().GuildID
		_, err = w.Write([]byte(guildID))
		if err != nil {
			log.Println(err)
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
	})
}
