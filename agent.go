package main

import (
	"fmt"
	"net/http"
)

func agent(w http.ResponseWriter, r *http.Request) {

	UserAgent := r.Header.Get("User-Agent")
	if UserAgent == "" {
		http.Error(w, "error", http.StatusBadRequest)
	}

	fmt.Fprint(w, "You are visiting us using:", UserAgent)
}
