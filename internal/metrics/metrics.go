package metrics

import (
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
)

type Metrics interface {
	CreateRuleCounterInc()
	UpdateRuleCounterInc()
	RemoveRuleCounterInc()
}

type metrics struct {
	createRuleCounter prometheus.Counter
	updateRuleCounter prometheus.Counter
	removeRuleCounter prometheus.Counter
}

func NewMetrics() Metrics {
	return &metrics{
		createRuleCounter: promauto.NewCounter(prometheus.CounterOpts{
			Name: "create_rule_total",
			Help: "Количество созданий правил",
		}),
		updateRuleCounter: promauto.NewCounter(prometheus.CounterOpts{
			Name: "update_rule_total",
			Help: "Количество обновлений правил",
		}),
		removeRuleCounter: promauto.NewCounter(prometheus.CounterOpts{
			Name: "remove_rule_total",
			Help: "Количество удалений правил",
		}),
	}
}

func (m *metrics) CreateRuleCounterInc() {
	m.createRuleCounter.Inc()
}

func (m *metrics) UpdateRuleCounterInc() {
	m.updateRuleCounter.Inc()
}

func (m *metrics) RemoveRuleCounterInc() {
	m.removeRuleCounter.Inc()
}
