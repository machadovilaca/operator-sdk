package main

import (
	"fmt"

	"github.com/machadovilaca/operator-observability/pkg/docs"

	"github.com/example/memcached-operator/internal/monitoring/metrics"
	"github.com/example/memcached-operator/internal/monitoring/rules"
)

const metricsTpl = `# Memcached Operator Metrics

{{- range . }}

{{ $deprecatedVersion := "" -}}
{{- with index .ExtraFields "DeprecatedVersion" -}}
    {{- $deprecatedVersion = printf " in %s" . -}}
{{- end -}}

{{- $stabilityLevel := "" -}}
{{- if and (.ExtraFields.StabilityLevel) (ne .ExtraFields.StabilityLevel "STABLE") -}}
	{{- $stabilityLevel = printf " [%s%s]" .ExtraFields.StabilityLevel $deprecatedVersion -}}
{{- end -}}

### {{ .Name }}{{ print $stabilityLevel }}
{{ .Help }}.

**Type:** {{ .Type -}}.
{{- end }}

## Developing new metrics

All metrics documented here are auto-generated and reflect exactly what is being
exposed. After developing new metrics or changing old ones please regenerate
this document.
`

func printMetrics() {
	metrics.SetupMetrics()
	rules.SetupRules()
	docsString := docs.BuildMetricsDocsWithCustomTemplate(metrics.ListMetrics(), rules.ListRecordingRules(), metricsTpl)
	fmt.Print(docsString)
}
