package routes

import (
	"io"
	"net/http"
)

// /
func (routes *RouteHandlers) GetIndex(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "Index!\n")
}
