package restclient

import (
	"encoding/json"
	"io"
	"net/http"
)

type HttpClientResponse interface {
	GetBody() ([]byte, error)
	Unmarshal(response any) error
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

	body, err = io.ReadAll(h.response.Body)
	if err != nil {
		return nil, err
	}
	return body, nil
}

func (h HttpRestClientResponse) Unmarshal(response any) error {
	body, err := h.GetBody()

	if err != nil {
		return err
	}

	err = json.Unmarshal(body, &response)

	return err
}

func (h HttpRestClientResponse) GetResponse() *http.Response {
	return h.response
}

func (h HttpRestClientResponse) GetError() error {
	return h.err
}
