package api

import (
	"esefexapi/api/middleware"
	"esefexapi/api/routes"
	"esefexapi/audioplayer"
	"esefexapi/clientnotifiy"
	"esefexapi/config"
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
	port     int
	cProto   string
	domain   string
	stop     chan struct{}
	ready    chan struct{}
}

func NewHttpApi(dbs *db.Databases, plr audioplayer.IAudioPlayer, ds *discordgo.Session, wsCN *clientnotifiy.WsClientNotifier, domain string) *HttpApi {
	return &HttpApi{
		handlers: routes.NewRouteHandlers(dbs, plr, ds, wsCN),
		mw:       middleware.NewMiddleware(dbs, ds),
		a:        plr,
		port:     config.Get().HttpApi.Port,
		cProto:   config.Get().HttpApi.CustomProtocol,
		domain:   domain,
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

	router.Handle("/api/sounds/{guild_id}", cors(h.GetSounds())).Methods("GET")

	router.Handle("/api/guild", cors(auth(h.GetGuild()))).Methods("GET").Headers("Cookie", "")
	router.Handle("/api/guilds", cors(auth(h.GetGuilds()))).Methods("GET").Headers("Cookie", "")

	router.Handle("/api/playsound/{user_id}/{guild_id}/{sound_id}", cors(h.PostPlaySoundInsecure())).Methods("POST")
	router.Handle("/api/playsound/{sound_id}", cors(auth(h.PostPlaySound()))).Methods("POST").Headers("Cookie", "")

	router.Handle("/joinsession/{guild_id}", cors(h.GetJoinSession())).Methods("GET")
	router.Handle("/link", cors(h.GetLinkDefer())).Methods("GET").Queries("t", "{t}")
	router.Handle("/api/link", cors(h.GetLinkRedirect())).Methods("GET").Queries("t", "{t}")

	router.Handle("/dump", cors(h.GetDump()))
	router.Handle("/", cors(h.GetIndex())).Methods("GET")

	router.PathPrefix("/static/").Handler(http.StripPrefix("/static/", http.FileServer(http.Dir("./api/public/"))))

	router.Handle("/api/ws", cors(auth(h.GetWs()))).Methods("GET")

	log.Printf("Webserver started on port %d (%s)\n", api.port, api.domain)

	// nolint:errcheck
	go http.ListenAndServe(fmt.Sprintf("0.0.0.0:%d", api.port), router)

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
