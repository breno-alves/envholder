package api

import (
	"encoding/json"
	"io"
)

type ListProjectsResponse struct {
	Projects []struct {
		ID          string `json:"id"`
		Name        string `json:"name"`
		Description string `json:"description"`
		isActive    bool   `json:"isActive"`
		createdAt   string `json:"createdAt"`
		updatedAt   string `json:"updatedAt"`
	}
}

func (api *Api) ListProjects(acessToken string) (*ListProjectsResponse, error) {
	res, err := api.Request("GET", "projects", acessToken, nil)
	if err != nil {
		return nil, err
	}

	bodyBytes, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	listProjectsResponse := &ListProjectsResponse{}
	err = json.Unmarshal(bodyBytes, listProjectsResponse)
	if err != nil {
		return nil, err
	}

	return listProjectsResponse, nil
}
