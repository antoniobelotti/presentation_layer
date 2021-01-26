package routes

import (
	"github.com/gorilla/mux"
	"html/template"
	"net/http"
	"web/main/models"
)


func GreatestHitsHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	period := vars["period"]

	rows, err := models.GreatestHits(period)
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
