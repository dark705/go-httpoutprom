package httpoutprom

import (
	"net/http"
	"strconv"
	"time"

	"github.com/prometheus/client_golang/prometheus"
)

type HTTPClient interface {
	Do(*http.Request) (*http.Response, error)
}

type Client struct {
	client    HTTPClient
	counter   *prometheus.CounterVec
	histogram *prometheus.HistogramVec
}

func NewClient(client HTTPClient, registerer prometheus.Registerer) *Client {
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

	return &Client{
		client:    client,
		counter:   counter,
		histogram: histogram,
	}
}

func (c *Client) Do(request *http.Request) (*http.Response, error) {
	start := time.Now()
	response, err := c.client.Do(request)
	if err == nil {
		labels := prometheus.Labels{
			"host":   request.URL.Host,
			"scheme": request.URL.Scheme,
			"method": request.Method,
			"code":   strconv.Itoa(response.StatusCode),
		}

		c.counter.With(labels).Inc()
		c.histogram.With(labels).Observe(time.Since(start).Seconds())
	}

	return response, err
}
