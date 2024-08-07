package main

import "net/http"

func (app *application) routes() http.Handler {
	mux := http.NewServeMux()
	mux.HandleFunc("GET /healthcheck", app.healthcheckHandler)
	mux.HandleFunc("GET /pokemon/{id}", app.pokemonByIdHandler)

	return mux
}
