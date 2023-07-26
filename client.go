package httpoutprom

import (
	"net/http"
	"strconv"
	"time"
)

type HTTPClient interface {
	Do(*http.Request) (*http.Response, error)
}

type Client struct {
	client   HTTPClient
	recorder *Recorder
}

func NewClient(recorder *Recorder, client HTTPClient) *Client {
	return &Client{
		recorder: recorder,
		client:   client,
	}
}

func (c *Client) Do(request *http.Request) (*http.Response, error) {
	start := time.Now()
	response, err := c.client.Do(request)
	if err == nil {
		c.recorder.ObserveDurationRequest(time.Since(start),
			request.URL.Host, request.URL.Scheme, request.Method, strconv.Itoa(response.StatusCode))
		c.recorder.IncRequest(request.URL.Host, request.URL.Scheme, request.Method, strconv.Itoa(response.StatusCode))
	}

	return response, err
}
