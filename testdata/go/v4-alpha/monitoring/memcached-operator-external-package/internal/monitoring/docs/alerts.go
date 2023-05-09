package main

import (
	"fmt"

	"github.com/machadovilaca/operator-observability/pkg/docs"

	"github.com/example/memcached-operator/internal/monitoring/rules"
)

const alertsTpl = `# Memcached Operator Alerts

{{- range . }}

### {{.Name}}
**Summary:** {{ index .Annotations "summary" }}.

**Description:** {{ index .Annotations "description" }}.

**Severity:** {{ index .Labels "severity" }}.
{{- if .For }}

**For:** {{ .For }}.
{{- end -}}
{{- end }}

## Developing new alerts

All alerts documented here are auto-generated and reflect exactly what is being
exposed. After developing new alerts or changing old ones please regenerate
this document.
`

func printAlerts() {
	rules.SetupRules()
	docsString := docs.BuildAlertsDocsWithCustomTemplate(rules.ListAlerts(), alertsTpl)
	fmt.Print(docsString)
}
