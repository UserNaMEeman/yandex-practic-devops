package metric

import (
	"fmt"
	"math/rand"
	"net/http"
	"reflect"
	"runtime"
)

func (sm *Metrics) SetGuage() {
	// var needMetrics []string
	rtm := &runtime.MemStats{}
	runtime.ReadMemStats(rtm)
	RandomValue := Metric{
		Name:   "RandomValue",
		Type:   "guage",
		ValueF: rand.Float64(),
	}
	sm.AddMetrics(RandomValue)
	val := reflect.ValueOf(rtm).Elem()
	for i := 0; i < val.NumField(); i++ {
		m := Metric{
			Name:   val.Type().Field(i).Name,
			Type:   "guage",
			ValueF: float64(val.FieldByName(val.Type().Field(i).Name)),
		}
		sm.AddMetrics(m)
	}
}

func (sm *Metrics) MetricPOSTt(url string) error {
	// "http://localhost:8080/update/gauge/" + key + "/" + aVal
	var postUrl string
	client := http.Client{}
	for _, i := range sm.M {
		if i.Type == "guage" {
			aVal := fmt.Sprint(i.ValueF)
			postUrl = url + i.Type + "/" + i.Name + "/" + aVal
		} else {
			aVal := fmt.Sprint(i.ValueC)
			postUrl = url + i.Type + "/" + i.Name + "/" + aVal
		}
		request, err := http.NewRequest(http.MethodPost, postUrl, nil)
		if err != nil {
			fmt.Println(err)
		}
		request.Header.Set("Content-Type", "text/plain")
		_, err = client.Do(request)
		if err != nil {
			// fmt.Println(err)
			return err
		}
	}
	return nil
}

// func (sm *GuageMetric) MetricPOST(url string) error {
// 	// "http://localhost:8080/update/gauge/" + key + "/" + aVal
// 	client := http.Client{}
// 	aVal := fmt.Sprint(sm.Value)
// 	postUrl := url + sm.Type + "/" + sm.Name + "/" + aVal
// 	request, err := http.NewRequest(http.MethodPost, postUrl, nil)
// 	if err != nil {
// 		fmt.Println(err)
// 	}
// 	request.Header.Set("Content-Type", "text/plain")
// 	_, err = client.Do(request)
// 	if err != nil {
// 		// fmt.Println(err)
// 		return err
// 	}
// 	return nil
// }

// func (sm *GuageMetric) SetMetric(mt *runtime.MemStats) {
// 	// mt := &runtime.MemStats{}
// 	// runtime.ReadMemStats(mt)
// 	sm.Type = "guage"
// 	switch sm.Name {
// 	case "Alloc":
// 		sm.Value = float64(mt.Alloc)
// 		return
// 	case "BuckHashSys":
// 		sm.Value = float64(mt.BuckHashSys)
// 		return
// 	case "Frees":
// 		sm.Value = float64(mt.Frees)
// 		return
// 	case "GCCPUFraction":
// 		sm.Value = float64(mt.GCCPUFraction)
// 		return
// 	case "GCSys":
// 		sm.Value = float64(mt.GCSys)
// 		return
// 	case "HeapAlloc":
// 		sm.Value = float64(mt.HeapAlloc)
// 		return
// 	case "HeapIdle":
// 		sm.Value = float64(mt.HeapIdle)
// 		return
// 	case "HeapInuse":
// 		sm.Value = float64(mt.HeapInuse)
// 		return
// 	case "HeapObjects":
// 		sm.Value = float64(mt.HeapObjects)
// 		return
// 	case "HeapReleased":
// 		sm.Value = float64(mt.HeapReleased)
// 		return
// 	case "HeapSys":
// 		sm.Value = float64(mt.HeapSys)
// 		return
// 	case "LastGC":
// 		sm.Value = float64(mt.LastGC)
// 		return
// 	case "Lookups":
// 		sm.Value = float64(mt.Lookups)
// 		return
// 	case "MCacheInuse":
// 		sm.Value = float64(mt.MCacheInuse)
// 		return
// 	case "MCacheSys":
// 		sm.Value = float64(mt.MCacheSys)
// 		return
// 	case "MSpanInuse":
// 		sm.Value = float64(mt.MSpanInuse)
// 		return
// 	case "MSpanSys":
// 		sm.Value = float64(mt.MSpanSys)
// 		return
// 	case "Mallocs":
// 		sm.Value = float64(mt.Mallocs)
// 		return
// 	case "NextGC":
// 		sm.Value = float64(mt.NextGC)
// 		return
// 	case "NumForcedGC":
// 		sm.Value = float64(mt.NumForcedGC)
// 		return
// 	case "NumGC":
// 		sm.Value = float64(mt.NumGC)
// 		return
// 	case "OtherSys":
// 		sm.Value = float64(mt.OtherSys)
// 		return
// 	case "PauseTotalNs":
// 		sm.Value = float64(mt.PauseTotalNs)
// 		return
// 	case "StackInuse":
// 		sm.Value = float64(mt.StackInuse)
// 		return
// 	case "StackSys":
// 		sm.Value = float64(mt.StackSys)
// 		return
// 	case "Sys":
// 		sm.Value = float64(mt.Sys)
// 		return
// 	case "TotalAlloc":
// 		sm.Value = float64(mt.TotalAlloc)
// 		return
// 	}
// }

// // 	aVal := fmt.Sprint(sm.Value)
// // 	postUrl := url + sm.Type + "/" + sm.Name + "/" + aVal
// // 	request, err := http.NewRequest(http.MethodPost, postUrl, nil)
// // 	if err != nil {
// // 		fmt.Println(err)
// // 	}
// // 	request.Header.Set("Content-Type", "text/plain")
// // 	_, err = client.Do(request)
// // 	if err != nil {
// // 		// fmt.Println(err)
// // 		return err
// // 	}
// // 	return nil
// // }
