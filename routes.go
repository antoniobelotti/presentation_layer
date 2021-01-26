package main

import (
	"github.com/gorilla/mux"
	"web/main/routes"
)


func setUpRoutes() *mux.Router {
	r := mux.NewRouter()

	r.HandleFunc("/books", routes.BookIndexHandler)
	r.HandleFunc("/", routes.IndexHandler)

	return r
}