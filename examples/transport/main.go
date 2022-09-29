package main

import (
	"fmt"
	"io"
	"net/http"

	"github.com/dark705/go-httpoutprom"
	"github.com/prometheus/client_golang/prometheus"
)

func main() {
	prometheus.MustRegister(httpoutprom.Collector())

	client := &http.Client{
		Transport: httpoutprom.NewTransport(http.DefaultTransport),
	}

	request, err := http.NewRequest(http.MethodGet, "https://httpbin.org/anything", nil)
	if err != nil {
		panic(err)
	}

	response, err := client.Do(request)
	if err != nil {
		panic(err)
	}
	defer response.Body.Close()

	body, err := io.ReadAll(response.Body)
	if err != nil {
		panic(err)
	}

	fmt.Printf("status: %d\nbody: %s\n", response.StatusCode, body) //nolint:forbidigo
}
