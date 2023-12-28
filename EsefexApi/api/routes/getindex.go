package routes

import (
	"io"
	"net/http"
)

// /
func (h *RouteHandlers) GetIndex(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	io.WriteString(w, "Index!\n")
}
