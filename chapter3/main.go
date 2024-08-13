package main

import (
	"fmt"
	"net/http"
)

const PORT = ":3011"

func main() {
	http.HandleFunc("/", homeHandler)
	http.HandleFunc("/about", aboutHandler)
	fmt.Println("Server is running on port", PORT)
	_ = http.ListenAndServe(PORT, nil)
}
