package middleware

import "esefexapi/db"

type Middleware struct {
	dbs *db.Databases
}

func NewMiddleware(dbs *db.Databases) *Middleware {
	return &Middleware{
		dbs: dbs,
	}
}
