package routes

import (
	"github.com/bmizerany/pat"
	"github.com/lapeko/udemy__building-modern-web-applications-with-go/tree/main/chapter3/pkg/handlers"
	"net/http"
)

func InitializeRoutes() *pat.PatternServeMux {
	m := pat.New()

	m.Get("/", http.HandlerFunc(handlers.HomeHandler))
	m.Get("/about", http.HandlerFunc(handlers.AboutHandler))

	return m
}
