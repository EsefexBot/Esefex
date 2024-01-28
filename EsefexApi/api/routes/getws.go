package routes

import (
	"esefexapi/types"
	"log"
	"net/http"

	"github.com/gorilla/websocket"
)

var wsUpgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

// api/ws
func (h *RouteHandlers) GetWs() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		userID := r.Context().Value("user").(types.UserID)

		conn, err := wsUpgrader.Upgrade(w, r, nil)
		if err != nil {
			log.Printf("Error upgrading websocket: %v", err)
			return
		}

		h.wsCN.AddConnection(userID, conn)
	})
}
