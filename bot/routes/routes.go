package routes

import (
	"encoding/json"
	"esefexbot/appcontext"
	"esefexbot/filedb"
	"fmt"
	"io"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

// api/sounds/<server_id>
func GetSounds(w http.ResponseWriter, r *http.Request, c appcontext.Context) {
	vars := mux.Vars(r)
	server_id := vars["server_id"]

	sounds := filedb.GetSounds(server_id)

	jsonResponse, err := json.Marshal(sounds)
	if err != nil {
		log.Fatal(err)
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(jsonResponse)

	fmt.Println("got /sounds request")
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
