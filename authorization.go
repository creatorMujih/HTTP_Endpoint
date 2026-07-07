package main 


import (
	"fmt"
	"net/http"
)

func authorizationHandler(w http.ResponseWriter, r *http.Request) {
	Apikey := r.Header.Get("X-API-Key")

	if Apikey != "secret123" {
		http.Error(w, "unauthorized: Invalid or missing API key", http.StatusUnauthorized)
		return
	}
	fmt.Fprintln(w, "Welcome to the secure dashboard!")

}