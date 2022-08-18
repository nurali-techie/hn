package main

import (
	"encoding/json"
)

type searchRes struct {
	Items []*item `json:"hits"`
}

type item struct {
	Title    string `json:"title"`
	Url      string `json:"url"`
	Points   int    `json:"points"`
	ObjectID string `json:"objectID"`
}

func parse(body []byte) ([]*item, error) {
	resp := &searchRes{}
	err := json.Unmarshal(body, resp)
	if err != nil {
		return nil, err
	}
	return resp.Items, nil
}
