package routes

import (
	"esefexapi/audioplayer"
	"esefexapi/clientnotifiy"
	"esefexapi/config"
	"esefexapi/db"

	"github.com/bwmarrin/discordgo"
)

type RouteHandlers struct {
	dbs    *db.Databases
	a      audioplayer.IAudioPlayer
	ds     *discordgo.Session
	wsCN   *clientnotifiy.WsClientNotifier
	cProto string
}

func NewRouteHandlers(dbs *db.Databases, a audioplayer.IAudioPlayer, ds *discordgo.Session, wsCN *clientnotifiy.WsClientNotifier) *RouteHandlers {
	return &RouteHandlers{
		a:      a,
		dbs:    dbs,
		ds:     ds,
		cProto: config.Get().HttpApi.CustomProtocol,
		wsCN:   wsCN,
	}
}
