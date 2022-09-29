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
	client HTTPClient
}

func NewClient(client HTTPClient) *Client {
	return &Client{
		client: client,
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

		counter.With(labels).Inc()
		histogram.With(labels).Observe(time.Since(start).Seconds())
	}

	return response, err
}
