# HTTP outgoing request Prometheus metrics

Provides simple wrappers around http.Client or http.RoundTripper for monitoring with Prometheus outgoing http requests.

## Metrics

Metrics:

| metric                                         | description                                      |
|------------------------------------------------|--------------------------------------------------|
| `http_outgoing_request_total`                  | A counter for outgoing http requests.            |
| `http_outgoing_request_duration_seconds`       | A histogram of outgoing http request latencies.  |

Labels:

| label    | description                                    |
|----------|------------------------------------------------|
| `scheme` | Request scheme, such as: http, https.          |
| `host`   | Request host name.                             |
| `method` | Request method, such as: GET, POST, and so on. |
| `code`   | Response status code.                          |

## Examples

### Wrap client:

``` go
client := httpoutprom.NewClient(http.DefaultClient, prometheus.DefaultRegisterer)

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
	
fmt.Printf("status: %d\nbody: %s\n", response.StatusCode, body)
```

### Wrap Transport:

``` go
client := &http.Client{
	Transport: httpoutprom.NewTransport(http.DefaultTransport, prometheus.DefaultRegisterer),
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
	
fmt.Printf("status: %d\nbody: %s\n", response.StatusCode, body)
```

## Licensing

This project is licensed under the MIT License.