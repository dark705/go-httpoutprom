package httpoutprom

import (
	"net/http"
	"strconv"
	"time"

	"github.com/prometheus/client_golang/prometheus"
)

type Transport struct {
	rt        http.RoundTripper
	counter   *prometheus.CounterVec
	histogram *prometheus.HistogramVec
}

func NewTransport(transport http.RoundTripper, registerer prometheus.Registerer) *Transport {
	counter := prometheus.NewCounterVec(
		prometheus.CounterOpts{
			Name: "http_outgoing_request_total",
			Help: "Number of outgoing http request.",
		},
		[]string{
			"host", "scheme", "method", "code",
		},
	)

	histogram := prometheus.NewHistogramVec(
		prometheus.HistogramOpts{
			Name:    "http_outgoing_request_duration_seconds",
			Help:    "A histogram of outgoing http request latencies.",
			Buckets: prometheus.DefBuckets,
		},
		[]string{
			"host", "scheme", "method", "code",
		},
	)

	registerer.MustRegister(counter, histogram)

	return &Transport{
		rt:        transport,
		counter:   counter,
		histogram: histogram,
	}
}

func (t Transport) RoundTrip(request *http.Request) (*http.Response, error) {
	start := time.Now()
	response, err := t.rt.RoundTrip(request)
	if err == nil {
		labels := prometheus.Labels{
			"host":   request.URL.Host,
			"scheme": request.URL.Scheme,
			"method": request.Method,
			"code":   strconv.Itoa(response.StatusCode),
		}

		t.counter.With(labels).Inc()
		t.histogram.With(labels).Observe(time.Since(start).Seconds())
	}

	return response, err
}
