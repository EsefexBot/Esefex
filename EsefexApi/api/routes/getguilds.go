package routes

import (
	"encoding/json"
	"esefexapi/types"
	"esefexapi/util/dcgoutil"
	"log"
	"net/http"
)

type GuildInfo struct {
	GuildID   string `json:"guildID"`
	GuildName string `json:"guildName"`
}

func (h *RouteHandlers) GetGuilds(w http.ResponseWriter, r *http.Request, userID types.UserID) {
	guilds, err := dcgoutil.UserGuilds(h.ds, userID)
	if err != nil {
		log.Println(err)
		http.Error(w, "Error getting user guilds", http.StatusInternalServerError)
		return
	}

	si := []GuildInfo{}

	for _, g := range guilds {
		si = append(si, GuildInfo{GuildID: g.ID, GuildName: g.Name})
	}

	js, err := json.Marshal(si)
	if err != nil {
		log.Println(err)
		http.Error(w, "Error marshalling json", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	_, err = w.Write(js)
	if err != nil {
		log.Println(err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
