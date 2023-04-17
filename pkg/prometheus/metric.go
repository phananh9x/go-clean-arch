package prometheus

import (
	"context"

	"github.com/prometheus/client_golang/prometheus"
	ginprometheus "github.com/zsais/go-gin-prometheus"
)

// MetricPrometheus ...
type MetricPrometheus struct {
	metrics []*ginprometheus.Metric
}

// NewMetricPrometheus ...
func NewMetricPrometheus(metrics []*ginprometheus.Metric) IMerchantMetric {
	return &MetricPrometheus{
		metrics: metrics,
	}
}

// IncreaseRequestCnt ..
func (t *MetricPrometheus) IncreaseRequestCnt(ctx context.Context, name string, action string, code string) {
	mt := t.findMetricByName(name)
	if mt == nil {
		return
	}
	counterOps := (mt.MetricCollector).(*prometheus.CounterVec)
	counterOps.WithLabelValues(action, code).Inc()
}

// IncreasePartnerRequestCnt ..
func (t *MetricPrometheus) IncreasePartnerRequestCnt(ctx context.Context, name string, partnerName string, action string, code string) {
	mt := t.findMetricByName(name)
	if mt == nil {
		return
	}
	counterOps := (mt.MetricCollector).(*prometheus.CounterVec)
	counterOps.WithLabelValues(partnerName, action, code).Inc()
}

//RecordsCurrentValue ...
func (t *MetricPrometheus) RecordsCurrentValue(ctx context.Context, name string, val float64, lvb ...string) {
	metrics := t.findMetricByName(name)
	if metrics == nil {
		return
	}
	gaugeVec := (metrics.MetricCollector).(*prometheus.GaugeVec)
	gaugeVec.WithLabelValues(lvb...).Set(val)
}

//MesureRequestDuration ...
func (t *MetricPrometheus) MesureRequestDuration(ctx context.Context, name string, duration float64, lvb ...string) {
	metrics := t.findMetricByName(name)
	if metrics == nil {
		return
	}
	histogramVec := (metrics.MetricCollector).(*prometheus.HistogramVec)
	histogramVec.WithLabelValues(lvb...).Observe(duration)
}

func (t *MetricPrometheus) findMetricByName(name string) *ginprometheus.Metric {
	for i := 0; i < len(t.metrics); i++ {
		if t.metrics[i].Name == name {
			return t.metrics[i]
		}
	}
	return nil
}

// IncreaseWithNameCnt ...
func (t *MetricPrometheus) IncreaseWithNameCnt(ctx context.Context, name string, lvb ...string) {
	metrics := t.findMetricByName(name)
	if metrics == nil {
		return
	}
	counterVec := (metrics.MetricCollector).(*prometheus.CounterVec)
	counterVec.WithLabelValues(lvb...).Inc()
}
