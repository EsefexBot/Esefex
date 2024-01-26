package middleware

import (
	"esefexapi/db"

	"github.com/bwmarrin/discordgo"
)

type Middleware struct {
	dbs *db.Databases
	ds  *discordgo.Session
}

func NewMiddleware(dbs *db.Databases, ds *discordgo.Session) *Middleware {
	return &Middleware{
		dbs: dbs,
		ds:  ds,
	}
}
