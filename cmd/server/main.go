package main

import (
	"net/http"

	"github.com/UserNaMEeman/yandex-practic-devops/cmd/server/handler"
)

func main() {
	// r := chi.NewRouter()
	// r.Post("/update/", handler.HandleMetric)
	http.HandleFunc("/update/", handler.HandleMetric)
	http.ListenAndServe(":8080", nil)
}
