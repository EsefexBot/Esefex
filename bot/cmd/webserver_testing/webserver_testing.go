package main

import (
	"net/http"

	c "esefexbot/appcontext"
	r "esefexbot/routes"

	"github.com/gorilla/mux"
)

func main() {
	println("Starting webserver...")

	context := c.NewContext()

	router := mux.NewRouter()

	router.HandleFunc("/api/sounds/{server_id}", c.Wrap(r.GetSounds, context))
	router.HandleFunc("/api/playsound/{server_id}/{sound_id}", c.Wrap(r.PlaySound, context))

	router.HandleFunc("/dump", c.Wrap(r.Dump, context))

	// http.Handle("/", router)
	println("Webserver started!")
	http.ListenAndServe(":8080", router)

}
