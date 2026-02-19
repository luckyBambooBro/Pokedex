package pokeapi

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

//starting url for http request calls
const baseURL = "https://pokeapi.co/api/v2/location-area/"

//create and return http.Get response from url provided
func (c *Client) ClientRequest(url string) (ApiStruct, error) {
	//determine url for first or subsequent http request calls
	if url == "" {
		url = baseURL
	}

	//create request and send with client.Do()
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return ApiStruct{}, fmt.Errorf("unable to create http request: %w", err)
	}
	resp, err := c.httpClient.Do(req)
	if err != nil {
		return ApiStruct{}, fmt.Errorf("http.Client.Do error, unable to create http request: %w", err)
	}
	
	defer resp.Body.Close()
	data, err := io.ReadAll(resp.Body)
	if err != nil {
		return ApiStruct{}, fmt.Errorf("unable to read http response body: %w", err)
	}

	locationResp := ApiStruct{}
	if err = json.Unmarshal(data, &locationResp); err != nil {
		return ApiStruct{}, fmt.Errorf("unable to unmarshal http response body: %w", err)
	}

	return locationResp, nil
}

