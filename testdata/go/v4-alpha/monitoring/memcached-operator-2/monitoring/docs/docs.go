package main

import (
	"bytes"
	"log"
	"sort"
	"strings"
	"text/template"

	dto "github.com/prometheus/client_model/go"

	"github.com/example/memcached-operator/monitoring/metrics"
	"github.com/example/memcached-operator/monitoring/metrics/registry"
)

const tpl = `# Operator Metrics

{{- range . }}

### {{.Name}}
[{{.StabilityLevel}}] {{.Help}}. Type: {{.Type}}.
{{- end }}

## Developing new metrics

All metrics documented here are auto-generated and reflect exactly what is being
exposed. After developing new metrics or changing old ones please regenerate
this document.
`

type metricDocs struct {
	Name           string
	Help           string
	Type           string
	StabilityLevel string
}

func main() {
	tpl, err := template.New("metrics").Parse(tpl)
	if err != nil {
		log.Fatalln(err)
	}

	metrics.SetupMetrics()
	ms := metrics.ListMetrics()

	buf := bytes.NewBufferString("")
	err = tpl.Execute(buf, buildMetricsDocs(ms))
	if err != nil {
		log.Fatalln(err)
	}

	log.Println(buf.String())
}

func buildMetricsDocs(metrics []*registry.MetricDesc) []metricDocs {
	metricsDocs := make([]metricDocs, len(metrics))
	for i, metric := range metrics {
		metricsDocs[i] = metricDocs{
			Name:           metric.FqName,
			Help:           metric.Help,
			Type:           getAndConvertMetricType(metric),
			StabilityLevel: string(metric.StabilityLevel),
		}
	}
	sortMetricsDocs(metricsDocs)

	return metricsDocs
}

func sortMetricsDocs(metricsDocs []metricDocs) {
	sort.Slice(metricsDocs, func(i, j int) bool {
		return metricsDocs[i].Name < metricsDocs[j].Name
	})
}

func getAndConvertMetricType(metric *registry.MetricDesc) string {
	t := dto.MetricType_name[int32(metric.MType)]
	return strings.ReplaceAll(t, "Vec", "")
}
