package httpoutprom

import "github.com/prometheus/client_golang/prometheus"

var counter = prometheus.NewCounterVec(
	prometheus.CounterOpts{
		Name: "http_outgoing_request_total",
		Help: "Number of outgoing http request.",
	},
	[]string{
		"host", "scheme", "method", "code",
	},
)

var histogram = prometheus.NewHistogramVec(
	prometheus.HistogramOpts{
		Name:    "http_outgoing_request_duration_seconds",
		Help:    "A histogram of outgoing http request latencies.",
		Buckets: prometheus.DefBuckets,
	},
	[]string{
		"host", "scheme", "method", "code",
	},
)

func Collectors() []prometheus.Collector {
	return []prometheus.Collector{counter, histogram}
}
