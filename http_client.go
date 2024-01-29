package restclient

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
	"time"
)

type HttpClient struct {
	url             string
	httpMethod      HttpMethod
	header          map[string]string
	params          map[string]string
	queries         map[string]string
	body            []byte
	bodyJson        any
	timeout         time.Duration
	timeoutDuration time.Duration
	authorization   Authorization
	interceptor     http.RoundTripper
}

func (h *HttpClient) Authorization(authorization Authorization) HttpIntegration {
	h.authorization = authorization
	return h
}

func (h *HttpClient) ContentType(contentType ContentType) HttpIntegration {
	h.header[ContentTypeDescription] = string(contentType)
	return h
}

func (h *HttpClient) Accept(accept ContentType) HttpIntegration {
	h.header[AcceptDescription] = string(accept)
	return h
}

func (h *HttpClient) AddHeader(key string, value string) HttpIntegration {
	h.header[key] = value
	return h
}

func (h *HttpClient) AddParams(key string, value string) HttpIntegration {
	h.params[key] = value
	return h
}

func (h *HttpClient) AddQuery(key string, value string) HttpIntegration {
	if len(value) > 0 {
		h.queries[key] = value
	}
	return h
}

func (h *HttpClient) Body(body []byte) HttpIntegration {
	h.body = body
	return h
}

func (h *HttpClient) Timeout(timeout time.Duration) HttpIntegration {
	h.timeout = timeout
	return h
}

func (h *HttpClient) BodyJson(body any) HttpIntegration {
	h.bodyJson = body
	return h
}

func (h *HttpClient) TimeoutDuration(timeoutDuration time.Duration) HttpIntegration {
	h.timeoutDuration = timeoutDuration
	return h
}

func (h *HttpClient) Interceptor(interceptor http.RoundTripper) HttpIntegration {
	h.interceptor = interceptor
	return h
}

func (h *HttpClient) Exec() HttpClientResponse {
	h.addParams()

	url := h.url + h.getRawQuery()

	err := h.makeBodyJson()

	if err != nil {
		return NewHttpRestClientResponse(nil, err)
	}

	req, err := http.NewRequest(string(h.httpMethod), url, bytes.NewBuffer(h.body))
	if err != nil {
		return NewHttpRestClientResponse(nil, err)
	}

	h.addHeaders(req)

	client := http.Client{
		Timeout:   h.getTimeout(),
		Transport: h.interceptor,
	}

	res, err := client.Do(req)
	if err != nil {
		return NewHttpRestClientResponse(nil, err)
	}
	return NewHttpRestClientResponse(res, err)
}

func (h *HttpClient) getTimeout() time.Duration {
	if h.timeout > 0 {
		if h.timeoutDuration > 0 {
			return h.timeout * h.timeoutDuration
		}
		return h.timeout * time.Second
	}
	if h.timeoutDuration > 0 {
		return 10 * h.timeoutDuration
	}
	return 10 * time.Second
}

func (h *HttpClient) makeBodyJson() error {
	if h.bodyJson != nil {
		b, err := json.Marshal(h.bodyJson)
		if err != nil {
			return err
		}
		h.body = b
	}
	return nil
}

func (h *HttpClient) addParams() {
	for k, v := range h.params {
		h.url = strings.Replace(h.url, fmt.Sprintf("${%s}", k), v, 3)
	}
}

func (h *HttpClient) addHeaders(request *http.Request) {
	for k, v := range h.header {
		request.Header.Add(k, v)
	}
	if h.authorization != nil {
		authorization := h.authorization.GetAuthorization()
		request.Header.Add(h.authorization.GetHeaderKey(), authorization)
	}
}

func (h *HttpClient) getRawQuery() string {
	queries := "?"
	for k, v := range h.queries {
		queries += fmt.Sprintf("%s=%s&", k, v)
	}
	return queries[:len(queries)-1]
}
