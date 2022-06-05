package main

import (
	"context"
	"fmt"
	"math/rand"
	"net/http"
	"os"
	"os/signal"
	"runtime"
	"syscall"
	"time"
)

type gauge float64
type counter int64

func POST(url string, client http.Client) {
	request, err := http.NewRequest(http.MethodPost, url, nil)
	if err != nil {
		fmt.Println(err)
	}
	request.Header.Set("Content-Type", "text/plain")
	_, err = client.Do(request)
	if err != nil {
		fmt.Println(err)
	}
}

func main() {
	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, syscall.SIGTERM, syscall.SIGINT, syscall.SIGQUIT)
	ctx := context.Background()
	mt := &runtime.MemStats{}
	var PollCount counter = 0
	pollInterval := 2 * time.Second
	reportInterval := 10 * time.Second
	var RandomValue gauge = gauge(rand.Float64())
	var sendMetrics map[string]gauge
	sendMetrics = make(map[string]gauge)

	go func() {
		c := <-sigChan
		fmt.Println("signal is ", c)
		os.Exit(0)
	}()
	defer func() {
		<-ctx.Done()
	}()

	go func() {
		for {
			runtime.ReadMemStats(mt)
			sendMetrics["Alloc"] = gauge(mt.Alloc)
			sendMetrics["BuckHashSys"] = gauge(mt.BuckHashSys)
			sendMetrics["Frees"] = gauge(mt.Frees)
			sendMetrics["GCCPUFraction"] = gauge(mt.GCCPUFraction)
			sendMetrics["GCSys"] = gauge(mt.GCSys)
			sendMetrics["HeapAlloc"] = gauge(mt.HeapAlloc)
			sendMetrics["HeapIdle"] = gauge(mt.HeapIdle)
			sendMetrics["HeapInuse"] = gauge(mt.HeapInuse)
			sendMetrics["HeapObjects"] = gauge(mt.HeapObjects)
			sendMetrics["HeapReleased"] = gauge(mt.HeapReleased)
			sendMetrics["HeapSys"] = gauge(mt.HeapSys)
			sendMetrics["LastGC"] = gauge(mt.LastGC)
			sendMetrics["Lookups"] = gauge(mt.Lookups)
			sendMetrics["MCacheInuse"] = gauge(mt.MCacheInuse)
			sendMetrics["MCacheSys"] = gauge(mt.MCacheSys)
			sendMetrics["MSpanInuse"] = gauge(mt.MSpanInuse)
			sendMetrics["MSpanSys"] = gauge(mt.MSpanSys)
			sendMetrics["Mallocs"] = gauge(mt.Mallocs)
			sendMetrics["NextGC"] = gauge(mt.NextGC)
			sendMetrics["NumForcedGC"] = gauge(mt.NumForcedGC)
			sendMetrics["NumGC"] = gauge(mt.NumGC)
			sendMetrics["OtherSys"] = gauge(mt.OtherSys)
			sendMetrics["PauseTotalNs"] = gauge(mt.PauseTotalNs)
			sendMetrics["StackInuse"] = gauge(mt.StackInuse)
			sendMetrics["StackSys"] = gauge(mt.StackSys)
			sendMetrics["Sys"] = gauge(mt.Sys)
			sendMetrics["TotalAlloc"] = gauge(mt.TotalAlloc)
			RandomValue = gauge(rand.Float64())
			PollCount++
			time.Sleep(pollInterval)
		}
	}()

	go func() {
		client := http.Client{}
		for {
			time.Sleep(reportInterval)
			for key, value := range sendMetrics {
				aVal := fmt.Sprint(value)
				url := "http://127.0.0.1:8080/update/gauge/" + key + "/" + aVal
				POST(url, client)
			}
			aPoll := fmt.Sprint(PollCount)
			url := "http://127.0.0.1:8080/update/counter/PollCount/" + aPoll
			POST(url, client)
			aRand := fmt.Sprint(RandomValue)
			url = "http://127.0.0.1:8080/update/gauge/RandomValue/" + aRand
			POST(url, client)
		}
	}()

}
