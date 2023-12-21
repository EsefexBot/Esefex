package api

import (
	"esefexapi/api/routes"
	"esefexapi/audioplayer"
	"esefexapi/service"
	"esefexapi/sounddb"

	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

var _ service.Service = &HttpApi{}

// HttpApi implements Service
type HttpApi struct {
	db      sounddb.ISoundDB
	a       audioplayer.IAudioPlayer
	apiPort int
	cProto  string
	stop    chan struct{}
	ready   chan struct{}
}

func NewHttpApi(db sounddb.ISoundDB, a audioplayer.IAudioPlayer, apiPort int, cProto string) *HttpApi {
	return &HttpApi{
		db:      db,
		a:       a,
		apiPort: apiPort,
		cProto:  cProto,
		stop:    make(chan struct{}, 1),
		ready:   make(chan struct{}),
	}
}

func (api *HttpApi) run() {
	defer close(api.stop)

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

	close(api.ready)

	go http.ListenAndServe(fmt.Sprintf(":%d", api.apiPort), router)

	<-api.stop
}

func (api *HttpApi) Start() <-chan struct{} {
	go api.run()
	return api.ready
}

func (api *HttpApi) Stop() <-chan struct{} {
	api.stop <- struct{}{}
	return api.stop
}