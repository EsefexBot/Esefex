package middleware

import (
	"esefexapi/userdb"
	"fmt"
	"log"
	"net/http"
)

// Auth middleware checks if the user is authenticated and injects the user into the request context
func (m *Middleware) Auth(next func(w http.ResponseWriter, r *http.Request, userID string)) func(w http.ResponseWriter, r *http.Request) {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		user_token := r.Header.Get("User-Token")

		Ouser, err := m.dbs.UserDB.GetUserByToken(userdb.Token(user_token))
		if err != nil || Ouser.IsNone() {
			errorMsg := fmt.Sprintf("Error getting user by token: %+v", err)

			log.Println(errorMsg)
			http.Error(w, errorMsg, http.StatusUnauthorized)
			return
		}

		next(w, r, Ouser.Unwrap().ID)
	})
}
