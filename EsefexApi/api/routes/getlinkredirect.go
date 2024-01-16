package routes

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
)

type LinkRedirect struct {
	RedirectUrl string
}

// api/link?<guild_id>
func (h *RouteHandlers) GetLinkRedirect() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		linkToken := r.URL.Query().Get("t")

		if linkToken == "" {
			http.Error(w, "Missing link token", http.StatusBadRequest)
			return
		}

		isValid, err := h.dbs.LinkTokenStore.ValidateToken(linkToken)
		if err != nil || !isValid {
			http.Error(w, fmt.Sprintf("Invalid link token, the token does not exist or is expired:\n%s", err), http.StatusBadRequest)
			return
		}

		userID, err := h.dbs.LinkTokenStore.GetUser(linkToken)
		if err != nil {
			http.Error(w, fmt.Sprintf("Invalid link token, the token does not exist or is expired:\n%s", err), http.StatusBadRequest)
			return
		}

		userToken, err := h.dbs.UserDB.NewToken(userID)
		if err != nil {
			http.Error(w, fmt.Sprintf("Error creating new user token:\n%s", err), http.StatusInternalServerError)
			return
		}

		cookie := http.Cookie{
			Name:   "User-Token",
			Value:  string(userToken),
			Path:   "/",
			MaxAge: 0,
			// enable this once we have https
			Secure:   false,
			HttpOnly: true,
			SameSite: http.SameSiteDefaultMode,
		}
		http.SetCookie(w, &cookie)

		w.Header().Set("Content-Type", "text/html")

		tmpl, err := template.ParseFiles("./api/templates/linkredirect.html")
		if err != nil {
			http.Error(w, "Error parsing template", http.StatusInternalServerError)
			return
		}

		err = tmpl.Execute(w, LinkRedirect{
			RedirectUrl: fmt.Sprintf("%s://link/%s", h.cProto, userToken),
		})
		if err != nil {
			http.Error(w, "Error executing template", http.StatusInternalServerError)
			return
		}

		err = h.dbs.LinkTokenStore.DeleteToken(userID)
		if err != nil {
			log.Printf("Error deleting link token: %s", err)
			http.Error(w, fmt.Sprintf("Error deleting link token: %s", err), http.StatusInternalServerError)
		}

		log.Printf("got /joinsession request\n")
	})
}
