package main

import (
	"chitChat/data"
	"net/http"
)

func index(w http.ResponseWriter, r *http.Request) {

	threads, err := data.Threads()
	if err == nil {
		_, err := session(w, r)

		privateFiles := []string{
			"private_navbar.html",
			"index.html",
			"layout.html",
		}

		publicFiles := []string{
			"public_navbar.html",
			"index.html",
			"layout.html",
		}

		if err != nil {
			generateHtml(w, threads, publicFiles...)
		} else {
			generateHtml(w, threads, privateFiles...)
		}
	}
}
