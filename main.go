package main

import "github.com/luckybamboobro/pokedex/internal/pokeapi"

const baseURL = "https://pokeapi.co/api/v2/location-area/"

func main() {
	cfg := config{
		pokeApiClient: pokeapi.NewClient(clientTimeout, cacheInterval),
		next: "",
		previous: "",
	}
	startRepl(&cfg)
}
