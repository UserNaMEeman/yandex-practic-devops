package metric

import (
	"errors"
)

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

func (m *Metrics) AddMetrics(am Metric) error {
	if m.M[am.Name].Name != "" {
		return errors.New("metric " + am.Name + " already exist")
	}
	m.M[am.Name] = am
	return nil
}

// type CounterM struct {
// 	M map[string]CounterMetric
// }

// func NewCounterM() *CounterM{
// 	return &CounterM{
// 		M: map[string]CounterMetric{}
// 	}
// }

// type CounterMetric struct {
// 	Name  string
// 	Type  string	`counter`
// 	Value int64
// }
