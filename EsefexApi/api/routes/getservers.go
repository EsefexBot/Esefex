package routes

import (
	"encoding/json"
	"esefexapi/util/dcgoutil"
	"log"
	"net/http"
)

type ServerInfo struct {
	ServerID   string `json:"serverID"`
	ServerName string `json:"serverName"`
}

func (h *RouteHandlers) GetServers(w http.ResponseWriter, r *http.Request, userID string) {
	guilds, err := dcgoutil.UserServers(h.ds, userID)
	if err != nil {
		log.Println(err)
		http.Error(w, "Error getting user servers", http.StatusInternalServerError)
		return
	}

	si := []ServerInfo{}

	for _, g := range guilds {
		si = append(si, ServerInfo{ServerID: g.ID, ServerName: g.Name})
	}

	js, err := json.Marshal(si)
	if err != nil {
		log.Println(err)
		http.Error(w, "Error marshalling json", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(js)
}
