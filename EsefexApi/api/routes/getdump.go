package routes

import (
	"net/http"

	"github.com/davecgh/go-spew/spew"
)

// /dump
func (h *RouteHandlers) GetDump() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		spew.Fdump(w, r)
		spew.Dump(r)
	})
}
