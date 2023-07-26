package main

import (
	"fmt"
	"io"
	"net/http"

	"github.com/dark705/go-httpoutprom"
)

func main() {
	client := httpoutprom.NewClient(httpoutprom.NewRecorder(httpoutprom.Config{}), http.DefaultClient)

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

	fmt.Printf("status: %d; body: %s", response.StatusCode, body) //nolint:forbidigo
}
