package routes

import (
	"esefexapi/types"
	"esefexapi/util/dcgoutil"
	"net/http"
)

// api/guild
// returns the guild a user is connected to
func (h *RouteHandlers) GetGuild(w http.ResponseWriter, r *http.Request, userID types.UserID) {
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
	w.Write([]byte(guildID))
}
