package main

import (
	"net/http"

	"go_rest_pokedex.stewie.in/data"
)

func (app *application) pokemonByIdHandler(w http.ResponseWriter, r *http.Request) {
	id, err := readIdParam(r)
	if err != nil {
		http.NotFound(w, r)
		return
	}

	pokemon := data.Pokemon{
		Rank:      id,
		Name:      "Bulbasaur",
		Type_1:    "Grass",
		Type_2:    "Poison",
		Legendary: false,
	}

	err = writeJson(w, http.StatusOK, envelope{"pokemon": pokemon}, nil)
	if err != nil {
		app.logger.Error(err.Error())
		http.Error(w, "The server encountered a problem and could not process your request", http.StatusInternalServerError)
	}
}
