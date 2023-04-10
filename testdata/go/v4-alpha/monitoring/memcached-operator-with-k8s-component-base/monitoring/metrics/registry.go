package metrics

import (
	"github.com/prometheus/client_golang/prometheus"
	dto "github.com/prometheus/client_model/go"
	"k8s.io/component-base/metrics/legacyregistry"
)

type Registry struct{}

func NewRegistry() Registry {
	return Registry{}
}

func (r Registry) Register(_ prometheus.Collector) error {
	return nil
}

func (r Registry) MustRegister(_ ...prometheus.Collector) {}

func (r Registry) Unregister(_ prometheus.Collector) bool {
	return false
}

func (r Registry) Gather() ([]*dto.MetricFamily, error) {
	return legacyregistry.DefaultGatherer.Gather()
}
