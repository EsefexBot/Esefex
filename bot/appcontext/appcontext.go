package appcontext

import (
	"net/http"
)

type Context struct {
	CustomProtocol string
}

func Wrap(fn func(http.ResponseWriter, *http.Request, Context), c Context) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		fn(w, r, c)
	}
}
