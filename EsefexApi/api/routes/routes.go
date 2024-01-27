package routes

import (
	"esefexapi/audioplayer"
	"esefexapi/clientnotifiy"
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

func NewRouteHandlers(dbs *db.Databases, a audioplayer.IAudioPlayer, ds *discordgo.Session, cProto string, wsCN *clientnotifiy.WsClientNotifier) *RouteHandlers {
	return &RouteHandlers{
		a:      a,
		dbs:    dbs,
		ds:     ds,
		cProto: cProto,
		wsCN:   wsCN,
	}
}
