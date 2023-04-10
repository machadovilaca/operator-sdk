package metrics

import (
	"context"

	k8smetrics "k8s.io/component-base/metrics"

	cachev1alpha1 "github.com/example/memcached-operator/api/v1alpha1"
)

type customResourceCollectorType struct {
	k8smetrics.BaseStableCollector

	customResourceCount *k8smetrics.Desc
}

func NewCustomResourceCollector() k8smetrics.StableCollector {
	return &customResourceCollectorType{
		customResourceCount: k8smetrics.NewDesc(
			MetricPrefix+"custom_resource_count",
			"Number of custom resources managed by the operator. Type: Gauge.",
			[]string{"namespace"},
			map[string]string{"cr_name": "memcached"},
			k8smetrics.ALPHA,
			"",
		),
	}
}

func (c *customResourceCollectorType) DescribeWithStability(ch chan<- *k8smetrics.Desc) {
	ch <- c.customResourceCount
}

func (c *customResourceCollectorType) CollectWithStability(ch chan<- k8smetrics.Metric) {
	metricsLog.Info("collecting custom resource metrics", "namespace", "default")

	result := &cachev1alpha1.MemcachedList{}

	if collectorK8sClient == nil {
		return
	}

	err := collectorK8sClient.List(context.TODO(), result)
	if err != nil {
		return
	}

	ch <- k8smetrics.NewLazyConstMetric(
		c.customResourceCount,
		k8smetrics.GaugeValue,
		float64(len(result.Items)),
		"default",
	)
}
