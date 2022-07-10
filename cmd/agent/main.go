package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"strconv"
	"sync"
	"syscall"
	"time"

	"github.com/UserNaMEeman/yandex-practic-devops/cmd/agent/metric"
)

const (
	pollInterval   time.Duration = 2 * time.Second
	reportInterval time.Duration = 10 * time.Second
)

type config struct {
	addr           string
	pollInterval   time.Duration
	reportInterval time.Duration
}

func defEnv() config {

	curentConfig := config{}

	pollIntervalVal := os.Getenv("POLL_INTERVAL")
	reportIntervalVal := os.Getenv("REPORT_INTERVAL")
	addr := os.Getenv("ADDRESS")

	if pollIntervalVal == "" {
		pollIntervalVal = "2"
	}
	if reportIntervalVal == "" {
		reportIntervalVal = "10"
	}
	if addr == "" {
		addr = "localhost:8080"
	}

	tp, _ := strconv.Atoi(pollIntervalVal)
	curentConfig.pollInterval = time.Duration(tp) * time.Second
	tr, _ := strconv.Atoi(reportIntervalVal)
	curentConfig.reportInterval = time.Duration(tr) * time.Second
	return curentConfig
}

func collectMetrics(met *metric.Metrics) {
	met.SetMetrics()
}

func main() {
	// myConfig := defEnv()
	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, syscall.SIGTERM, syscall.SIGINT, syscall.SIGQUIT)
	ctx := context.Background()
	met := metric.NewMetrics()
	var mutex sync.Mutex

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
			mutex.Lock()
			collectMetrics(met)
			// fmt.Println(time.Now())
			mutex.Unlock()
			time.Sleep(pollInterval)
			// time.Sleep(myConfig.pollInterval)
		}
	}()

	go func() {
		for {
			// time.Sleep(reportInterval)
			mutex.Lock()
			// targAddr := "http://" + myConfig.addr + "/update/"
			// met.MetricPOST(targAddr)
			met.MetricPOST("http://localhost:8080/update/")
			// fmt.Println("POST", time.Now())
			mutex.Unlock()
			time.Sleep(reportInterval)
			// time.Sleep(myConfig.reportInterval)
		}
	}()
}
