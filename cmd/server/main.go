package main

import (
	"net/http"

	"github.com/UserNaMEeman/yandex-practic-devops/cmd/server/handler"
)

func main() {
	http.HandleFunc("/update/", handler.HandleMetric)
	http.ListenAndServe(":8080", nil)
}
