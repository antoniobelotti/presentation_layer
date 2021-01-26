package main

import (
	"github.com/gorilla/mux"
	"web/main/routes"
)


func setUpRoutes() *mux.Router {
	r := mux.NewRouter()

	r.HandleFunc("/greatest_hits", routes.GreatestHitsHandler)
	r.HandleFunc("/", routes.IndexHandler)

	return r
}