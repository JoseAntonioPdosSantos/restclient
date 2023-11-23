package restclient

import (
	"net/http"
	"time"
)

type HttpIntegration interface {
	ContentType(contentType ContentType) HttpIntegration
	Accept(accept ContentType) HttpIntegration
	Authorization(authorization Authorization) HttpIntegration
	Timeout(timeout time.Duration) HttpIntegration
	TimeoutDuration(timeoutDuration time.Duration) HttpIntegration
	AddHeader(key string, value string) HttpIntegration
	AddParams(key string, value string) HttpIntegration
	AddQuery(key string, value string) HttpIntegration
	AddQueryIfNoEmpty(key string, value string) HttpIntegration
	Interceptor(interceptor http.RoundTripper) HttpIntegration
	Body(body []byte) HttpIntegration
	BodyJson(body any) HttpIntegration
	Exec() HttpClientResponse
}
