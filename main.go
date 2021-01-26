package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"web/main/models"
)


func main() {
	port, success := os.LookupEnv("EXPOSE_PORT")
	if ! success{
		log.Println("Unable to read EXPOSE_PORT environment variable. Defaulting to port 8080")
		port = "8080"
	}

	err := models.InitDB()
	if err != nil{
		log.Fatal("Unable to establish connection to database")
	}

	router := setUpRoutes()
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%s",port), router))
}
