package main

import (
	"log"
	"net/http"
)

func main() {
	mux := http.NewServeMux()

	files := http.FileServer(http.Dir("/public"))
	mux.Handle("/static/", http.StripPrefix("/static/", files))
	mux.HandleFunc("/", index)
	mux.HandleFunc("/authenticate", authenticate)

	server := &http.Server{
		Addr:    "0.0.0.0:8080",
		Handler: mux,
	}
	err := server.ListenAndServe()
	if err != nil {
		log.Fatal(err)
	}
}
