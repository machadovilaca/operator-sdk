package rules

import (
	"github.com/machadovilaca/operator-observability/pkg/operatorrules"
	promv1 "github.com/prometheus-operator/prometheus-operator/pkg/apis/monitoring/v1"
)

const (
	recordingRulesPrefix = "memcached_operator_"
	namespace            = "memcached-operator"
)

var (
	// Add your custom recording rules here
	recordingRules = [][]operatorrules.RecordingRule{
		operatorRecordingRules,
	}

	// Add your custom alerts here
	alerts = [][]promv1.Rule{
		operatorAlerts,
	}
)

func SetupRules() {
	err := operatorrules.RegisterRecordingRules(recordingRules...)
	if err != nil {
		panic(err)
	}

	err = operatorrules.RegisterAlerts(alerts...)
	if err != nil {
		panic(err)
	}
}

func ListRecordingRules() []operatorrules.RecordingRule {
	return operatorrules.ListRecordingRules()
}

func ListAlerts() []promv1.Rule {
	return operatorrules.ListAlerts()
}
