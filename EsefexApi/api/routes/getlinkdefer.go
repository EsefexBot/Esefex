package routes

import (
	"html/template"
	"net/http"
)

type LinkDeferData struct {
	LinkToken string
}

// link?<guild_id>
func (h *RouteHandlers) GetLinkDefer(w http.ResponseWriter, r *http.Request) {
	linkToken := r.URL.Query().Get("t")

	if linkToken == "" {
		http.Error(w, "Missing link token", http.StatusBadRequest)
		return
	}

	data := LinkDeferData{
		LinkToken: linkToken,
	}

	tmpl, err := template.ParseFiles("./api/templates/link.html")
	if err != nil {
		http.Error(w, "Error parsing template", http.StatusInternalServerError)
		return
	}

	err = tmpl.Execute(w, data)
	if err != nil {
		http.Error(w, "Error executing template", http.StatusInternalServerError)
		return
	}
}
