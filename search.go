package main

import (
	"encoding/json"
	"io"
	"net/http"
	"os"
	"time"
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

func searchCall(url string) []*item {
	client := http.Client{
		Timeout: time.Second * 10,
	}

	resp, err := client.Get(url)
	if err != nil {
		Err("search failed with error, %v", err)
		os.Exit(1)
	}

	if resp.StatusCode != http.StatusOK {
		Err("search failed with error, %s", resp.Status)
		os.Exit(1)
	}

	defer resp.Body.Close()
	content, err := io.ReadAll(resp.Body)
	if err != nil {
		Err("read search result failed with error, %v", err)
		os.Exit(1)
	}

	items, err := parse(content)
	if err != nil {
		Err("parse search result failed with error, %v", err)
		os.Exit(1)
	}

	return items
}

func parse(body []byte) ([]*item, error) {
	resp := &searchRes{}
	err := json.Unmarshal(body, resp)
	if err != nil {
		return nil, err
	}
	return resp.Items, nil
}
