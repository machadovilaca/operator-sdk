package metrics

import (
	k8smetrics "k8s.io/component-base/metrics"
)

var (
	operatorMetrics = []k8smetrics.Registerable{
		reconcileCount,
		reconcileAction,
	}

	reconcileCount = k8smetrics.NewCounter(
		&k8smetrics.CounterOpts{
			Name:           MetricPrefix + "reconcile_count",
			Help:           "Number of times the operator has executed the reconcile loop. Type: Counter.",
			StabilityLevel: k8smetrics.STABLE,
		},
	)

	reconcileAction = k8smetrics.NewCounterVec(
		&k8smetrics.CounterOpts{
			Name:           MetricPrefix + "reconcile_action_count",
			Help:           "Number of times the operator has executed the reconcile loop with a given action. Type: Counter.",
			StabilityLevel: k8smetrics.ALPHA,
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
