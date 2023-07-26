# HTTP outgoing request Prometheus metrics

Provides simple wrappers around http.Client or http.RoundTripper for monitoring with Prometheus outgoing http requests.

## Metrics

Metrics:

| metric                                         | description                                      |
|------------------------------------------------|--------------------------------------------------|
| `http_outgoing_request_total`                  | A counter for outgoing http requests.            |
| `http_outgoing_request_duration_seconds`       | A histogram of outgoing http request latencies.  |

Default labels:

| label    | description                                    |
|----------|------------------------------------------------|
| `scheme` | Request scheme, such as: http, https.          |
| `host`   | Request host name.                             |
| `method` | Request method, such as: GET, POST, and so on. |
| `code`   | Response status code.                          |

## Examples

See example folder.

## Licensing

This project is licensed under the MIT License.