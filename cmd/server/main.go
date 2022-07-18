package main

import (
	"net/http"
	"os"
	"strconv"
	"time"

	"github.com/UserNaMEeman/yandex-practic-devops/cmd/server/handler"
	"github.com/UserNaMEeman/yandex-practic-devops/cmd/server/storage"
	"github.com/go-chi/chi"
)

type config struct {
	addrServ      string
	storeInterval time.Duration
	storeFile     string
	restore       bool
}

func defEnv() config {
	currentConfig := config{}
	addr, stateAddr := os.LookupEnv("ADDRESS")
	storeInterval, stateStoreInterval := os.LookupEnv("STORE_INTERVAL")
	storeFile, statestoreFile := os.LookupEnv("STORE_FILE")
	restore, staterestore := os.LookupEnv("RESTORE")
	if !stateAddr {
		addr = "127.0.0.1:8080"
	}
	if !stateStoreInterval {
		storeInterval = "300"
	}
	if !statestoreFile {
		storeFile = "/tmp/devops-metrics-db.json"
	}
	if !staterestore {
		restore = "true"
	}
	tp, _ := strconv.Atoi(storeInterval)
	currentConfig.storeInterval = time.Duration(tp) * time.Second
	currentConfig.restore, _ = strconv.ParseBool(restore)
	currentConfig.addrServ = addr
	currentConfig.storeFile = storeFile
	return currentConfig
}

func main() {
	currentConfig := defEnv()

	var recMetric storage.Metrics
	pullMetrics := make(map[string]storage.Metrics)
	r := chi.NewRouter()
	// r.Use(middleware.RequestID)
	// r.Use(middleware.RealIP)
	// r.Use(middleware.Recoverer)
	// r.Use(middleware.Logger)

	if currentConfig.storeFile != "" && currentConfig.storeInterval != 0*time.Second {
		ticker := time.NewTicker(currentConfig.storeInterval) //currentConfig.storeInterval
		defer ticker.Stop()
		go func() {
			for {
				<-ticker.C
				storage.StoreData(pullMetrics, currentConfig.storeFile)
			}
		}()
	}
	if currentConfig.restore {
		// fmt.Println(currentConfig.storeFile)
		// storage.GetDataFromFile(currentConfig.storeFile)
		pullMetrics = storage.GetDataFromFile(currentConfig.storeFile)
		// fmt.Println(pullMetrics)
	}
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
			if currentConfig.storeFile != "" && currentConfig.storeInterval == 0*time.Second {
				storage.StoreData(pullMetrics, currentConfig.storeFile)
			}
		})
		r.Post("/", func(w http.ResponseWriter, r *http.Request) {
			recMetric = handler.HandleJSONMetric(w, r, pullMetrics)
			pullMetrics[recMetric.ID] = recMetric
			if currentConfig.storeFile != "" && currentConfig.storeInterval == 0*time.Second {
				storage.StoreData(pullMetrics, currentConfig.storeFile)
			}
			// fmt.Println(JSONMetrics)
		})
	})
	// r.Route("/value", func(r chi.Router) {
	r.Post("/value/", func(w http.ResponseWriter, r *http.Request) {
		handler.ShowJSONMetrics(w, r, pullMetrics)
	})
	// }
	// addrServ := "localhost:8080"
	http.ListenAndServe(currentConfig.addrServ, r)
}
