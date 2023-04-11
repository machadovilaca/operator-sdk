package metrics

import (
	ctrl "sigs.k8s.io/controller-runtime"

	"github.com/machadovilaca/operator-observability/pkg/operatormetrics"
)

const metricPrefix = "memcached_operator_"

var metricsLog = ctrl.Log.WithName("metrics")

func SetupMetrics() {
	metricsLog.Info("registering metrics")
	err := operatormetrics.RegisterMetrics(operatorMetrics)
	if err != nil {
		panic(err)
	}

	metricsLog.Info("registering collectors")
	err = operatormetrics.RegisterCollector(customResourceCollector)
	if err != nil {
		panic(err)
	}
}

// ListMetrics returns a list of all metrics exposed by the operator
func ListMetrics() []operatormetrics.Metric {
	return operatormetrics.ListMetrics()
}
