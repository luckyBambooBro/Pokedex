package main

import "github.com/luckybamboobro/pokedex/internal/pokeapi"

func commandMap(c *config) error {
	url := ""
	if c.next == "" {
		url = pokeLocationAreaUrl
	} else {
		url = c.next
	} 
	resp, err := c.pokeApiClient.ClientRequest(c.next)
}

