package routes

import (
	"io"
	"net/http"
)

// /
func (h *RouteHandlers) GetIndex() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		_, _ = io.WriteString(w, "Index!\n")
	})
}
