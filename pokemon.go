package main

import (
	"fmt"
	"net/http"
)

func (app *application) pokemonByIdHandler(w http.ResponseWriter, r *http.Request) {
	id, err := readIdParam(r)
	if err != nil {
		http.NotFound(w, r)
		return
	}

	fmt.Fprintf(w, "show pokemon %d\n", id)
}
