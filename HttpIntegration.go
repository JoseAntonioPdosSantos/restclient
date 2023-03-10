package restclient

import (
	"net/http"
)

type HttpIntegration interface {
	Url(url string) HttpIntegration
	Exec() HttpClientResponse
	Authorization(authorization Authorization) HttpIntegration
	ContentType(contentType ContentType) HttpIntegration
	Accept(accept ContentType) HttpIntegration
	AddHeader(key string, value string) HttpIntegration
	AddParams(key string, value string) HttpIntegration
	Interceptor(interceptor http.RoundTripper) HttpIntegration
	Body(body []byte) HttpIntegration
	AddQuery(key string, value string) HttpIntegration
}
