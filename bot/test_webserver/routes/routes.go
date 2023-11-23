package routes

import (
	"fmt"
	"io"
	"net/http"
	"webserver/appcontext"

	"github.com/gorilla/mux"
)

// api/sounds/<server_id>
func GetSounds(w http.ResponseWriter, r *http.Request, c appcontext.Context) {
	vars := mux.Vars(r)
	server_id := vars["server_id"]

	response := fmt.Sprintf("Get sounds for server %s", server_id)

	fmt.Println(response)
	io.WriteString(w, response)
}

// api/playsound/<server_id>/<sound_id>
func PlaySound(w http.ResponseWriter, r *http.Request, c appcontext.Context) {
	fmt.Printf("got /playsound request\n")
	io.WriteString(w, "Play sound!\n")
}

func Dump(w http.ResponseWriter, r *http.Request, c appcontext.Context) {
	fmt.Printf("%+v\n", c)
	fmt.Printf("%+v\n", r)

	io.WriteString(w, "Dump!\n")
}