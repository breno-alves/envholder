package api

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

type Api struct {
	client      *http.Client
	credentials *struct {
		bearerToken string
		apiKey      string
	}
}

func NewApi() *Api {
	return &Api{
		client: &http.Client{
			Timeout: 30 * time.Second,
		},
	}
}

type RequestOptions struct {
	method string
	url    string
	body   map[string]interface{}
	token  string
}

func (api *Api) Request(reqOptions *RequestOptions) (*http.Response, error) {
	var (
		body   []byte = nil
		method        = "GET"
		url           = URL + reqOptions.url
		token         = ""
		err    error
	)

	if reqOptions.body != nil {
		body, err = json.Marshal(reqOptions.body)
		if err != nil {
			return nil, err
		}
	}

	if reqOptions.token != "" {
		token = reqOptions.token
	}

	if reqOptions.method != "GET" {
		method = reqOptions.method
	}

	req, err := http.NewRequest(method, url, bytes.NewBuffer(body))
	if err != nil {
		fmt.Errorf("could not create request: %s\n", err)
		return nil, err
	}
	req.Header.Set("Content-Type", "application/json")

	if {
		req.Header.Set("Authorization", "Bearer "+token)
	}

	res, err := api.client.Do(req)
	if err != nil {
		fmt.Errorf("client: could not send request: %s\n", err)
		return nil, err
	}

	return res, nil
}
