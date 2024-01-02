package api

import (
	"esefexapi/api/middleware"
	"esefexapi/api/routes"
	"esefexapi/audioplayer"
	"esefexapi/db"
	"esefexapi/service"

	"fmt"
	"log"
	"net/http"

	"github.com/bwmarrin/discordgo"
	"github.com/gorilla/mux"
)

var _ service.IService = &HttpApi{}

// HttpApi implements Service
type HttpApi struct {
	handlers *routes.RouteHandlers
	mw       *middleware.Middleware
	a        audioplayer.IAudioPlayer
	apiPort  int
	cProto   string
	stop     chan struct{}
	ready    chan struct{}
}

func NewHttpApi(dbs *db.Databases, plr audioplayer.IAudioPlayer, ds *discordgo.Session, apiPort int, cProto string) *HttpApi {
	return &HttpApi{
		handlers: routes.NewRouteHandlers(dbs, plr, ds, cProto),
		mw:       middleware.NewMiddleware(dbs),
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
	auth := api.mw.Auth
	cors := api.mw.Cors
	h := api.handlers

	router.HandleFunc("/api/sounds/{server_id}", cors(h.GetSounds)).Methods("GET")

	router.HandleFunc("/api/server", cors(auth(h.GetServer))).Methods("GET").Headers("User-Token", "")
	router.HandleFunc("/api/servers", cors(auth(h.GetServers))).Methods("GET").Headers("User-Token", "")

	router.HandleFunc("/api/playsound/{user_id}/{server_id}/{sound_id}", cors(h.PostPlaySoundInsecure)).Methods("POST")
	router.HandleFunc("/api/playsound/{sound_id}", cors(auth(h.PostPlaySound))).Methods("POST").Headers("User-Token", "")

	router.HandleFunc("/joinsession/{server_id}", cors(h.GetJoinSession)).Methods("GET")
	router.HandleFunc("/link", cors(h.GetLink)).Methods("GET").Queries("t", "{t}")

	router.HandleFunc("/dump", cors(h.GetDump))
	router.HandleFunc("/", cors(h.GetIndex)).Methods("GET")

	router.PathPrefix("/static/").Handler(http.StripPrefix("/static/", http.FileServer(http.Dir("./api/public/"))))

	log.Printf("Webserver started on port %d (http://localhost:%d)\n", api.apiPort, api.apiPort)

	go http.ListenAndServe(fmt.Sprintf("0.0.0.0:%d", api.apiPort), router)

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
