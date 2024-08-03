package main

import (
	"net/http"
)

func main() {
	http.HandleFunc("/greet", greetHandler)

	http.ListenAndServe(":8080", nil)
}
