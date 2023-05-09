# Memcached Operator Metrics

### memcached_operator_cr_count [DEPRECATED in 1.14.0]
Number of existing memcached custom resources.

**Type:** Gauge.

### memcached_operator_reconcile_action_count [ALPHA]
Number of times the operator has executed the reconcile loop with a given action.

**Type:** Counter.

### memcached_operator_reconcile_count
Number of times the operator has executed the reconcile loop.

**Type:** Counter.

### memcached_operator_number_of_pods
Number of memcached operator pods in the cluster.

**Type:** Gauge.

### memcached_operator_number_of_ready_pods [ALPHA]
Number of ready memcached operator pods in the cluster.

**Type:** Gauge.

## Developing new metrics

All metrics documented here are auto-generated and reflect exactly what is being
exposed. After developing new metrics or changing old ones please regenerate
this document.
