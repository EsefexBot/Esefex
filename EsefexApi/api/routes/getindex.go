package routes

import (
	"io"
	"net/http"
)

// /
func (h *RouteHandlers) GetIndex(w http.ResponseWriter, r *http.Request) {
	_, _ = io.WriteString(w, "Index!\n")
}
