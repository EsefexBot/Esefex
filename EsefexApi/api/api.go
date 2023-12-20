package api

import (
	"esefexapi/api/routes"
	"esefexapi/audioplayer"
	"esefexapi/sounddb"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type HttpApi struct {
	db      sounddb.ISoundDB
	a       audioplayer.IAudioPlayer
	apiPort int
	cProto  string
	stop    chan struct{}
}

func NewHttpApi(db sounddb.ISoundDB, a audioplayer.IAudioPlayer, apiPort int, cProto string) *HttpApi {
	return &HttpApi{
		db:      db,
		a:       a,
		apiPort: apiPort,
		cProto:  cProto,
		stop:    make(chan struct{}, 1),
	}
}

func (api *HttpApi) run() {
	log.Println("Starting webserver...")

	routes := routes.NewRouteHandler(api.db, api.a, api.cProto)

	router := mux.NewRouter()

	router.HandleFunc("/api/sounds/{server_id}", routes.GetSounds)
	router.HandleFunc("/api/playsound/{user_id}/{server_id}/{sound_id}", routes.PlaySound)

	router.HandleFunc("/joinsession/{server_id}", routes.JoinSession)

	router.HandleFunc("/dump", routes.Dump)
	router.HandleFunc("/", routes.Index)

	// http.Handle("/", router)
	log.Printf("Webserver started on port %d\n", api.apiPort)
	go http.ListenAndServe(fmt.Sprintf(":%d", api.apiPort), router)

	<-api.stop
}

func (api *HttpApi) Start() {
	go api.run()
}

func (api *HttpApi) Stop() {
	close(api.stop)
}
