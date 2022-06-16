package metric

import (
	"fmt"
	"math/rand"
	"net/http"
	"reflect"
	"runtime"
	"strconv"
)

func (sm *Metrics) SetGuage() {
	rtm := &runtime.MemStats{}
	runtime.ReadMemStats(rtm)
	RandomValue := Metric{
		Name:   "RandomValue",
		Type:   "guage",
		ValueF: rand.Float64(),
	}
	sm.addMetrics(RandomValue)
	val := reflect.ValueOf(rtm).Elem()
	for i := 0; i < val.NumField(); i++ {
		a := fmt.Sprint(val.FieldByName(val.Type().Field(i).Name))
		value, _ := strconv.ParseFloat(a, 64)
		m := Metric{
			Name:   val.Type().Field(i).Name,
			Type:   "guage",
			ValueF: value,
		}
		sm.addMetrics(m)
	}
	if sm.M["PollCount"].Name == ""{
		countM := Metric{
			Name: "PollCount",
			Type: "counter",
			ValueC: 0,
		}
		sm.addMetrics(countM)
	} else{
		sm.countCount()
	}
}

func (sm *Metrics) MetricPOST(url string) error {
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
