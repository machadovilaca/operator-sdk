package registry

import (
	"fmt"
	runtimemetrics "sigs.k8s.io/controller-runtime/pkg/metrics"

	"github.com/prometheus/client_golang/prometheus"
	dto "github.com/prometheus/client_model/go"
)

type CollectorDesc struct {
	metrics     map[string]*MetricDesc
	collectFunc func(ch chan<- prometheus.Metric)
}

var registeredCollectors []*CollectorDesc

func SetupCollectors(collectors []*CollectorDesc) error {
	for _, collector := range collectors {
		err := runtimemetrics.Registry.Register(collector)
		if err != nil {
			return err
		}
		registeredCollectors = append(registeredCollectors, collector)
	}

	return nil
}

func NewCollector(metrics map[string]*MetricDesc, collectFunc func(ch chan<- prometheus.Metric)) *CollectorDesc {
	return &CollectorDesc{
		metrics:     metrics,
		collectFunc: collectFunc,
	}
}

func (co *CollectorDesc) Describe(ch chan<- *prometheus.Desc) {
	for _, metric := range co.metrics {
		desc := prometheus.NewDesc(metric.FqName, metric.Help, metric.Labels, metric.ConstLabelPairs)
		ch <- desc
	}
}

func (co *CollectorDesc) Collect(ch chan<- prometheus.Metric) {
	co.collectFunc(ch)
}

func NewMetricWithValue(metric *MetricDesc, value float64, labels ...string) (prometheus.Metric, error) {
	desc := prometheus.NewDesc(metric.FqName, metric.Help, metric.Labels, metric.ConstLabelPairs)

	valueType, err := dtoToValue(metric.MType)
	if err != nil {
		return nil, err
	}

	return prometheus.NewConstMetric(desc, valueType, value, labels...)
}

func dtoToValue(metricType dto.MetricType) (prometheus.ValueType, error) {
	switch metricType {
	case dto.MetricType_COUNTER:
		return prometheus.CounterValue, nil
	case dto.MetricType_GAUGE:
		return prometheus.GaugeValue, nil
	case dto.MetricType_UNTYPED:
		return prometheus.UntypedValue, nil
	default:
		return prometheus.UntypedValue, fmt.Errorf("invalid collector metric type %v", metricType)
	}
}
