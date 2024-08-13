package main

import "net/http"

func homeHandler(w http.ResponseWriter, req *http.Request) {
	renderPage(w, "home")
}

func aboutHandler(w http.ResponseWriter, req *http.Request) {
	renderPage(w, "about")
}
