package metrics

import (
	"github.com/prometheus/client_golang/prometheus"
	"github.com/sirupsen/logrus"
	"os"
	"strconv"
	"time"
)

// Metrics represents the configuration for configuring
// Prometheus metrics.
type Metrics struct {
	// Enabled determines if Metrics are enabled.
	// Determined by the `METRICS` environment label.
	Enabled bool
}

var (
	RequestsMetric = prometheus.NewCounterVec(prometheus.CounterOpts{
		Name: "nino_api_requests",
		Help: "How many HTTP requests processed, partitioned by its status code, method, and URI path.",
	}, []string{"status_code", "method", "path"})

	RequestLatencyMetric = prometheus.NewHistogramVec(prometheus.HistogramOpts{
		Name: "nino_api_request_latency",
		Help: "How much latency did the HTTP request take.",
	}, []string{"status_code", "method", "path"})
)

// NewMetrics initializes a new Metrics struct.
func NewMetrics() *Metrics {
	var enabled bool
	metrics := os.Getenv("METRICS")

	if metrics == "" {
		enabled = false
	} else {
		enabled = metrics == "true"
	}

	return &Metrics{
		Enabled: enabled,
	}
}

// Register registers all counters / gauge / histograms.
func (m *Metrics) Register() {
	logrus.WithField("type", "Prometheus").Info("Registering metrics...")
	prometheus.MustRegister(RequestsMetric, RequestLatencyMetric)
}

// IncRequest increments the RequestsMetric.
func (m *Metrics) IncRequest(status int, method string, path string) {
	RequestsMetric.WithLabelValues(strconv.Itoa(status), method, path).Inc()
}

// ObserveLatency observes the latency into the RequestLatencyMetric histogram.
func (m *Metrics) ObserveLatency(start time.Time, status int, method string, path string) {
	RequestLatencyMetric.WithLabelValues(strconv.Itoa(status), method, path).Observe(float64(time.Since(start).Nanoseconds() / 1000000))
}
