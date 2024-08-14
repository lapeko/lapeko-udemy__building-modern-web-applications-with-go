package main

import (
	"github.com/lapeko/udemy__building-modern-web-applications-with-go/tree/main/chapter3/pkg/handlers"
	"log"
	"net/http"
)

const PORT = ":3011"

// TODO implement template caching
// template.New + GlobalReact from folder like "/template/*.layout.gohtml

func main() {
	http.HandleFunc("/", handlers.HomeHandler)
	http.HandleFunc("/about", handlers.AboutHandler)
	log.Println("Server is running on port", PORT)
	_ = http.ListenAndServe(PORT, nil)
}
