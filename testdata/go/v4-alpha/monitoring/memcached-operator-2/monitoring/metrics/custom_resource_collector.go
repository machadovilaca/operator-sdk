package metrics

import (
	"context"

	"sigs.k8s.io/controller-runtime/pkg/client"

	"github.com/example/memcached-operator/monitoring/metrics/registry"
	"github.com/prometheus/client_golang/prometheus"
	dto "github.com/prometheus/client_model/go"

	cachev1alpha1 "github.com/example/memcached-operator/api/v1alpha1"
)

const customResourceCount = metricPrefix + "custom_resource_count"

var (
	collectorK8sClient client.Client
)

func SetupCustomResourceCollector(k8sClient client.Client) {
	collectorK8sClient = k8sClient
}

var (
	collectorMetrics = map[string]*registry.MetricDesc{
		customResourceCount: {
			FqName:         customResourceCount,
			Help:           "Number of custom resources",
			MType:          dto.MetricType_GAUGE,
			StabilityLevel: registry.STABLE,
			Labels:         []string{"namespace"},
		},
	}

	customResourceCollector = registry.NewCollector(
		collectorMetrics,
		customResourceCollectorFunc,
	)
)

func customResourceCollectorFunc(ch chan<- prometheus.Metric) {
	if collectorK8sClient == nil {
		return
	}

	result := &cachev1alpha1.MemcachedList{}
	err := collectorK8sClient.List(context.TODO(), result, client.InNamespace("default"))

	value, err := registry.NewMetricWithValue(collectorMetrics[customResourceCount], float64(len(result.Items)), "default")
	if err != nil {
		panic(err)
	}

	ch <- value
}
