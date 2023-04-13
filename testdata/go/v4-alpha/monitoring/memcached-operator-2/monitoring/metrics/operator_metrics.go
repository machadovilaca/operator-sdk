package metrics

import (
	"github.com/example/memcached-operator/monitoring/metrics/registry"
	"github.com/prometheus/client_golang/prometheus"
	dto "github.com/prometheus/client_model/go"
)

const (
	reconcileCount   = metricPrefix + "reconcile_count"
	reconcileAction  = metricPrefix + "reconcile_action_count"
	histogramExample = metricPrefix + "histogram_example"
)

var operatorMetrics = []*registry.MetricDesc{
	{
		FqName:         reconcileCount,
		Help:           "Number of times the operator has executed the reconcile loop",
		MType:          dto.MetricType_COUNTER,
		StabilityLevel: registry.STABLE,
		InitFunc:       registry.NewCounter,
		ConstLabelPairs: map[string]string{
			"controller": "guestbook",
		},
	},
	{
		FqName:         reconcileAction,
		Help:           "Number of times the operator has executed the reconcile loop with a given action",
		MType:          dto.MetricType_COUNTER,
		Labels:         []string{"action"},
		StabilityLevel: registry.ALPHA,
		InitFunc:       registry.NewCounterVec,
	},
	{
		FqName:         histogramExample,
		Help:           "Example of a histogram metric",
		MType:          dto.MetricType_HISTOGRAM,
		Labels:         []string{"label"},
		StabilityLevel: registry.ALPHA,
		InitFunc:       registry.NewHistogramVec,
		ExtraHistogramOpts: &prometheus.HistogramOpts{
			Buckets: []float64{0.1, 0.2, 0.3, 0.4, 0.5},
		},
	},
}

func IncrementReconcileCountMetric() error {
	c, err := registry.GetMetricCollector(reconcileCount)
	if err != nil {
		return err
	}

	c.(prometheus.Counter).Inc()
	return nil
}

func IncrementReconcileActionMetric(action string) error {
	c, err := registry.GetMetricCollector(reconcileAction)
	if err != nil {
		return err
	}

	c.(*prometheus.CounterVec).WithLabelValues(action).Inc()
	return nil
}

func ObserveHistogramExampleMetric(label string, value float64) error {
	c, err := registry.GetMetricCollector(histogramExample)
	if err != nil {
		return err
	}

	c.(*prometheus.HistogramVec).WithLabelValues(label).Observe(value)
	return nil
}
