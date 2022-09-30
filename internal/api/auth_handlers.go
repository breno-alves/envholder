package api

import (
	"encoding/json"
	"io"
)

type LoginResponse struct {
	AccessToken string `json:"access_token"`
}

func (api *Api) Login(username, password string) (*LoginResponse, error) {
	// TODO: refactor
	body := make(map[string]interface{})
	body["username"] = username
	body["password"] = password

	res, err := api.Request("POST", "auth/login", "", body)
	if err != nil {
		return nil, err
	}

	bodyBytes, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	loginResponse := &LoginResponse{}
	err = json.Unmarshal(bodyBytes, loginResponse)
	if err != nil {
		return nil, err
	}

	return loginResponse, nil
}
