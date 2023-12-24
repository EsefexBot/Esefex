package routes

import (
	"io"
	"net/http"
)

// /
func (routes *RouteHandlers) Index(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "Index!\n")
}
