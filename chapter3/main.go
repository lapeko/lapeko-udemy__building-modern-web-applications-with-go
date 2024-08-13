package main

import (
	"fmt"
	"html/template"
	"net/http"
)

const PORT = ":3011"

func homeHandler(w http.ResponseWriter, req *http.Request) {
	renderPage(w, "home")
}

func aboutHandler(w http.ResponseWriter, req *http.Request) {
	renderPage(w, "about")
}

func main() {
	http.HandleFunc("/", homeHandler)
	http.HandleFunc("/about", aboutHandler)
	fmt.Println("Server is running on port", PORT)
	_ = http.ListenAndServe(PORT, nil)
}

func renderPage(w http.ResponseWriter, pageName string) {
	templatePath := "./templates/" + pageName + ".html"
	_template, err := template.ParseFiles(templatePath)
	if err != nil {
		fmt.Println(err)
	}
	err = _template.Execute(w, nil)
	if err != nil {
		fmt.Println(err)
	}
}
