package routes

import (
	"encoding/json"
	"github.com/gorilla/mux"
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

func UserPlaylistsBasicInfoHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	username := vars["username"]

	playlists, err := models.GetPlaylistsBasicInfoByUsername(username)
	if err != nil {
		// return error page
	}
	json.NewEncoder(w).Encode(playlists)
}

func UserPlaylistSongsHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	username := vars["username"]
	playlistId := vars["playlistId"]

	playlists, err := models.GetPlaylistsSongs(username, playlistId)
	if err != nil {
		// return error page
	}
	json.NewEncoder(w).Encode(playlists)
}