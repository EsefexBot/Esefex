package routes

import (
	"esefexapi/audioplayer"
	"esefexapi/db"

	"github.com/bwmarrin/discordgo"
)

type RouteHandlers struct {
	dbs    *db.Databases
	a      audioplayer.IAudioPlayer
	ds     *discordgo.Session
	cProto string
}

func NewRouteHandlers(dbs *db.Databases, a audioplayer.IAudioPlayer, ds *discordgo.Session, cProto string) *RouteHandlers {
	return &RouteHandlers{
		a:      a,
		dbs:    dbs,
		ds:     ds,
		cProto: cProto,
	}
}
