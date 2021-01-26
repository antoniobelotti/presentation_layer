package routes

import (
	"html/template"
	"net/http"
)

type GreatestHitsData struct {
	Title string
	Message string
}

func GreatestHitsHandler(w http.ResponseWriter, r *http.Request) {
	t := template.Must(template.ParseFiles("templates/base.html", "templates/greatest_hits/index.html"))
	t.Execute(w, GreatestHitsData{
		Title: "",
		Message: "",
	})
}