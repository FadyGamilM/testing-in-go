package main

import (
	"log"
	"net/http"
)

type application struct{}

func main() {
	// setup a config for our app
	app := application{}

	// setup application routes
	mux := app.routes()

	// start the server
	log.Println("starting server on port 8080")
	err := http.ListenAndServe(":8080", mux)
	if err != nil {
		log.Fatal(err)
	}
}
