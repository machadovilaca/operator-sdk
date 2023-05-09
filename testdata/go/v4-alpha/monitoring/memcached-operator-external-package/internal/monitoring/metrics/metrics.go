package metrics

import "github.com/machadovilaca/operator-observability/pkg/operatormetrics"

const (
	metricsPrefix = "memcached_operator_"
	namespace     = "memcached-operator"
)

var (
	// Add your custom metrics here
	metrics = [][]operatormetrics.Metric{
		operatorMetrics,
	}

	// Add your custom collectors here
	collectors = []operatormetrics.Collector{
		customResourceCollector,
	}
)

func SetupMetrics() {
	err := operatormetrics.RegisterMetrics(metrics...)
	if err != nil {
		panic(err)
	}

	err = operatormetrics.RegisterCollector(collectors...)
	if err != nil {
		panic(err)
	}
}

func ListMetrics() []operatormetrics.Metric {
	return operatormetrics.ListMetrics()
}
