package main

import "github.com/luckybamboobro/pokedex/internal/pokeapi"

const baseURL = "https://pokeapi.co/api/v2/"

func main() {
	cfg := config{
		pokeApiClient: pokeapi.NewClient(clientTimeout, cacheInterval),
		next: "",
		previous: "",
		pokedex: map[string]pokeapi.Pokemon{},
	}
	startRepl(&cfg)
}
