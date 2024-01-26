package middleware

import (
	"context"
	"esefexapi/userdb"
	"fmt"
	"log"
	"net/http"
)

// Auth middleware checks if the user is authenticated and injects the user into the request context
func (m *Middleware) Auth(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		user_token, err := r.Cookie("User-Token")
		if err != nil {
			errorMsg := fmt.Sprintf("Error getting user token cookie: %+v", err)

			log.Println(errorMsg)
			http.Error(w, errorMsg, http.StatusUnauthorized)
			return
		}

		Ouser, err := m.dbs.UserDB.GetUserByToken(userdb.Token(user_token.Value))
		if err != nil || Ouser.IsNone() {
			errorMsg := fmt.Sprintf("Error getting user by token: %+v", err)

			log.Println(errorMsg)
			http.Error(w, errorMsg, http.StatusUnauthorized)
			return
		}

		// Inject the user into the request context
		ctx := context.WithValue(r.Context(), "user", Ouser)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
