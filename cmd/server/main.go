package main

import(
	"github.com/UserNaMEeman/yandex-practic-devops/cmd/server/handler"
	"net/http"
)

func main() {
	http.HandleFunc("/update/", handler.GetMetric)
	http.ListenAndServe(":8080", nil)
}
