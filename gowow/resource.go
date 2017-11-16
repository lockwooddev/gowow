package gowow

import (
	"net/http"
	"strings"
)

type resource struct {
	Region   string
	Endpoint string
	Params   []string
	Options  map[string]string
}

func (r resource) url() string {
	baseURL := regions[r.Region]
	endpoint := r.Endpoint
	argPart := strings.Join(r.Params, "/")
	return baseURL + endpoint + argPart
}

func (r resource) buildRequest() (*http.Request, error) {
	request, err := http.NewRequest("GET", r.url(), nil)
	if err != nil {
		return nil, err
	}

	qs := request.URL.Query()
	for k, v := range r.Options {
		qs.Add(k, v)
	}
	request.URL.RawQuery = qs.Encode()

	return request, err
}
