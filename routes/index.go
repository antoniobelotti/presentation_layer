package routes

import (
	"encoding/json"
	"fmt"
	"html/template"
	"net/http"
	"web/main/models"
)


func IndexHandler(w http.ResponseWriter, r *http.Request) {
	stats,err := models.GetGeneralStats()
	if err!= nil{
		fmt.Println(err)
	}

	t := template.Must(template.ParseFiles("templates/base.html", "templates/index.html"))
	t.Execute(w, stats)
}

func StatsPlaylistsLengthDistribution(w http.ResponseWriter, r *http.Request) {
	distribution,err := models.GetPlaylistLengthDistribution()
	if err!= nil{
		fmt.Println(err)
	}

	json.NewEncoder(w).Encode(distribution)
}