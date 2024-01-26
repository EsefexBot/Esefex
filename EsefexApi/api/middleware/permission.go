package middleware

import (
	"esefexapi/permissions"
	"esefexapi/types"
	"esefexapi/userdb"
	"esefexapi/util/dcgoutil"
	"esefexapi/util/refl"
	"net/http"
)

func (m *Middleware) Permission(next http.Handler, perms ...string) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Get the user from the request context
		user := r.Context().Value("user").(userdb.User)
		userChan, err := dcgoutil.UserVCAny(m.ds, user.ID)
		if err != nil {
			errorMsg := "Error getting user channel: " + err.Error()
			http.Error(w, errorMsg, http.StatusInternalServerError)
			return
		}

		if userChan.IsNone() {
			errorMsg := "User is not in a voice channel"
			http.Error(w, errorMsg, http.StatusUnauthorized)
			return
		}

		p, err := m.dbs.PermissionDB.Query(types.GuildID(userChan.Unwrap().ChannelID), user.ID)
		if err != nil {
			errorMsg := "Error querying permissions: " + err.Error()
			http.Error(w, errorMsg, http.StatusInternalServerError)
			return
		}

		// Check if the user has any of the required permissions
		for _, perm := range perms {
			ps, err := refl.GetNestedFieldValue(p, perm)
			if err != nil {
				errorMsg := "Error getting nested field value: " + err.Error()
				http.Error(w, errorMsg, http.StatusInternalServerError)
				return
			}

			if !ps.(permissions.PermissionState).Allowed() {
				errorMsg := "User does not have the required permissions"
				http.Error(w, errorMsg, http.StatusUnauthorized)
				return
			}
		}

		next.ServeHTTP(w, r)
	})
}
