package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, req *http.Request) {
		n, err := fmt.Fprint(w, "Hello world")
		if err != nil {
			fmt.Println(err)
		}
		fmt.Printf("Success. %d bites sent\n", n)
	})
	_ = http.ListenAndServe(":3011", nil)
}
