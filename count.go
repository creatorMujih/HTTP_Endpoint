package main

import (
	"fmt"
	"io"
	"net/http"
)

func count(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		fmt.Fprint(w, "Send a POST request with text to count words")
	}

	if r.Method == http.MethodPost {
		Body, err := io.ReadAll(r.Body)
		if err != nil {
			http.Error(w, "Failed to read request body", http.StatusInternalServerError)
			return
		}

		fmt.Fprint(w, len([]rune(string(Body))))
		return
	}
	http.Error(w, "Method Not Found", http.StatusInternalServerError)
}
