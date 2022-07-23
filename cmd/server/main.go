package main

import (
	"flag"
	"fmt"
	"net/http"
	"os"
	"strconv"
	"time"

	"github.com/UserNaMEeman/yandex-practic-devops/cmd/server/handler"
	"github.com/UserNaMEeman/yandex-practic-devops/cmd/server/storage"
	"github.com/go-chi/chi"
)

type config struct {
	addrServ      *string
	storeInterval *time.Duration
	storeFile     *string
	restore       *bool
}

func defEnv() (config, bool) {
	useEnv := false
	currentConfig := config{}
	addr, stateAddr := os.LookupEnv("ADDRESS")
	storeInterval, stateStoreInterval := os.LookupEnv("STORE_INTERVAL")
	storeFile, statestoreFile := os.LookupEnv("STORE_FILE")
	restore, staterestore := os.LookupEnv("RESTORE")
	if !stateAddr {
		addr = "localhost:8080"
	} else {
		useEnv = true
	}
	if !stateStoreInterval {
		storeInterval = "300"
	} else {
		useEnv = true
	}
	if !statestoreFile {
		storeFile = "/tmp/devops-metrics-db.json"
	} else {
		useEnv = true
	}
	if !staterestore {
		restore = "true"
	} else {
		useEnv = true
	}
	tp, _ := strconv.Atoi(storeInterval)
	si := time.Duration(tp) * time.Second
	rs, _ := strconv.ParseBool(restore)
	currentConfig.storeInterval = &si
	currentConfig.restore = &rs
	currentConfig.addrServ = &addr
	currentConfig.storeFile = &storeFile
	return currentConfig, useEnv
}

var flagConfig config

func init() {
	flagConfig.addrServ = flag.String("a", "localhost:8080", "listenning host and port")
	flagConfig.storeInterval = flag.Duration("i", 300*time.Second, "storeInterval")
	flagConfig.storeFile = flag.String("f", "/tmp/devops-metrics-db.json", "store metrics file")
	flagConfig.restore = flag.Bool("r", true, "restore metrics from file (true of false)")
}

func main() {
	currentConfig, state := defEnv()
	if !state {
		flag.Parse()
		currentConfig = flagConfig
	}

	var recMetric storage.Metrics
	pullMetrics := make(map[string]storage.Metrics)
	r := chi.NewRouter()
	// r.Use(middleware.RequestID)
	// r.Use(middleware.RealIP)
	// r.Use(middleware.Recoverer)
	// r.Use(middleware.Logger)

	fmt.Println("listen on ", *currentConfig.addrServ)
	fmt.Println("restore ", *currentConfig.restore)
	fmt.Println("storeFile ", *currentConfig.storeFile)
	fmt.Println("storeInterval ", *currentConfig.storeInterval)

	if *currentConfig.restore {
		// fmt.Println(currentConfig.storeFile)
		// storage.GetDataFromFile(currentConfig.storeFile)
		tempMetrics, err := storage.GetDataFromFile(*currentConfig.storeFile)
		if err == nil {
			pullMetrics = tempMetrics
		}
		// fmt.Println(pullMetrics)
	}

	if *currentConfig.storeFile != "" && *currentConfig.storeInterval != 0*time.Second {
		ticker := time.NewTicker(*currentConfig.storeInterval) //currentConfig.storeInterval
		defer ticker.Stop()
		go func() {
			for {
				<-ticker.C
				storage.StoreData(pullMetrics, *currentConfig.storeFile)
			}
		}()
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
			if *currentConfig.storeFile != "" && *currentConfig.storeInterval == 0*time.Second {
				storage.StoreData(pullMetrics, *currentConfig.storeFile)
			}
		})
		r.Post("/", func(w http.ResponseWriter, r *http.Request) {
			recMetric = handler.HandleJSONMetric(w, r, pullMetrics)
			pullMetrics[recMetric.ID] = recMetric
			if *currentConfig.storeFile != "" && *currentConfig.storeInterval == 0*time.Second {
				// fmt.Println("store data")8
				storage.StoreData(pullMetrics, *currentConfig.storeFile)
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
	http.ListenAndServe(*currentConfig.addrServ, r)
}
