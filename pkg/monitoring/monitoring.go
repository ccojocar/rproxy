package monitoring

import (
	"fmt"
	"time"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
)

// RequestCountMetric defines the request count metric which records the number of requests served by the proxy
var RequestCountMetric = promauto.NewCounter(prometheus.CounterOpts{
	Name: "request_count",
	Help: "The number of requests served",
})

// requestLatencyMetric defines the request latency metric which records the latency of a request per service
type requestLatencyMetric struct {
	latencyHistogram *prometheus.HistogramVec
	started          time.Time
}

// RequestLatencyMeric defines a request latency metric
var RequestLatencyMeric *requestLatencyMetric

// Init initialises the monitoring metrics and register them with Prometheus
func Init() error {
	var err error
	if RequestLatencyMeric, err = newRequestLatencyMetric(); err != nil {
		return fmt.Errorf("initialising monitoring: %w", err)
	}
	return nil
}

//newRequestLatencyMetric creates a new request latency metric recorder.
func newRequestLatencyMetric() (*requestLatencyMetric, error) {
	rlm := &requestLatencyMetric{
		latencyHistogram: prometheus.NewHistogramVec(prometheus.HistogramOpts{
			Name:    "request_latency_seconds",
			Help:    "The request latency per downstream service",
			Buckets: prometheus.DefBuckets,
		}, []string{"service"}),
	}
	if err := prometheus.Register(rlm.latencyHistogram); err != nil {
		return nil, fmt.Errorf("registering request latency metric with Prometheus: %w", err)
	}
	return rlm, nil
}

// Start starts recording the request latency
func (r *requestLatencyMetric) Start() {
	// skip recording when the request latency recorder in not initialized
	if r == nil {
		return
	}
	r.started = time.Now()
}

// Finish finishes the request latency and send the metric to Prometheus
func (r *requestLatencyMetric) Finish(service string) {
	// skip recording when the request latency recorder in not initialized
	if r == nil {
		return
	}
	duration := time.Since(r.started).Seconds()
	r.latencyHistogram.WithLabelValues(service).Observe(duration)
}
