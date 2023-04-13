package registry

import (
	"fmt"

	runtimemetrics "sigs.k8s.io/controller-runtime/pkg/metrics"

	"github.com/prometheus/client_golang/prometheus"
	dto "github.com/prometheus/client_model/go"
)

type MetricDesc struct {
	FqName          string
	Help            string
	MType           dto.MetricType
	ConstLabelPairs map[string]string
	Labels          []string
	StabilityLevel  StabilityLevel

	ExtraHistogramOpts *prometheus.HistogramOpts
	ExtraSummaryOpts   *prometheus.SummaryOpts

	InitFunc func(desc *MetricDesc) (prometheus.Collector, error)

	collector prometheus.Collector
}

type StabilityLevel string

const (
	ALPHA  StabilityLevel = "ALPHA"
	BETA   StabilityLevel = "BETA"
	STABLE StabilityLevel = "STABLE"
)

var registeredMetrics = map[string]*MetricDesc{}

func SetupMetrics(metricsRegisters [][]*MetricDesc) error {
	for _, register := range metricsRegisters {
		for _, metric := range register {
			err := registerMetric(metric)
			if err != nil {
				return err
			}
		}
	}

	return nil
}

func registerMetric(metric *MetricDesc) error {
	if metric.InitFunc == nil {
		return fmt.Errorf("InitFunc not set for metric %v", metric.FqName)
	}

	collector, err := metric.InitFunc(metric)
	if err != nil {
		return err
	}

	fmt.Println("Registering metric", metric.FqName)
	err = runtimemetrics.Registry.Register(collector)
	if err != nil {
		return err
	}

	metric.collector = collector
	registeredMetrics[metric.FqName] = metric

	return nil
}

func ListMetrics() []*MetricDesc {
	var metrics []*MetricDesc

	for _, desc := range registeredMetrics {
		metrics = append(metrics, desc)
	}

	for _, collector := range registeredCollectors {
		for _, desc := range collector.metrics {
			metrics = append(metrics, desc)
		}
	}

	return metrics
}

func GetMetricCollector(metricName string) (prometheus.Collector, error) {
	metric, ok := registeredMetrics[metricName]
	if !ok {
		return nil, fmt.Errorf("metric %v not found", metricName)
	}

	return metric.collector, nil
}
