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

func (api *Api) Request(method, url string, body *struct{}) (*http.Response, error) {
	bodyJson, err := json.Marshal(body)
	if err != nil {
		fmt.Errorf("failed to marshal body: %w", err)
		return nil, err
	}

	req, err := http.NewRequest(method, url, bytes.NewReader(bodyJson))
	if err != nil {
		fmt.Errorf("client: could not create request: %s\n", err)
		return nil, err
	}

	res, err := api.client.Do(req)
	if err != nil {
		fmt.Errorf("client: could not send request: %s\n", err)
		return nil, err
	}

	return res, nil
}

func (api *Api) auth() (string, error) {

}