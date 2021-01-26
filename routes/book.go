package routes

import (
	"html/template"
	"net/http"
)

type BookIndexData struct {
	Title string
	Message string
}

func BookIndexHandler(w http.ResponseWriter, r *http.Request) {
	/*
		vars := mux.Vars(r)
		title := vars["title"]
		page := vars["page"] */

	t := template.Must(template.ParseFiles("templates/base.html", "templates/books/index.html"))
	t.Execute(w, BookIndexData{
		Title: "Book Index",
		Message: "Book index wow",
	})
}