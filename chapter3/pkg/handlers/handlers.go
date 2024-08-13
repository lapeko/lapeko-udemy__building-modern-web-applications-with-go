package handlers

import (
	"github.com/lapeko/udemy__building-modern-web-applications-with-go/tree/main/chapter3/pkg/helpers"
	"net/http"
)

func HomeHandler(w http.ResponseWriter, req *http.Request) {
	helpers.RenderPage(w, "home")
}

func AboutHandler(w http.ResponseWriter, req *http.Request) {
	helpers.RenderPage(w, "about")
}
