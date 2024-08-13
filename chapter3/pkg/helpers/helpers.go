package helpers

import (
	"fmt"
	"html/template"
	"net/http"
)

func RenderPage(w http.ResponseWriter, pageName string) {
	templatePath := "templates/" + pageName + ".page.gohtml"
	layoutPath := "templates/base.layout.gohtml"
	_template, err := template.ParseFiles(templatePath, layoutPath)
	if err != nil {
		fmt.Println(err)
	}
	err = _template.Execute(w, nil)
	if err != nil {
		fmt.Println(err)
	}
}
