package helpers

import (
	"fmt"
	"html/template"
	"net/http"
)

func RenderPage(w http.ResponseWriter, pageName string) {
	templatePath := "templates/" + pageName + ".page.tmpl"
	_template, err := template.ParseFiles(templatePath)
	if err != nil {
		fmt.Println(err)
	}
	err = _template.Execute(w, nil)
	if err != nil {
		fmt.Println(err)
	}
}
