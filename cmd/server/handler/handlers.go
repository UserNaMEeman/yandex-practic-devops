package handler

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"io/ioutil"
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
		fmt.Println("not POST")
		http.Error(w, "Only POST requests are allowed", http.StatusMethodNotAllowed)
		return false
	}
	if r.Header.Get("Content-Type") != "text/plain" {
		if r.Header.Get("Content-Type") != "" {
			// fmt.Println(r.Header.Get("Content-Type"))
			http.Error(w, "Content-Type must be text/plain; charset=utf-8", http.StatusInternalServerError)
			return false
		}
	}
	if len(strings.Split(path, "/")) != 5 {
		// fmt.Println("Invalid URL", path, len(strings.Split(path, "/")))
		http.Error(w, "Invalid URL", http.StatusNotFound)
		return false
	}
	if strings.Split(path, "/")[1] != "update" {
		http.Error(w, "Invalid URL", http.StatusNotImplemented)
		// fmt.Println("Invalid URL")
		return false
	}
	if strings.Split(path, "/")[2] != "gauge" {
		if strings.Split(path, "/")[2] != "counter" {
			http.Error(w, "Invalid URL", http.StatusNotImplemented)
			// fmt.Println(strings.Split(path, "/")[2])
			return false
		}
	}
	if strings.Split(path, "/")[2] == "counter" {
		_, err = strconv.Atoi(strings.Split(path, "/")[4])
		if err != nil {
			// fmt.Println(strings.Split(path, "/")[4], "Invalid value", err)
			http.Error(w, "Invalid value", 400)
			return false
		}
	}
	if strings.Split(path, "/")[2] == "gauge" {
		_, err = strconv.ParseFloat(strings.Split(path, "/")[4], 64)
		if err != nil {
			// fmt.Println(strings.Split(path, "/")[4], "Invalid value", err)
			http.Error(w, "Invalid value", 400)
			return false
		}
	}

	return true
}

func HandleMetric(w http.ResponseWriter, r *http.Request, allMetrics map[string]storage.Metrics) (storage.Metrics, error) {
	var recMetric storage.Metrics
	state := checkRequest(w, r)
	if !state {
		// fmt.Println(state)
		return recMetric, errors.New("bad URL")
	}
	// fmt.Println("ok")
	p, err := url.Parse(fmt.Sprintf("%v", r.URL))
	if err != nil {
		panic(err)
	}
	path := p.Path
	elemData := strings.Split(path, "/")
	if elemData[2] == "gauge" {
		value, err := strconv.ParseFloat(elemData[4], 64)
		if err != nil {
			log.Printf("%s", err)
		}
		recMetric.ID = elemData[3]
		// fmt.Println(recMetric.Name, ":", value)
		recMetric.MType = "gauge"
		recMetric.Value = &value
	} else {
		value, err := strconv.Atoi(elemData[4])
		if err != nil {
			log.Printf("%s", err)
		}
		recMetric.ID = elemData[3]

		if allMetrics[recMetric.ID].ID != "" {
			val := *allMetrics[recMetric.ID].Delta + int64(value)
			recMetric.Delta = &val
		} else {
			val := int64(value)
			recMetric.Delta = &val
		}

		recMetric.MType = "counter"
	}
	return recMetric, nil
	// recMetric.SaveData()
	// fmt.Printf("%v\n", recMetric)
}

func HandleJSONMetric(w http.ResponseWriter, r *http.Request, allMetrics map[string]storage.Metrics) storage.Metrics {
	var recJSONMetric storage.Metrics
	if r.Method != http.MethodPost {
		fmt.Println("not POST")
		http.Error(w, "Only POST requests are allowed", http.StatusMethodNotAllowed)
		return storage.Metrics{}
	}
	if r.Header.Get("Content-Type") != "application/json" {
		// fmt.Println(r.Header.Get("Content-Type"))
		http.Error(w, "Content-Type must be application/json; charset=utf-8", http.StatusInternalServerError)
		return storage.Metrics{}
	}
	// fmt.Println("ok")
	p, err := url.Parse(fmt.Sprintf("%v", r.URL))
	if err != nil {
		panic(err)
	}
	path := p.Path
	elemData := strings.Split(path, "/")
	if elemData[1] != "update" {
		http.Error(w, "link is bad", http.StatusBadRequest)
		return storage.Metrics{}
	}
	data, err := ioutil.ReadAll(r.Body)
	if err != nil {
		http.Error(w, err.Error(), 500)
	}
	// fmt.Println(string(data))

	if err := json.Unmarshal(data, &recJSONMetric); err != nil {
		fmt.Println(err)
	}

	if allMetrics[recJSONMetric.ID].MType == "counter" {
		// fmt.Println(*allMetrics[recJSONMetric.ID].Delta)
		if allMetrics[recJSONMetric.ID].ID != "" {
			a := *allMetrics[recJSONMetric.ID].Delta + *recJSONMetric.Delta
			recJSONMetric.Delta = &a
		}
		// if *allMetrics[recJSONMetric.ID].Delta != 0 {
		// 	a := *allMetrics[recJSONMetric.ID].Delta + *recJSONMetric.Delta
		// 	recJSONMetric.Delta = &a
		// }
	}
	// fmt.Printf("%+v", recJSONMetric)
	return recJSONMetric
}

func ShowJSONMetrics(w http.ResponseWriter, r *http.Request, allMetrics map[string]storage.Metrics) {
	var reqJSON storage.Metrics
	var respJSON storage.Metrics
	if r.Header.Get("Content-Type") != "application/json" {
		// fmt.Println(r.Header.Get("Content-Type"))
		http.Error(w, "Content-Type must be application/json; charset=utf-8", http.StatusInternalServerError)
		return
	}
	body, err := io.ReadAll(r.Body)
	if err != nil {
		fmt.Println(err)
		return
	}
	err = json.Unmarshal(body, &reqJSON)
	if err != nil {
		fmt.Println(err)
	}

	// if reqJSON.ID == "" || reqJSON.MType == "" {
	// 	w.WriteHeader(http.StatusNotFound)
	// 	return
	// }
	for _, i := range allMetrics {
		if i.ID == reqJSON.ID && i.MType == reqJSON.MType {
			respJSON = i
			break
		}
	}
	sendData, err := json.Marshal(respJSON)
	if err != nil {
		fmt.Println(err)
	}
	// fmt.Println(allMetrics)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	// w.Header().Set("Content-Type", "application/json")
	fmt.Fprint(w, string(sendData))
	// fmt.Printf("%+v\n", reqJSON)
}

func ShowAllMetrics(w http.ResponseWriter, pullMetrics map[string]storage.Metrics) {
	w.WriteHeader(http.StatusOK)
	for _, i := range pullMetrics {
		// w.WriteHeader(http.StatusAccepted)
		if i.MType == "gauge" {
			fmt.Fprintf(w, "Name: %s	Type: %s	Value: %v\n", i.ID, i.MType, *i.Value)
		} else {
			fmt.Fprintf(w, "Name: %s	Type: %s	Value: %v\n", i.ID, i.MType, *i.Delta)
		}
	}
}

func ShowOneMetric(w http.ResponseWriter, r *http.Request, pullMetrics map[string]storage.Metrics) {
	// w.WriteHeader(http.StatusForbidden)
	p, err := url.Parse(fmt.Sprintf("%v", r.URL))
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Println(err)
	}
	path := p.Path
	elemData := strings.Split(path, "/")
	valType := elemData[2]
	valName := elemData[3]
	if pullMetrics[valName].ID == "" {
		w.WriteHeader(http.StatusNotFound)
		return
	}
	if valType == "gauge" {
		w.WriteHeader(http.StatusOK)
		fmt.Fprintf(w, "%v\n", *pullMetrics[valName].Value)
	} else {
		w.WriteHeader(http.StatusOK)
		fmt.Fprintf(w, "%v\n", *pullMetrics[valName].Delta)
	}
	// http.Error(w, "valid data", http.StatusOK)
}

// func ShowAllMetrics(w http.ResponseWriter, r *http.Request) {
// 	storage.SelectAllMetrics(w)
// }

// func ShowMetrics(w http.ResponseWriter, r *http.Request) {
// 	p, err := url.Parse(fmt.Sprintf("%v", r.URL))
// 	if err != nil {
// 		panic(err)
// 	}
// 	path := p.Path
// 	elemData := strings.Split(path, "/")
// 	valName := fmt.Sprintf("%s", elemData[3])
// 	storage.SelectMetric(w, valName)
// }
