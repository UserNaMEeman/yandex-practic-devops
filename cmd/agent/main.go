package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"sync"
	"syscall"
	"time"

	"github.com/UserNaMEeman/yandex-practic-devops/cmd/agent/metric"
)

const (
	pollInterval   time.Duration = 2 * time.Second
	reportInterval time.Duration = 10 * time.Second
)

func collectMetrics(met *metric.Metrics) {
	met.SetMetrics()
}

func main() {
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
		}
	}()

	go func() {
		for {
			// time.Sleep(reportInterval)
			mutex.Lock()
			met.MetricPOST("http://localhost:8080/update/")
			// fmt.Println("POST", time.Now())
			mutex.Unlock()
			time.Sleep(reportInterval)
		}
	}()
}
