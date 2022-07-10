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

// type Metrics struct {
// 	M map[string]Metric
// }

type Metrics struct {
	M map[string]Metric2
}

type Metric2 struct {
	ID    string   `json:"id"`              // имя метрики
	MType string   `json:"type"`            // параметр, принимающий значение gauge или counter
	Delta *int64   `json:"delta,omitempty"` // значение метрики в случае передачи counter
	Value *float64 `json:"value,omitempty"` // значение метрики в случае передачи gauge
}

func NewMetrics() *Metrics {
	return &Metrics{
		M: map[string]Metric2{},
	}
}

// func NewMetrics() *Metrics {
// 	return &Metrics{
// 		M: map[string]Metric{},
// 	}
// }

func (sm *Metrics) countCount() {
	valueCount := sm.M["PollCount"].Delta
	delta := *valueCount + 1
	sVal := Metric2{
		ID:    "PollCount",
		MType: "counter",
		Delta: &delta,
	}
	sm.M["PollCount"] = sVal
}

// func (sm *Metrics) countCount() {
// 	valueCount := sm.M["PollCount"].ValueC
// 	sVal := Metric{
// 		Name:   "PollCount",
// 		Type:   "counter",
// 		ValueC: valueCount + 1,
// 	}
// 	sm.M["PollCount"] = sVal
// }

func (sm *Metrics) addMetrics(am Metric2) {
	// if m.M[am.Name].Name != "" {
	// 	return errors.New("metric " + am.Name + " already exist")
	// }
	sm.M[am.ID] = am
}

// func (sm *Metrics) addMetrics(am Metric) {
// 	// if m.M[am.Name].Name != "" {
// 	// 	return errors.New("metric " + am.Name + " already exist")
// 	// }
// 	sm.M[am.Name] = am
// }
