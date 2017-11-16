package gowow

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRequestWithArguments(t *testing.T) {
	assert := assert.New(t)

	params := []string{"Khadgar"}
	options := map[string]string{"locale": "en_GB", "jsonp": "callback"}

	r := resource{Region: "eu", Endpoint: "wow/auction/data/", Params: params, Options: options}

	request, _ := r.buildRequest()

	expectedURL := "https://eu.api.battle.net/wow/auction/data/Khadgar?jsonp=callback&locale=en_GB"
	assert.Equal(expectedURL, request.URL.String())
}

func TestRequestEmptyArguments(t *testing.T) {
	assert := assert.New(t)

	params := []string{"Khadgar"}
	options := map[string]string{}

	r := resource{Region: "eu", Endpoint: "wow/auction/data/", Params: params, Options: options}

	request, _ := r.buildRequest()

	expectedURL := "https://eu.api.battle.net/wow/auction/data/Khadgar"
	assert.Equal(expectedURL, request.URL.String())
}
