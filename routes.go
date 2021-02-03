package main

import (
	"github.com/gorilla/mux"
	"web/main/routes"
)


func setUpRoutes() *mux.Router {
	r := mux.NewRouter()

	r.HandleFunc("/greatest_hits/{period}", routes.GreatestHitsHandler)
	r.HandleFunc("/playlists", routes.PlaylistsHandler)
	r.HandleFunc("/playlists/{username}", routes.UserPlaylistsBasicInfoHandler)
	r.HandleFunc("/playlists/{username}/{playlistId}", routes.UserPlaylistSongsHandler)
	r.HandleFunc("/", routes.IndexHandler)

	return r
}