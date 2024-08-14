package main

import (
	"github.com/lapeko/udemy__building-modern-web-applications-with-go/tree/main/chapter3/pkg/routes"
	"log"
	"net/http"
)

const PORT = ":3011"

func main() {
	m := routes.InitializeRoutes()
	http.Handle("/", m)
	log.Println("Server is running on port", PORT)
	_ = http.ListenAndServe(PORT, nil)
}
