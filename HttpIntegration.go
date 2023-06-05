package restclient

import (
	"net/http"
	"time"
)

type HttpIntegration interface {
	Url(url string) HttpIntegration
	Host(url string) HttpIntegration
	Exec() HttpClientResponse
	Authorization(authorization Authorization) HttpIntegration
	ContentType(contentType ContentType) HttpIntegration
	Accept(accept ContentType) HttpIntegration
	AddHeader(key string, value string) HttpIntegration
	AddParams(key string, value string) HttpIntegration
	Interceptor(interceptor http.RoundTripper) HttpIntegration
	Body(body []byte) HttpIntegration
	Timeout(timeout time.Duration) HttpIntegration
	TimeoutDuration(timeoutDuration time.Duration) HttpIntegration
	AddQuery(key string, value string) HttpIntegration
}
