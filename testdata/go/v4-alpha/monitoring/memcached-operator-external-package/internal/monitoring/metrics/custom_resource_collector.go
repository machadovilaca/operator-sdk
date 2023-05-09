package metrics

import (
	"context"

	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"

	"github.com/machadovilaca/operator-observability/pkg/operatormetrics"

	cachev1alpha1 "github.com/example/memcached-operator/api/v1alpha1"
)

var (
	collectorK8sClient client.Client
	logger             = ctrl.Log.WithName("customResourceCollector")
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
			Name:        metricsPrefix + "cr_count",
			Help:        "Number of existing memcached custom resources",
			ConstLabels: map[string]string{"controller": "memcached"},
			ExtraFields: map[string]string{
				"StabilityLevel":    "DEPRECATED",
				"DeprecatedVersion": "1.14.0",
			},
		},
	)
)

func customResourceCollectorCallback() []operatormetrics.CollectorResult {
	if collectorK8sClient == nil {
		return []operatormetrics.CollectorResult{}
	}

	result := &cachev1alpha1.MemcachedList{}
	err := collectorK8sClient.List(context.TODO(), result, client.InNamespace(namespace))
	if err != nil {
		logger.Error(err, "Failed to list memcached custom resources")
		return []operatormetrics.CollectorResult{}
	}

	return []operatormetrics.CollectorResult{
		{
			Metric: crCount,
			Labels: []string{namespace},
			Value:  float64(len(result.Items)),
		},
	}
}
