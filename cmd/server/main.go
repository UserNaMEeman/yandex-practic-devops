package main

import (
	"net/http"

	"github.com/UserNaMEeman/yandex-practic-devops/cmd/server/handler"
	"github.com/UserNaMEeman/yandex-practic-devops/cmd/server/storage"
	"github.com/go-chi/chi"
)

func main() {
	var recMetric storage.Metrics
	// var JSONMetrics storage.Metrics
	pullMetrics := make(map[string]storage.Metrics)
	// pullJSONMetrics := make(map[string]storage.Metrics)
	r := chi.NewRouter()
	// r.Use(middleware.RequestID)
	// r.Use(middleware.RealIP)
	// r.Use(middleware.Recoverer)
	// r.Use(middleware.Logger)

	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		handler.ShowAllMetrics(w, pullMetrics)
	})
	r.Get("/value/{type}/{name}", func(w http.ResponseWriter, r *http.Request) {
		handler.ShowOneMetric(w, r, pullMetrics)
	})
	// r.Get("/value/{type}/{name}", handler.ShowMetrics)
	r.Route("/update", func(r chi.Router) {
		r.Post("/{type}/{name}/{value}", func(w http.ResponseWriter, r *http.Request) {
			recMetric, _ = handler.HandleMetric(w, r, pullMetrics)
			pullMetrics[recMetric.ID] = recMetric
		})
		r.Post("/", func(w http.ResponseWriter, r *http.Request) {
			recMetric = handler.HandleJSONMetric(w, r, pullMetrics)
			pullMetrics[recMetric.ID] = recMetric
			// fmt.Println(JSONMetrics)
		})
	})
	// r.Route("/value", func(r chi.Router) {
	r.Post("/value", func(w http.ResponseWriter, r *http.Request) {
		handler.ShowJSONMetrics(w, r, pullMetrics)
	})
	// }
	http.ListenAndServe(":8080", r)
}
