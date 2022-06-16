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

func (m *Metrics) countCount(){
	valueCount := m.M["PollCount"].ValueC
	sVal := Metric{
		Name: "PollCount",
		Type: "counter",
		ValueC: valueCount + 1,
	}
	m.M["PollCount"] = sVal
}

func (m *Metrics) addMetrics(am Metric) {
	// if m.M[am.Name].Name != "" {
	// 	return errors.New("metric " + am.Name + " already exist")
	// }
	m.M[am.Name] = am
}