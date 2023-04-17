package infra

import (
	ginprometheus "github.com/zsais/go-gin-prometheus"
	"go-clean-arch/pkg/prometheus"
)

func InitMetrics(metrics []*ginprometheus.Metric, subsystem string) prometheus.IMerchantMetric {
	counter := ginprometheus.NewPrometheus(subsystem, metrics)
	metricPro := prometheus.NewMetricPrometheus(counter.MetricsList)
	return metricPro
}
