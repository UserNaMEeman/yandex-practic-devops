package main

import (
	"net/http"

	"github.com/UserNaMEeman/yandex-practic-devops/cmd/server/handler"
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
)

func main() {
	r := chi.NewRouter()
	r.Use(middleware.RequestID)
	r.Use(middleware.RealIP)
	r.Use(middleware.Recoverer)
	r.Use(middleware.Logger)

	r.Get("/", handler.ShowAllMetrics)
	r.Get("/value/{type}/{name}", handler.ShowMetrics)
	r.Route("/update", func(r chi.Router) {
		r.Post("/{type}/{name}/{value}", handler.HandleMetric)
	})
	http.ListenAndServe(":8080", r)
}
