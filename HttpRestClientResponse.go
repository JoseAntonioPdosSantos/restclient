package restclient

import (
	"io/ioutil"
	"net/http"
)

type HttpClientResponse interface {
	GetBody() ([]byte, error)
	GetResponse() *http.Response
	GetError() error
}

type HttpRestClientResponse struct {
	response *http.Response
	err      error
}

func NewHttpRestClientResponse(response *http.Response, err error) HttpClientResponse {
	return HttpRestClientResponse{response: response, err: err}
}

func (h HttpRestClientResponse) GetBody() (body []byte, err error) {
	defer h.response.Body.Close()

	body, err = ioutil.ReadAll(h.response.Body)
	if err != nil {
		return nil, err
	}
	return body, nil
}

func (h HttpRestClientResponse) GetResponse() *http.Response {
	return h.response
}

func (h HttpRestClientResponse) GetError() error {
	return h.err
}
