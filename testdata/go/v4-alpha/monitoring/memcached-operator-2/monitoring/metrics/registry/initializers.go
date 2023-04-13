package registry

import (
	"fmt"

	"github.com/prometheus/client_golang/prometheus"
)

func NewCounter(opts *MetricDesc) (prometheus.Collector, error) {
	return prometheus.NewCounter(prometheus.CounterOpts{
		Name:        opts.FqName,
		Help:        opts.Help,
		ConstLabels: opts.ConstLabelPairs,
	}), nil
}

func NewCounterVec(opts *MetricDesc) (prometheus.Collector, error) {
	if len(opts.Labels) == 0 {
		return nil, fmt.Errorf("CounterVec labels not set for metric %v", opts.FqName)
	}

	return prometheus.NewCounterVec(prometheus.CounterOpts{
		Name:        opts.FqName,
		Help:        opts.Help,
		ConstLabels: opts.ConstLabelPairs,
	}, opts.Labels), nil
}

func NewGauge(opts *MetricDesc) (prometheus.Collector, error) {
	return prometheus.NewGauge(prometheus.GaugeOpts{
		Name:        opts.FqName,
		Help:        opts.Help,
		ConstLabels: opts.ConstLabelPairs,
	}), nil
}

func NewGaugeVec(opts *MetricDesc) (prometheus.Collector, error) {
	if len(opts.Labels) == 0 {
		return nil, fmt.Errorf("GaugeVec labels not set for metric %v", opts.FqName)
	}

	return prometheus.NewGaugeVec(prometheus.GaugeOpts{
		Name:        opts.FqName,
		Help:        opts.Help,
		ConstLabels: opts.ConstLabelPairs,
	}, opts.Labels), nil
}

func NewHistogram(opts *MetricDesc) (prometheus.Collector, error) {
	if opts.ExtraHistogramOpts == nil {
		return nil, fmt.Errorf("HistogramOpts not set for metric %v", opts.FqName)
	}

	opts.ExtraHistogramOpts.Name = opts.FqName
	opts.ExtraHistogramOpts.Help = opts.Help
	opts.ExtraHistogramOpts.ConstLabels = opts.ConstLabelPairs

	return prometheus.NewHistogram(*opts.ExtraHistogramOpts), nil
}

func NewHistogramVec(opts *MetricDesc) (prometheus.Collector, error) {
	if opts.ExtraHistogramOpts == nil {
		return nil, fmt.Errorf("HistogramOpts not set for metric %v", opts.FqName)
	}

	if len(opts.Labels) == 0 {
		return nil, fmt.Errorf("HistogramVec labels not set for metric %v", opts.FqName)
	}

	opts.ExtraHistogramOpts.Name = opts.FqName
	opts.ExtraHistogramOpts.Help = opts.Help
	opts.ExtraHistogramOpts.ConstLabels = opts.ConstLabelPairs

	return prometheus.NewHistogramVec(*opts.ExtraHistogramOpts, opts.Labels), nil
}

func NewSummary(opts *MetricDesc) (prometheus.Collector, error) {
	if opts.ExtraSummaryOpts == nil {
		return nil, fmt.Errorf("SummaryOpts not set for metric %v", opts.FqName)
	}

	opts.ExtraSummaryOpts.Name = opts.FqName
	opts.ExtraSummaryOpts.Help = opts.Help
	opts.ExtraSummaryOpts.ConstLabels = opts.ConstLabelPairs

	return prometheus.NewSummary(*opts.ExtraSummaryOpts), nil
}

func NewSummaryVec(opts *MetricDesc) (prometheus.Collector, error) {
	if opts.ExtraSummaryOpts == nil {
		return nil, fmt.Errorf("SummaryOpts not set for metric %v", opts.FqName)
	}

	if len(opts.Labels) == 0 {
		return nil, fmt.Errorf("SummaryVec labels not set for metric %v", opts.FqName)
	}

	opts.ExtraSummaryOpts.Name = opts.FqName
	opts.ExtraSummaryOpts.Help = opts.Help
	opts.ExtraSummaryOpts.ConstLabels = opts.ConstLabelPairs

	return prometheus.NewSummaryVec(*opts.ExtraSummaryOpts, opts.Labels), nil
}
