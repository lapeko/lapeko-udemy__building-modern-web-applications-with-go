package routes

import (
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/lapeko/udemy__building-modern-web-applications-with-go/tree/main/chapter3/pkg/handlers"
	"net/http"
)

func InitializeRoutes() *chi.Mux {
	r := chi.NewRouter()

	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)

	r.Get("/", http.HandlerFunc(handlers.HomeHandler))
	r.Get("/about", http.HandlerFunc(handlers.AboutHandler))

	return r
}
