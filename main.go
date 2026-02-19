package main

import "github.com/luckybamboobro/pokedex/internal/pokeapi"

func main() {
	cfg := config{
		pokeApiClient: pokeapi.NewClient(clientTimeout, cacheInterval),
		next: "",
		previous: "",
	}
	startRepl(&cfg)
}
