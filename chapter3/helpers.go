package main

import (
	"fmt"
	"html/template"
	"net/http"
)

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
