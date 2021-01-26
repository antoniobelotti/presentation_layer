package routes

import (
	"html/template"
	"net/http"
)

func PlaylistsHandler(w http.ResponseWriter, r *http.Request) {
	t := template.Must(template.ParseFiles("templates/base.html", "templates/playlists/index.html"))
	t.Execute(w, nil)
}