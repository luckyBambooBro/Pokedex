package pokeapi

import (
	"encoding/json"
	"fmt"
	"json"
	"net/http"
)

//starting url for http request calls
const baseURL = "hhttps://pokeapi.co/api/v2/location-area/"

//create and return http.Get response from url provided
func (c *Client) ClientRequest(url string) (*http.Response, error) {
	//determine url for first or subsequent http request calls
	if url == "" {
		url = baseURL
	}

	//create request and send with client.Do()
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, fmt.Errorf("unable to create http request: %w", err)
	}
	resp, err := c.httpClient.Do(req)
	if err != nil {
				return nil, fmt.Errorf("http.Client.Do error, unable to create http request: %w, err")
	}
	
	defer resp.Body.Close()
	data, err := json.Unmarshal()
}

