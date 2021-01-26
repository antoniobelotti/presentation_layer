package routes

import (
	"html/template"
	"net/http"
)

type PlaylistsData struct {
	Title string
	Message string
}

func PlaylistsHandler(w http.ResponseWriter, r *http.Request) {
	t := template.Must(template.ParseFiles("templates/base.html", "templates/playlists/index.html"))
	t.Execute(w, GreatestHitsData{
		Title: "",
		Message: "",
	})
}