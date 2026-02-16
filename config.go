package main

import (
	"github.com/luckybamboobro/pokedex/internal/pokeapi"
	"time"
)
const 
(
	clientTimeout = 5 * time.Second
)

type config struct {
	pokeApiClient pokeapi.Client
	next     string
	previous string
}
