package metric

// import (
// 	"errors"
// )

type Metric struct {
	Name   string
	Type   string
	ValueF float64
	ValueC int64
}

type Metrics struct {
	M map[string]Metric
}

func NewMetrics() *Metrics {
	return &Metrics{
		M: map[string]Metric{},
	}
}

func (sm *Metrics) countCount() {
	valueCount := sm.M["PollCount"].ValueC
	sVal := Metric{
		Name:   "PollCount",
		Type:   "counter",
		ValueC: valueCount + 1,
	}
	sm.M["PollCount"] = sVal
}

func (sm *Metrics) addMetrics(am Metric) {
	// if m.M[am.Name].Name != "" {
	// 	return errors.New("metric " + am.Name + " already exist")
	// }
	sm.M[am.Name] = am
}
