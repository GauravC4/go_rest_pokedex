package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func greetHandler(w http.ResponseWriter, r *http.Request) {
	type Pokemon struct {
		Name string
		Type string
	}

	poki := Pokemon{Name: "Pikachu", Type: "Electric"}

	pokiJson, err := json.Marshal(poki)

	if err != nil {
		fmt.Println("Json mashling error", err)
	}

	w.Header().Set("Content-Type", "application/json")

	fmt.Fprint(w, string(pokiJson))

}
