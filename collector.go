package httpoutprom

import "github.com/prometheus/client_golang/prometheus"

func Collector() *collector { //nolint:golint,revive
	return col
}

type collector struct {
	counter   *prometheus.CounterVec
	histogram *prometheus.HistogramVec
}

func (c *collector) Describe(in chan<- *prometheus.Desc) {
	c.histogram.Describe(in)
	c.counter.Describe(in)
}

func (c *collector) Collect(in chan<- prometheus.Metric) {
	c.histogram.Collect(in)
	c.counter.Collect(in)
}

var col = &collector{
	counter: prometheus.NewCounterVec(
		prometheus.CounterOpts{
			Name: "http_outgoing_request_total",
			Help: "Number of outgoing http request.",
		},
		[]string{
			"host", "scheme", "method", "code",
		},
	),
	histogram: prometheus.NewHistogramVec(
		prometheus.HistogramOpts{
			Name:    "http_outgoing_request_duration_seconds",
			Help:    "A histogram of outgoing http request latencies.",
			Buckets: prometheus.DefBuckets,
		},
		[]string{
			"host", "scheme", "method", "code",
		},
	),
}
