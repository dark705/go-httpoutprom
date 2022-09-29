package httpoutprom

import (
	"net/http"
	"strconv"
	"time"

	"github.com/prometheus/client_golang/prometheus"
)

type Transport struct {
	rt http.RoundTripper
}

func NewTransport(transport http.RoundTripper) *Transport {
	return &Transport{
		rt: transport,
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

		col.counter.With(labels).Inc()
		col.histogram.With(labels).Observe(time.Since(start).Seconds())
	}

	return response, err
}
