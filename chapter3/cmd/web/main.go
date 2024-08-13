package main

import (
	"fmt"
	"github.com/lapeko/udemy__building-modern-web-applications-with-go/tree/main/chapter3/pkg/handlers"
	"net/http"
)

const PORT = ":3011"

func main() {
	http.HandleFunc("/", handlers.HomeHandler)
	http.HandleFunc("/about", handlers.AboutHandler)
	fmt.Println("Server is running on port", PORT)
	_ = http.ListenAndServe(PORT, nil)
}
