package main

import (
	"fmt"
	"net/http"
)

func pongPin(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "pong")
}
