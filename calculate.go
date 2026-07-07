package main

import (
	"fmt"
	"net/http"
	"strconv"
)

func calculate(w http.ResponseWriter, r *http.Request) {

	op := r.URL.Query().Get("op")
	aStr := r.URL.Query().Get("a")
	bStr := r.URL.Query().Get("b")

	if op == "" || aStr == "" || bStr == "" {
		http.Error(w, "Empty Query", http.StatusBadRequest)
		return
	}
	a, err := strconv.Atoi(aStr)
	if err != nil {
		http.Error(w, "Error With Conversion", http.StatusBadRequest)
		return
	}

	b, err := strconv.Atoi(bStr)
	if err != nil {
		http.Error(w, "Error With Conversion", http.StatusBadRequest)
		return
	}

	var result int
	switch op {
	case "add":
		result = a + b
	case "subtract":
		result = a + b
	case "multiply":
		result = a * b
	default:
		http.Error(w, "400 Bad Request", http.StatusBadRequest)
	}
	fmt.Fprintf(w, "Result: %d", result)
}
