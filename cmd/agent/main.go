package main

import (
	"github.com/UserNaMEeman/yandex-practic-devops/cmd/agent/metric"
	"os/signal"
	"time"
	"sync"
	"fmt"
	"os"
	"syscall"
	"context"
)

const(
	pollInterval int = 2
	reportInterval int = 10
)

func collectMetrics(met *metric.Metrics) {
	met.SetGuage()
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

	go func(){for{
		mutex.Lock()
		collectMetrics(met)
		mutex.Unlock()
		time.Sleep(2 * time.Second)
	}}()

	go func(){for{
		mutex.Lock()
		met.MetricPOST("http://localhost:8080/update/")
		mutex.Unlock()
		time.Sleep(10 * time.Second)
	}}()
}