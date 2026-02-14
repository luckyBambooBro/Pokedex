package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"github.com/luckybamboobro/pokedex/internal/pokeapi"
)


func commandMap(config *config) error {
	//set starting URL
	url := ""
	if config.next == "" {
		url = config.startURL
	} else {
		url = config.next 
	}

	//get http.Response from pokeAPI
	resp, err := http.Get(url)
	if err != nil {
		return fmt.Errorf("error obtaining http.get response from %s: %w", url, err)
	}
	//convert resp to io.Reader type so we can Unmarshal
	data, err := io.ReadAll(resp.Body)
	if err != nil {
		return fmt.Errorf("unable to obtain body of http.response: %w", err)
	}
	if resp.StatusCode > 299 {
		return fmt.Errorf("bad status code: %d", resp.StatusCode)
	}
	//safeguard close. can put another close call later if wanted
	defer resp.Body.Close()

	var apiResponse pokeapi.ApiStruct
	if err = json.Unmarshal(data, &apiResponse); err != nil {
		return fmt.Errorf("Error unmarshaling http response body: %w", err)
	}

	//print api response to terminal
	for _, item := range apiResponse.Results {
		fmt.Println(item.Name)
	}
	
	//update url to increment map page for next function call to call the next page of 20 results
	config.previous = apiResponse.Previous
	config.next = apiResponse.Next
	return nil
}

func commandMapB(config *config) error {
	//set starting URL
	url := ""
	if config.previous == "" {
		fmt.Println("you're on the first page")
		return nil		
	} else {
		url = config.previous 
	}

	//get http.Response from pokeAPI
	resp, err := http.Get(url)
	if err != nil {
		return fmt.Errorf("error obtaining http.get response from %s: %w", url, err)
	}
	//convert resp to io.Reader type so we can Unmarshal
	data, err := io.ReadAll(resp.Body)
	if err != nil {
		return fmt.Errorf("unable to obtain body of http.response: %w", err)
	}
	if resp.StatusCode > 299 {
		return fmt.Errorf("bad status code: %d", resp.StatusCode)
	}
	//safeguard close. can put another close call later if wanted
	defer resp.Body.Close()

	var apiResponse pokeapi.ApiStruct
	if err = json.Unmarshal(data, &apiResponse); err != nil {
		return fmt.Errorf("Error unmarshaling http response body: %w", err)
	}

	//print api response to terminal
	for _, item := range apiResponse.Results {
		fmt.Println(item.Name)
	}
	
	//update url to increment map page for next function call to call the next page of 20 results
	config.previous = apiResponse.Previous
	config.next = apiResponse.Next
	return nil
}
