package main

import (
	"net/http"
	"fmt"
	"io"
)

const pokeLocationAreaURL = "https://pokeapi.co/api/v2/location-area/"

func commandMap() error {
	resp, err := http.Get(pokeLocationAreaURL)
	if err != nil {
		return fmt.Errorf("error obtaining http.get response from %s: %w", pokeLocationAreaURL, err)
	}
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return fmt.Errorf("unable to obtain body of http.response: %w", err)
	}
	if resp.StatusCode > 299 {
		return fmt.Errorf("bad status code: %d", resp.StatusCode)
	}
	defer resp.Body.Close()
	return nil
}

