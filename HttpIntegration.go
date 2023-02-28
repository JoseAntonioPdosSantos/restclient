package restclient

import (
	"net/http"
)

type HttpIntegration interface {
	Url(url string) HttpIntegration
	Exec() ([]byte, error)
	Authorization(authorization Authorization) HttpIntegration
	AddHeader(key string, value string) HttpIntegration
	AddParams(key string, value string) HttpIntegration
	Interceptor(interceptor http.RoundTripper) HttpIntegration
	Body(body []byte) HttpIntegration
	AddQuery(key string, value string) HttpIntegration
}
