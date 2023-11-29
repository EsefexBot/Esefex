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

	routes := routes.Routes{C: c}

	router := mux.NewRouter()

	router.HandleFunc("/api/sounds/{server_id}", routes.GetSounds)
	router.HandleFunc("/api/playsound/{server_id}/{sound_id}", routes.PlaySound)

	router.HandleFunc("/joinsession/{server_id}", routes.JoinSession)

	router.HandleFunc("/dump", routes.Dump)

	// http.Handle("/", router)
	log.Printf("Webserver started on port %s\n", c.ApiPort)
	go http.ListenAndServe(fmt.Sprintf(":%s", c.ApiPort), router)

	<-c.Channels.Stop
}
