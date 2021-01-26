package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
)


func main() {
	port, success := os.LookupEnv("EXPOSE_PORT")
	if ! success{
		log.Println("Unable to read EXPOSE_PORT environment variable. Defaulting to port 8080")
		port = "8080"
	}

	router := setUpRoutes()
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%s",port), router))
}
