package metrics

import (
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
)

// Metrics are being used to monitor the service status.
type Metrics struct {
	requests    *prometheus.CounterVec
	failedCalls *prometheus.CounterVec
	latency     *prometheus.HistogramVec
}

// NewMetrics returns a metrics struct with registered promehteus metrics.
func NewMetrics() *Metrics {
	return &Metrics{
		requests: promauto.NewCounterVec(prometheus.CounterOpts{
			Name: "total_requests",
			Help: "total number of requests",
		}, []string{"endpoint"}),
		failedCalls: promauto.NewCounterVec(prometheus.CounterOpts{
			Name: "failed_calls",
			Help: "total number of failed requests",
		}, []string{"endpoint"}),
		latency: promauto.NewHistogramVec(prometheus.HistogramOpts{
			Name: "requests_latency",
			Help: "requests handling latency",
		}, []string{"endpoint"}),
	}
}

func (m *Metrics) AddRequest(endpoint string) {
	m.requests.With(prometheus.Labels{"endpoint": endpoint}).Add(1)
}

func (m *Metrics) AddFailedCall(endpoint string) {
	m.failedCalls.With(prometheus.Labels{"endpoint": endpoint}).Add(1)
}

func (m *Metrics) ObserveLatency(endpoint string, value float64) {
	m.latency.With(prometheus.Labels{"endpoint": endpoint}).Observe(value)
}
