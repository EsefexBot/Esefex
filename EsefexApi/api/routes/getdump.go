package routes

import (
	"net/http"

	"github.com/davecgh/go-spew/spew"
)

// /dump
func (h *RouteHandlers) GetDump(w http.ResponseWriter, r *http.Request) {
	spew.Fdump(w, r)
	spew.Dump(r)
}
