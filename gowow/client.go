package gowow

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
)

type apiClient struct {
	handler handler
	apikey  string
}

func NewAPIClient(APIKey string) apiClient {
	return apiClient{
		handler: httpHandler{},
		apikey:  APIKey,
	}
}

func (a apiClient) fetchResource(r resource) (*http.Response, error) {
	request, err := r.buildRequest()
	if err != nil {
		return nil, err
	}

	// set api key
	qs := request.URL.Query()
	qs.Add("apikey", a.apikey)
	request.URL.RawQuery = qs.Encode()

	response, err := a.handler.Get(request)
	if err != nil {
		return nil, err
	}

	if response.StatusCode != 200 {
		msg := fmt.Sprintf("Status code not ok: %d", response.StatusCode)
		return nil, errors.New(msg)
	}

	apiError := response.Header.Get("X-Mashery-Error-Code")
	if apiError != "" {
		msg := fmt.Sprintf("Invalid response: %s", apiError)
		return nil, errors.New(msg)
	}

	return response, nil
}

func (a apiClient) GetAuctionDataStatus(region string, realm string, locale string, jsonp string) (*AuctionDataStatus, error) {
	params := []string{realm}
	options := map[string]string{"locale": locale, "jsonp": jsonp}

	r := resource{
		Region:   region,
		Endpoint: "wow/auction/data/",
		Params:   params,
		Options:  options,
	}

	response, err := a.fetchResource(r)
	if err != nil {
		return nil, err
	}

	defer response.Body.Close()

	auctionDataStatus := &AuctionDataStatus{}
	err = json.NewDecoder(response.Body).Decode(auctionDataStatus)
	if err != nil {
		return nil, err
	}

	return auctionDataStatus, nil
}

func (a apiClient) GetRealmStatus(region string, locale string, jsonp string) (*RealmStatus, error) {
	options := map[string]string{"locale": locale, "jsonp": jsonp}

	r := resource{
		Region:   region,
		Endpoint: "wow/realm/status",
		Options:  options,
	}

	response, err := a.fetchResource(r)
	if err != nil {
		return nil, err
	}

	defer response.Body.Close()

	realmStatus := &RealmStatus{}
	err = json.NewDecoder(response.Body).Decode(realmStatus)
	if err != nil {
		return nil, err
	}

	return realmStatus, nil
}
