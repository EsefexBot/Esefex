package routes

import (
	"esefexapi/util/dcgoutil"
	"net/http"
)

// api/server
// returns the server a user is connected to
func (h *RouteHandlers) GetServer(w http.ResponseWriter, r *http.Request, userID string) {
	Oserver, err := dcgoutil.UserVC(h.ds, userID)
	if err != nil {
		errorMsg := "Error getting user server"
		http.Error(w, errorMsg, http.StatusInternalServerError)
		return
	}
	if Oserver.IsNone() {
		http.Error(w, "User not connected to server", http.StatusNotFound)
		return
	}

	serverID := Oserver.Unwrap().GuildID
	w.Write([]byte(serverID))
}
