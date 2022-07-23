package main

import (
	"context"
	"flag"
	"fmt"
	"os"
	"os/signal"
	"strconv"
	"sync"
	"syscall"
	"time"

	"github.com/UserNaMEeman/yandex-practic-devops/cmd/agent/metric"
)

// const (
// 	pollInterval   time.Duration = 2 * time.Second
// 	reportInterval time.Duration = 10 * time.Second
// )

type config struct {
	addr           *string
	pollInterval   *time.Duration
	reportInterval *time.Duration
}

func defEnv() (config, bool) {
	useEnv := false
	curentConfig := config{}
	pollIntervalVal, statePoll := os.LookupEnv("POLL_INTERVAL")
	reportIntervalVal, stateRep := os.LookupEnv("REPORT_INTERVAL")
	addr, stateAddr := os.LookupEnv("ADDRESS")

	if !statePoll {
		pollIntervalVal = "2"
	} else {
		useEnv = true
	}
	if !stateRep {
		reportIntervalVal = "10"
	} else {
		useEnv = true
	}

	if !stateAddr {
		addr = "127.0.0.1:8080"
	} else {
		useEnv = true
	}

	tp, _ := strconv.Atoi(pollIntervalVal)
	vp := time.Duration(tp) * time.Second
	curentConfig.pollInterval = &vp
	tr, _ := strconv.Atoi(reportIntervalVal)
	vr := time.Duration(tr) * time.Second
	curentConfig.reportInterval = &vr
	curentConfig.addr = &addr
	return curentConfig, useEnv
}

func collectMetrics(met *metric.Metrics) {
	met.SetMetrics()
}

var flagConfig config

func init() {
	flagConfig.addr = flag.String("a", "localhost:8080", "target host and port")
	flagConfig.pollInterval = flag.Duration("p", 2*time.Second, "pollInterval")
	flagConfig.reportInterval = flag.Duration("r", 10*time.Second, "reportInterval")
}

func main() {
	myConfig, state := defEnv()
	if !state {
		flag.Parse()
		myConfig = flagConfig
	}
	fmt.Println("send to ", *myConfig.addr)
	fmt.Println("pollInterval ", *myConfig.pollInterval)
	fmt.Println("reportInterval ", *myConfig.reportInterval)
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
			// time.Sleep(pollInterval)
			time.Sleep(*myConfig.pollInterval)
		}
	}()

	go func() {
		for {
			// time.Sleep(reportInterval)
			mutex.Lock()
			targAddr := "http://" + *myConfig.addr + "/update/"
			met.MetricPOST(targAddr)
			// met.MetricPOST("http://localhost:8080/update/")
			// fmt.Println("POST", time.Now())
			mutex.Unlock()
			// time.Sleep(reportInterval)
			time.Sleep(*myConfig.reportInterval)
		}
	}()
}
