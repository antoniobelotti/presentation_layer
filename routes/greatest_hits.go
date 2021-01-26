package routes

import (
	"html/template"
	"net/http"
	"web/main/models"
)


func GreatestHitsHandler(w http.ResponseWriter, r *http.Request) {
	rows, err := models.TodayGreatestHits()
	if err != nil {
		// return error page
	}

	funcMap := template.FuncMap{
		// The name "inc" is what the function will be called in the template text.
		"inc": func(i int) int {
			return i + 1
		},
	}
	t := template.Must(template.New("base.html").Funcs(funcMap).ParseFiles("templates/base.html", "templates/greatest_hits/index.html"))
	t.Execute(w, rows)
}
