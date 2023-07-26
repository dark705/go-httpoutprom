package httpoutprom

import (
	"time"

	"github.com/prometheus/client_golang/prometheus"
)

type Config struct {
	DurationBuckets []float64
	Registry        prometheus.Registerer
	Service         string
	ServiceLabel    string
	HostLabel       string
	SchemeLabel     string
	MethodLabel     string
	StatusCodeLabel string
}

func NewRecorder(cfg Config) *Recorder {
	if len(cfg.DurationBuckets) == 0 {
		cfg.DurationBuckets = prometheus.DefBuckets
	}

	if cfg.Registry == nil {
		cfg.Registry = prometheus.DefaultRegisterer
	}

	if cfg.ServiceLabel == "" {
		cfg.ServiceLabel = "service"
	}

	if cfg.HostLabel == "" {
		cfg.HostLabel = "host"
	}

	if cfg.SchemeLabel == "" {
		cfg.SchemeLabel = "scheme"
	}

	if cfg.MethodLabel == "" {
		cfg.MethodLabel = "method"
	}

	if cfg.StatusCodeLabel == "" {
		cfg.StatusCodeLabel = "code"
	}

	recorder := &Recorder{
		counter: prometheus.NewCounterVec(
			prometheus.CounterOpts{
				Name: "http_outgoing_request_total",
				Help: "Number of outgoing http request.",
			},
			[]string{cfg.ServiceLabel, cfg.HostLabel, cfg.SchemeLabel, cfg.MethodLabel, cfg.StatusCodeLabel},
		),
		histogram: prometheus.NewHistogramVec(
			prometheus.HistogramOpts{
				Name:    "http_outgoing_request_duration_seconds",
				Help:    "A histogram of outgoing http request latencies.",
				Buckets: cfg.DurationBuckets,
			},
			[]string{cfg.ServiceLabel, cfg.HostLabel, cfg.SchemeLabel, cfg.MethodLabel, cfg.StatusCodeLabel},
		),
		service: cfg.Service,
	}

	cfg.Registry.MustRegister(
		recorder.counter,
		recorder.histogram,
	)

	return recorder
}

type Recorder struct {
	service   string
	counter   *prometheus.CounterVec
	histogram *prometheus.HistogramVec
}

func (r *Recorder) Describe(in chan<- *prometheus.Desc) {
	r.histogram.Describe(in)
	r.counter.Describe(in)
}

func (r *Recorder) Collect(in chan<- prometheus.Metric) {
	r.histogram.Collect(in)
	r.counter.Collect(in)
}

func (r *Recorder) ObserveDurationRequest(duration time.Duration, host, scheme, method, statusCode string) {
	r.histogram.WithLabelValues(r.service, host, scheme, method, statusCode).Observe(duration.Seconds())
}

func (r *Recorder) IncRequest(host, scheme, method, statusCode string) {
	r.counter.WithLabelValues(r.service, host, scheme, method, statusCode).Inc()
}
