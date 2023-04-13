package metrics

import "github.com/example/memcached-operator/monitoring/metrics/registry"

const metricPrefix = "memcached_operator_"

var (
	metricsRegisters = [][]*registry.MetricDesc{
		operatorMetrics,
	}

	collectorRegister = []*registry.CollectorDesc{
		customResourceCollector,
	}
)

func SetupMetrics() {
	err := registry.SetupMetrics(metricsRegisters)
	if err != nil {
		panic(err)
	}

	err = registry.SetupCollectors(collectorRegister)
	if err != nil {
		panic(err)
	}
}

func ListMetrics() []*registry.MetricDesc {
	return registry.ListMetrics()
}
