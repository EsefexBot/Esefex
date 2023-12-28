package routes

import (
	"fmt"
	"log"
	"net/http"
)

// link?<server_id>
func (h *RouteHandlers) GetLink(w http.ResponseWriter, r *http.Request) {
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

	redirectUrl := fmt.Sprintf("%s://link/%s", h.cProto, userToken)
	response := fmt.Sprintf(`<meta http-equiv="refresh" content="0; URL=%s" />`, redirectUrl)
	fmt.Fprint(w, response)

	cookie := http.Cookie{
		Name:     "User-Token",
		Value:    string(userToken),
		Path:     "/",
		MaxAge:   0,
		Secure:   true,
		HttpOnly: true,
		SameSite: http.SameSiteStrictMode,
	}
	http.SetCookie(w, &cookie)

	w.Header().Set("Content-Type", "text/html")
	w.Header().Set("Access-Control-Allow-Origin", "*")

	h.dbs.LinkTokenStore.DeleteToken(userID)

	log.Printf("got /joinsession request\n")
}
