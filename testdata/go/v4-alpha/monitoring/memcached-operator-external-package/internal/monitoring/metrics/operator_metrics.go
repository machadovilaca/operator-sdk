package metrics

import "github.com/machadovilaca/operator-observability/pkg/operatormetrics"

var (
	operatorMetrics = []operatormetrics.Metric{
		reconcileCount,
		reconcileAction,
	}

	reconcileCount = operatormetrics.NewCounter(
		operatormetrics.MetricOpts{
			Name: metricsPrefix + "reconcile_count",
			Help: "Number of times the operator has executed the reconcile loop",
			ConstLabels: map[string]string{
				"controller": "memcached",
			},
			ExtraFields: map[string]string{
				"StabilityLevel": "STABLE",
			},
		},
	)

	reconcileAction = operatormetrics.NewCounterVec(
		operatormetrics.MetricOpts{
			Name: metricsPrefix + "reconcile_action_count",
			Help: "Number of times the operator has executed the reconcile loop with a given action",
			ExtraFields: map[string]string{
				"StabilityLevel": "ALPHA",
			},
		},
		[]string{"action"},
	)
)

func IncrementReconcileCountMetric() {
	reconcileCount.Inc()
}

func IncrementReconcileActionMetric(action string) {
	reconcileAction.WithLabelValues(action).Inc()
}
