package restclient

import (
	"bytes"
	"net/http"
	"strings"
	"time"
)

type HttpClient struct {
	url           string
	httpMethod    HttpMethod
	header        map[string]string
	params        map[string]string
	queries       map[string]string
	body          []byte
	authorization Authorization
	interceptor   http.RoundTripper
}

func (h HttpClient) Url(url string) HttpIntegration {
	h.url = url
	return h
}

func (h HttpClient) Authorization(authorization Authorization) HttpIntegration {
	h.authorization = authorization
	return h
}

func (h HttpClient) ContentType(contentType ContentType) HttpIntegration {
	h.header[ContentTypeDescription] = string(contentType)
	return h
}

func (h HttpClient) Accept(accept ContentType) HttpIntegration {
	h.header[AcceptDescription] = string(accept)
	return h
}

func (h HttpClient) AddHeader(key string, value string) HttpIntegration {
	h.header[key] = value
	return h
}

func (h HttpClient) AddParams(key string, value string) HttpIntegration {
	h.params[key] = value
	return h
}

func (h HttpClient) AddQuery(key string, value string) HttpIntegration {
	h.queries[key] = value
	return h
}

func (h HttpClient) Body(body []byte) HttpIntegration {
	h.body = body
	return h
}

func (h HttpClient) Interceptor(interceptor http.RoundTripper) HttpIntegration {
	h.interceptor = interceptor
	return h
}

func (h HttpClient) Exec() HttpClientResponse {
	h.addParams()

	url := h.url + h.getRawQuery()

	req, err := http.NewRequest(string(h.httpMethod), url, bytes.NewBuffer(h.body))
	if err != nil {
		return NewHttpRestClientResponse(nil, err)
	}

	h.addHeaders(req)

	client := http.Client{
		Timeout:   10 * time.Second,
		Transport: h.interceptor,
	}

	res, err := client.Do(req)
	if err != nil {
		return NewHttpRestClientResponse(nil, err)
	}
	return NewHttpRestClientResponse(res, err)
}

func (h *HttpClient) addParams() {
	for k, v := range h.params {
		h.url = strings.Replace(h.url, "${"+k+"}", v, 3)
	}
}

func (h HttpClient) addHeaders(request *http.Request) {
	for k, v := range h.header {
		request.Header.Add(k, v)
	}
	if h.authorization != nil {
		authorization := h.authorization.GetAuthorization()
		request.Header.Add("Authorization", authorization)
	}
}

func (h HttpClient) getRawQuery() string {
	queries := "?"
	for k, v := range h.queries {
		queries += k + "=" + v + "&"
	}
	return queries[:len(queries)-1]
}
