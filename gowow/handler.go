package gowow

import (
	"net/http"
	"time"
)

var httpClient = &http.Client{
	Timeout: time.Second * time.Duration(30),
}

type handler interface {
	Get(request *http.Request) (*http.Response, error)
}

type httpHandler struct{}

func (h httpHandler) Get(request *http.Request) (*http.Response, error) {
	response, err := httpClient.Do(request)
	if err != nil {
		return nil, err
	}
	return response, err
}
