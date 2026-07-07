package main

import (
	"net/http"
	"log"
)

func main() {
	http.HandleFunc("/ping", pongPin)
	http.HandleFunc("/hello", query)
	http.HandleFunc("/count", count)
	http.HandleFunc("/agent", agent)
	http.HandleFunc("/calculate", calculate)
	http.HandleFunc("/dashboard", authorizationHandler)
	http.HandleFunc("/legacy",redirectorOld)
	http.HandleFunc("/v2",redirectorNew)

	log.Fatal(http.ListenAndServe(":8080", nil))
}
