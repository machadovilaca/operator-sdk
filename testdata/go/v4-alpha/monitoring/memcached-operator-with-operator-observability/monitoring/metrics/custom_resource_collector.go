package metrics

import (
	"context"

	"sigs.k8s.io/controller-runtime/pkg/client"

	"github.com/machadovilaca/operator-observability/pkg/operatormetrics"

	cachev1alpha1 "github.com/example/memcached-operator/api/v1alpha1"
)

var (
	collectorK8sClient client.Client
)

func SetupCustomResourceCollector(k8sClient client.Client) {
	collectorK8sClient = k8sClient
}

var (
	customResourceCollector = operatormetrics.Collector{
		Metrics: []operatormetrics.CollectorMetric{
			{
				Metric: crCount,
				Labels: []string{"namespace"},
			},
		},
		CollectCallback: customResourceCollectorCallback,
	}

	crCount = operatormetrics.NewGauge(
		operatormetrics.MetricOpts{
			Name:        metricPrefix + "cr_count",
			Help:        "Number of existing guestbook custom resources",
			ConstLabels: map[string]string{"controller": "guestbook"},
			ExtraFields: map[string]string{
				"StabilityLevel":    "DEPRECATED",
				"DeprecatedVersion": "1.15.0",
			},
		},
	)
)

func customResourceCollectorCallback() []operatormetrics.CollectorResult {
	if collectorK8sClient == nil {
		return []operatormetrics.CollectorResult{}
	}

	result := &cachev1alpha1.MemcachedList{}
	err := collectorK8sClient.List(context.TODO(), result, client.InNamespace("default"))
	if err != nil {
		return []operatormetrics.CollectorResult{}
	}

	return []operatormetrics.CollectorResult{
		{
			Metric: crCount,
			Labels: []string{"default"},
			Value:  float64(len(result.Items)),
		},
	}
}
