package routes

import (
	"io"
	"log"
	"net/http"
)

// /dump
func (routes *RouteHandlers) GetDump(w http.ResponseWriter, r *http.Request) {
	log.Printf("%+v\n", r)

	io.WriteString(w, "Dump!\n")
}
