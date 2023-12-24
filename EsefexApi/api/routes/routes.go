package routes

import (
	"esefexapi/audioplayer"
	"esefexapi/sounddb"
)

type RouteHandlers struct {
	a      audioplayer.IAudioPlayer
	db     sounddb.ISoundDB
	cProto string
}

func NewRouteHandler(db sounddb.ISoundDB, a audioplayer.IAudioPlayer, cProto string) *RouteHandlers {
	return &RouteHandlers{
		a:      a,
		db:     db,
		cProto: cProto,
	}
}
