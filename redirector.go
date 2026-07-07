package main

import (
	"net/http"
	"fmt"
)

func redirectorOld(w http.ResponseWriter, r *http.Request) {
	http.Redirect(w, r, "/v2", http.StatusMovedPermanently)
}

func redirectorNew(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Welcome to version 2")
}
