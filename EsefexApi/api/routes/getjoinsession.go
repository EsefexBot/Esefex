package routes

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

// joinsession/<guild_id>
func (h *RouteHandlers) GetJoinSession() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		guild_id := vars["guild_id"]

		redirectUrl := fmt.Sprintf("%s://joinsession/%s", h.cProto, guild_id)
		response := fmt.Sprintf(`<meta http-equiv="refresh" content="0; URL=%s" />`, redirectUrl)

		w.Header().Set("Content-Type", "text/html")
		fmt.Fprint(w, response)

		log.Printf("got /joinsession request\n")
	})
}
