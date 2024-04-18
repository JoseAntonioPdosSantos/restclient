package restclient

import (
	"encoding/json"
	"errors"
	"fmt"
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
	return &HttpRestClientResponse{response: response, err: err}
}

func (h *HttpRestClientResponse) GetBody() (body []byte, err error) {
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			fmt.Printf("an error occurred while trying to close the body, got: %s", err)
		}
	}(h.response.Body)
	if h.response == nil {
		return nil, errors.New("response can not be nil")
	}
	if h.response.Body == nil {
		return nil, errors.New("body is nil")
	}
	body, err = io.ReadAll(h.response.Body)
	if err != nil {
		return nil, err
	}
	return body, nil
}

func (h *HttpRestClientResponse) Unmarshal(objectToConvert any) error {
	body, err := h.GetBody()

	if err != nil {
		return err
	}

	err = json.Unmarshal(body, &objectToConvert)

	return err
}

func (h *HttpRestClientResponse) GetResponse() *http.Response {
	return h.response
}

func (h *HttpRestClientResponse) GetError() error {
	return h.err
}
