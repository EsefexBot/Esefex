package routes

import (
	"esefexapi/audioplayer"
	"esefexapi/db"
)

type RouteHandlers struct {
	a      audioplayer.IAudioPlayer
	dbs    db.Databases
	cProto string
}

func NewRouteHandlers(dbs db.Databases, a audioplayer.IAudioPlayer, cProto string) *RouteHandlers {
	return &RouteHandlers{
		a:      a,
		dbs:    dbs,
		cProto: cProto,
	}
}
