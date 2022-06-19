package metric

import (
	"fmt"
	"math/rand"
	"net/http"
	"reflect"
	"runtime"
	"strconv"
)

func (sm *Metrics) SetMetrics() {
	rtm := &runtime.MemStats{}
	runtime.ReadMemStats(rtm)
	RandomValue := Metric{
		Name:   "RandomValue",
		Type:   "gauge",
		ValueF: rand.Float64(),
	}
	sm.addMetrics(RandomValue)
	val := reflect.ValueOf(rtm).Elem()
	for i := 0; i < val.NumField(); i++ {
		a := fmt.Sprint(val.FieldByName(val.Type().Field(i).Name))
		value, _ := strconv.ParseFloat(a, 64)
		m := Metric{
			Name:   val.Type().Field(i).Name,
			Type:   "gauge",
			ValueF: value,
		}
		sm.addMetrics(m)
	}
	if sm.M["PollCount"].Name == "" {
		countM := Metric{
			Name:   "PollCount",
			Type:   "counter",
			ValueC: 0,
		}
		sm.addMetrics(countM)
	} else {
		sm.countCount()
	}
}

func (sm *Metrics) MetricPOST(url string) error {
	// "http://localhost:8080/update/gauge/" + key + "/" + aVal
	var postURL string
	client := http.Client{}
	for _, i := range sm.M {
		if i.Type == "gauge" {
			aVal := fmt.Sprintf("%f", i.ValueF)
			postURL = url + i.Type + "/" + i.Name + "/" + aVal
			// fmt.Printf("%v:%f:%s\n", aVal, i.ValueF, i.Name)
		} else {
			aVal := fmt.Sprintf("%d", i.ValueC)
			postURL = url + i.Type + "/" + i.Name + "/" + aVal
			// fmt.Printf("%v:%d:%s\n", aVal, i.ValueC, i.Name)
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
