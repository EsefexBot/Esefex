package routes

import (
	"io"
	"log"
	"net/http"
)

// /dump
func (h *RouteHandlers) GetDump(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "Dump!\n")

	log.Printf("%+v\n", r)
}
