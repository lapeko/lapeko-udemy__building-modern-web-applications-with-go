package helpers

import (
	"html/template"
	"log"
	"net/http"
	"path/filepath"
)

var cachedTemplates = map[string]*template.Template{}

func RenderPage(w http.ResponseWriter, pageName string) {
	_template := retrieveTemplateInCache(pageName)
	err := _template.Execute(w, nil)
	if err != nil {
		log.Fatalln(err)
	}
}

func retrieveTemplateInCache(pageName string) *template.Template {
	cache, ok := cachedTemplates[pageName]
	if ok {
		log.Println("Getting", pageName, "from cache")
		return cache
	}
	templatePath := filepath.Join("templates", pageName+".page.gohtml")
	layouts, err := filepath.Glob(filepath.Join("templates", "*.layout.gohtml"))

	if err != nil {
		panic(err)
	}
	if len(layouts) == 0 {
		log.Fatalln("No layouts found")
	}
	_template, err := template.ParseFiles(append([]string{templatePath}, layouts...)...)

	if err != nil {
		panic(err)
	}
	log.Println("Caching", pageName, "page template")
	cachedTemplates[pageName] = _template
	return _template
}
