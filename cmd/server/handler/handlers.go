package handler

import (
	"fmt"
	"log"
	"net/http"
	"net/url"
	"strconv"
	"strings"

	// "metric-server/service"
	"github.com/UserNaMEeman/yandex-practic-devops/cmd/server/storage"
)

func checkRequest(w http.ResponseWriter, r *http.Request) bool {
	// urlQ := fmt.Sprintf("%v", r.URL)
	p, err := url.Parse(fmt.Sprintf("%v", r.URL))
	if err != nil {
		panic(err)
	}
	path := p.Path
	if r.Method != http.MethodPost {
		// fmt.Println("not POST")
		http.Error(w, "Only POST requests are allowed", http.StatusMethodNotAllowed)
		return false
	}
	if r.Header.Get("Content-Type") != "text/plain" {
		if r.Header.Get("Content-Type") != "" {
			fmt.Println(r.Header.Get("Content-Type"))
			http.Error(w, "Content-Type must be text/plain; charset=utf-8", 500)
			return false
		}
	}
	if len(strings.Split(path, "/")) != 5 {
		fmt.Println("Invalid URL", path, len(strings.Split(path, "/")))
		http.Error(w, "Invalid URL", 404)
		return false
	}
	if strings.Split(path, "/")[1] != "update" {
		http.Error(w, "Invalid URL", 501)
		fmt.Println("Invalid URL")
		return false
	}
	if strings.Split(path, "/")[2] != "gauge" {
		if strings.Split(path, "/")[2] != "counter" {
			http.Error(w, "Invalid URL", 501)
			fmt.Println(strings.Split(path, "/")[2])
			return false
		}
	}

	_, err = strconv.Atoi(strings.Split(path, "/")[4])
	if err != nil {
		http.Error(w, "Invalid value", 400)
		return false
	}
	_, err = strconv.ParseFloat(strings.Split(path, "/")[4], 64)
	if err != nil {
		http.Error(w, "Invalid value", 400)
		return false
	}

	return true
}

func HandleMetric(w http.ResponseWriter, r *http.Request) {
	// metrics := storage.RecieveMetrics
	var recMetric storage.DataStore
	state := checkRequest(w, r)
	if !state {
		fmt.Println(state)
		return
	}
	// fmt.Println("ok")
	p, err := url.Parse(fmt.Sprintf("%v", r.URL))
	if err != nil {
		panic(err)
	}
	path := p.Path
	elemData := strings.Split(path, "/")
	if elemData[2] == "guage" {
		value, err := strconv.ParseFloat(elemData[4], 64)
		if err != nil {
			log.Printf("%s", err)
		}
		recMetric.Name = elemData[3]
		recMetric.Type = "guage"
		recMetric.ValueF = value
	} else {
		value, err := strconv.Atoi(elemData[4])
		if err != nil {
			log.Printf("%s", err)
		}
		recMetric.Name = elemData[3]
		recMetric.Type = "counter"
		recMetric.ValueC = int64(value)
	}
	http.Error(w, "valid data", http.StatusOK)
	// recMetric.SaveData()
	// w.WriteHeader(http.StatusOK)
	fmt.Println(recMetric)
}
