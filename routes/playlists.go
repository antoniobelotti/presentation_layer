package routes

import (
	"html/template"
	"net/http"
	"web/main/models"
)

func PlaylistsHandler(w http.ResponseWriter, r *http.Request) {
	usernames, err := models.GetAllUsernames()
	if err != nil {
		// return error page
	}
	t := template.Must(template.ParseFiles("templates/base.html", "templates/playlists/index.html"))
	t.Execute(w, usernames)
}