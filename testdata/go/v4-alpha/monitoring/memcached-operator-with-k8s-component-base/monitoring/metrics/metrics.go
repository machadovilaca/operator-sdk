package metrics

import (
	k8smetrics "k8s.io/component-base/metrics"
	"k8s.io/component-base/metrics/legacyregistry"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

const MetricPrefix = "memcached_operator_"

var metricsLog = ctrl.Log.WithName("metrics")

var (
	MetricRegisters = [][]k8smetrics.Registerable{
		operatorMetrics,
	}

	CollectorRegisters = []k8smetrics.StableCollector{
		NewCustomResourceCollector(),
	}

	collectorK8sClient client.Client
)

func init() {
	// Metrics
	metricsLog.Info("registering metrics")
	for _, register := range MetricRegisters {
		legacyregistry.MustRegister(register...)
	}

	// Collectors
	metricsLog.Info("registering collectors")
	for _, register := range CollectorRegisters {
		legacyregistry.CustomMustRegister(register)
	}
}

// SetupCustomResourceCollector sets the k8s client to be used by the custom resource collector
func SetupCustomResourceCollector(k8sClient client.Client) {
	collectorK8sClient = k8sClient
}
