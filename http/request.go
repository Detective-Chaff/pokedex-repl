package http

import (
	"encoding/json"
	"net/http"
)

type locationResponse struct {
	Count    int    `json:"count"`
	Next     string `json:"next"`
	Previous any    `json:"previous"`
	Results  []struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	} `json:"results"`
}

type httpObj struct {
	req    *http.Request
	client *http.Client
	res    *http.Response
}

func createRequest(method string, url string) (*http.Request, error) {
	return http.NewRequest(method, url, nil)
}

func createClient() (*http.Client, error) {
	return &http.Client{}, nil
}

func GetLocations(direction string) (locationResponse, error) {
	var endpoint string
	var err error
	var locRes locationResponse
	if direction == "" {
		endpoint = "https://pokeapi.co/api/v2/location-area/"
	} else {
		endpoint = direction
	}

	obj := httpObj{}
	obj.req, err = createRequest("GET", endpoint)
	if err != nil {
		return locationResponse{}, err
	}

	obj.client, err = createClient()
	if err != nil {
		return locationResponse{}, err
	}

	obj.res, err = obj.client.Do(obj.req)
	if err != nil {
		return locationResponse{}, err
	}

	decoder := json.NewDecoder(obj.res.Body)
	err = decoder.Decode(&locRes)
	if err != nil {
		return locationResponse{}, err
	}

	return locRes, nil
}
