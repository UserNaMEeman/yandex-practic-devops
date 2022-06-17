package metric

import (
	"fmt"
	"math/rand"
	"net/http"
	"reflect"
	"runtime"
	"strconv"
)

func (statmetric *Metrics) SetGuage() {
	rtm := &runtime.MemStats{}
	runtime.ReadMemStats(rtm)
	RandomValue := Metric{
		Name:   "RandomValue",
		Type:   "guage",
		ValueF: rand.Float64(),
	}
	statmetric.addMetrics(RandomValue)
	val := reflect.ValueOf(rtm).Elem()
	for i := 0; i < val.NumField(); i++ {
		a := fmt.Sprint(val.FieldByName(val.Type().Field(i).Name))
		value, _ := strconv.ParseFloat(a, 64)
		m := Metric{
			Name:   val.Type().Field(i).Name,
			Type:   "guage",
			ValueF: value,
		}
		statmetric.addMetrics(m)
	}
	if statmetric.M["PollCount"].Name == "" {
		countM := Metric{
			Name:   "PollCount",
			Type:   "counter",
			ValueC: 0,
		}
		statmetric.addMetrics(countM)
	} else {
		statmetric.countCount()
	}
}

func (sm *Metrics) MetricPOST(url string) error {
	// "http://localhost:8080/update/gauge/" + key + "/" + aVal
	var postURL string
	client := http.Client{}
	for _, i := range sm.M {
		if i.Type == "guage" {
			aVal := fmt.Sprint(i.ValueF)
			postURL = url + i.Type + "/" + i.Name + "/" + aVal
		} else {
			aVal := fmt.Sprint(i.ValueC)
			postURL = url + i.Type + "/" + i.Name + "/" + aVal
		}
		request, err := http.NewRequest(http.MethodPost, postURL, nil)
		if err != nil {
			fmt.Println(err)
		}
		request.Header.Set("Content-Type", "text/plain")
		response, err := client.Do(request)
		if err != nil {
			// fmt.Println(err)
			return err
		}
		defer response.Body.Close()
	}
	return nil
}
