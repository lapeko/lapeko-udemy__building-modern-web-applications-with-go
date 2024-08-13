package main

import (
	"fmt"
	"net/http"
)

const PORT = ":3011"

func aboutHandler(w http.ResponseWriter, req *http.Request) {
	a, b := 1, 2
	sumRes := sum(a, b)
	_, _ = fmt.Fprintf(w, "Sum of %d and %d is %d", a, b, sumRes)
}

func homeHandler(w http.ResponseWriter, req *http.Request) {
	_, _ = fmt.Fprintf(w, "Home page")
}

func sum(a, b int) int {
	return a + b
}

func main() {
	http.HandleFunc("/", homeHandler)
	http.HandleFunc("/about", aboutHandler)
	fmt.Println("Server is running on port", PORT)
	_ = http.ListenAndServe(PORT, nil)
}
