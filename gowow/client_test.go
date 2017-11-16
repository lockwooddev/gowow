package gowow

import (
	"bytes"
	"io/ioutil"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
)

type MockHandler struct {
	StatusCode int
	Headers    map[string]string
	Data       []byte
}

func (m MockHandler) Get(request *http.Request) (*http.Response, error) {
	body := ioutil.NopCloser(bytes.NewReader(m.Data))
	res := &http.Response{
		StatusCode: m.StatusCode,
		Body:       body,
		Header:     http.Header{},
	}

	for k, v := range m.Headers {
		res.Header.Set(k, v)
	}

	return res, nil
}

func TestNewAPIClientInstance(t *testing.T) {
	assert := assert.New(t)

	api := NewAPIClient("testkey")
	assert.Equal("testkey", api.apikey)
}

func TestFetchResourceNotOK(t *testing.T) {
	assert := assert.New(t)

	body := []byte(`{}`)
	mockHandler := MockHandler{StatusCode: 404, Data: body}
	api := apiClient{apikey: "testkey", handler: mockHandler}
	auctionDataStatus, err := api.GetAuctionDataStatus("eu", "khadgar", "en_GB", "callback")
	assert.Error(err)
	assert.Nil(auctionDataStatus)

	assert.Contains("Status code not ok: 404", err.Error())
}

func TestFetchResourceMasheryError(t *testing.T) {
	assert := assert.New(t)

	body := []byte(`{}`)

	headers := map[string]string{"X-Mashery-Error-Code": "ERR_403_DEVELOPER_INACTIVE"}
	mockHandler := MockHandler{StatusCode: 200, Data: body, Headers: headers}
	api := apiClient{apikey: "testkey", handler: mockHandler}
	auctionDataStatus, err := api.GetAuctionDataStatus("eu", "khadgar", "en_GB", "callback")
	assert.Error(err)
	assert.Nil(auctionDataStatus)
	assert.Contains("Invalid response: ERR_403_DEVELOPER_INACTIVE", err.Error())
}

func TestGetAuctionDataStatus(t *testing.T) {
	assert := assert.New(t)

	body := []byte(`
		{
			"files": [{
				"url": "http://auction-api-us.worldofwarcraft.com/auction-data/123/auctions.json",
				"lastModified": 1510507964000
			}]
		}
	`)

	mockHandler := MockHandler{StatusCode: 200, Data: body}
	api := apiClient{apikey: "testkey", handler: mockHandler}

	auctionDataStatus, err := api.GetAuctionDataStatus("eu", "khadgar", "en_GB", "callback")
	assert.NoError(err)

	file := AuctionFile{
		URL:          "http://auction-api-us.worldofwarcraft.com/auction-data/123/auctions.json",
		LastModified: 1510507964000,
	}

	expected := AuctionDataStatus{
		Files: []AuctionFile{
			file,
		},
	}

	assert.Equal(&expected, auctionDataStatus)
}
