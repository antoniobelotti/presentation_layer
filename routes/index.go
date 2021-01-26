package routes

import (
	"html/template"
	"net/http"
)

type IndexData struct {
	Title string
	Message string
}

func IndexHandler(w http.ResponseWriter, r *http.Request) {
	/*
		vars := mux.Vars(r)
		title := vars["title"]
		page := vars["page"] */

	t := template.Must(template.ParseFiles("templates/base.html"))
	t.Execute(w, IndexData{
		Title: "example - Index",
		Message: "Hello World",
	})
}