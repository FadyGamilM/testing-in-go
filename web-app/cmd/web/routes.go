package main

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func (app *application) routes() http.Handler {
	// define a something that implements the http Handler
	mux := chi.NewRouter()

	// register middlewares
	mux.Use(middleware.Recoverer)

	// register routes
	mux.Get("/home", app.Home)

	// return the handler
	return mux
}
