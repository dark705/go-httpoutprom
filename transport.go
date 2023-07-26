package httpoutprom

import (
	"net/http"
	"strconv"
	"time"
)

type Transport struct {
	rt       http.RoundTripper
	recorder *Recorder
}

func NewTransport(recorder *Recorder, transport http.RoundTripper) *Transport {
	return &Transport{
		rt:       transport,
		recorder: recorder,
	}
}

func (t Transport) RoundTrip(request *http.Request) (*http.Response, error) {
	start := time.Now()
	response, err := t.rt.RoundTrip(request)
	if err == nil {
		t.recorder.ObserveDurationRequest(time.Since(start),
			request.URL.Host, request.URL.Scheme, request.Method, strconv.Itoa(response.StatusCode))
		t.recorder.IncRequest(request.URL.Host, request.URL.Scheme, request.Method, strconv.Itoa(response.StatusCode))
	}

	return response, err
}
