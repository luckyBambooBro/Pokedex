package main

import "github.com/luckybamboobro/pokedex/internal/pokeapi"

const 
(
	pokeLocationAreaUrl = "https://pokeapi.co/api/v2/location-area/"
	clientTimeout = 5
)

type config struct {
	pokeApiClient pokeapi.Client
	next     string
	previous string
}
