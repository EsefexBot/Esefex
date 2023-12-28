package api

import (
	"esefexapi/api/routes"
	"esefexapi/audioplayer"
	"esefexapi/db"
	"esefexapi/service"

	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

var _ service.IService = &HttpApi{}

// HttpApi implements Service
type HttpApi struct {
	handlers *routes.RouteHandlers
	a        audioplayer.IAudioPlayer
	apiPort  int
	cProto   string
	stop     chan struct{}
	ready    chan struct{}
}

func NewHttpApi(dbs db.Databases, plr audioplayer.IAudioPlayer, apiPort int, cProto string) *HttpApi {
	return &HttpApi{
		handlers: routes.NewRouteHandlers(dbs, plr, cProto),
		a:        plr,
		apiPort:  apiPort,
		cProto:   cProto,
		stop:     make(chan struct{}, 1),
		ready:    make(chan struct{}),
	}
}

func (api *HttpApi) run() {
	defer close(api.stop)
	log.Println("Starting webserver...")
	defer log.Println("Webserver stopped")

	router := mux.NewRouter()

	router.HandleFunc("/api/sounds/{server_id}", api.handlers.GetSounds).Methods("GET")
	router.HandleFunc("/api/playsound/{user_id}/{server_id}/{sound_id}", api.handlers.PostPlaySoundInsecure).Methods("POST")
	router.HandleFunc("/api/playsound/{sound_id}", api.handlers.PostPlaySound).Methods("POST").Headers("User-Token", "")

	router.HandleFunc("/joinsession/{server_id}", api.handlers.GetJoinSession).Methods("GET")
	router.HandleFunc("/link", api.handlers.GetLink).Methods("GET").Queries("t", "{t}")

	router.HandleFunc("/dump", api.handlers.GetDump)
	router.HandleFunc("/", api.handlers.GetIndex).Methods("GET")

	// http.Handle("/", router)
	log.Printf("Webserver started on port %d (http://localhost:%d)\n", api.apiPort, api.apiPort)

	go http.ListenAndServe(fmt.Sprintf(":%d", api.apiPort), router)

	close(api.ready)
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
