package api

import (
	"esefexbot/api/routes"
	"esefexbot/appcontext"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func Run(c *appcontext.Context) {
	log.Println("Starting webserver...")

	router := mux.NewRouter()

	router.HandleFunc("/api/sounds/{server_id}", appcontext.Wrap(routes.GetSounds, c))
	router.HandleFunc("/api/playsound/{server_id}/{sound_id}", appcontext.Wrap(routes.PlaySound, c))

	router.HandleFunc("/joinsession/{server_id}", appcontext.Wrap(routes.JoinSession, c))

	router.HandleFunc("/dump", appcontext.Wrap(routes.Dump, c))

	// http.Handle("/", router)
	log.Printf("Webserver started on port %s\n", c.ApiPort)
	go http.ListenAndServe(fmt.Sprintf(":%s", c.ApiPort), router)

	<-c.Channels.Stop
}
