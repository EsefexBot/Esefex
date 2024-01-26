package middleware

import "esefexapi/db"

type CommandMiddleware struct {
	dbs *db.Databases
}

func NewCommandMiddleware(dbs *db.Databases) *CommandMiddleware {
	return &CommandMiddleware{
		dbs: dbs,
	}
}
