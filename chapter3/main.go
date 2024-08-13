package main

import (
	"errors"
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

func divideHandler(w http.ResponseWriter, req *http.Request) {
	a, b := 10.0, 5.0
	res, err := divide(a, b)
	if err != nil {
		_, _ = fmt.Fprintf(w, "Can not devide %.2f by %.2f", a, b)
		return
	}
	_, _ = fmt.Fprintf(w, "Divide %.2f by %.2f is %.2f", a, b, res)
}

func sum(a, b int) int {
	return a + b
}

func divide(dividend, divider float64) (float64, error) {
	if divider == 0 {
		return 0, errors.New("divide by 0")
	}
	return dividend / divider, nil
}

func main() {
	http.HandleFunc("/", homeHandler)
	http.HandleFunc("/about", aboutHandler)
	http.HandleFunc("/divide", divideHandler)
	fmt.Println("Server is running on port", PORT)
	_ = http.ListenAndServe(PORT, nil)
}
