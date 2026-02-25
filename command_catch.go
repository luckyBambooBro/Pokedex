package main

import (
	"fmt"
	"math/rand"

	"github.com/luckybamboobro/pokedex/internal/pokeapi"
)

func commandCatch(cfg *config, argument ...string) error {
	//check and set url
	if len(argument) == 0 {
		fmt.Println("Please provide a pokemon name")
		return nil
	}
	url_section := "pokemon/"
	pokemonName := argument[0]
	url := baseURL + url_section + pokemonName

	//send http request
	pokemon, err := cfg.pokeApiClient.RequestPokemon(url)
	if err != nil {
		fmt.Println(fmt.Errorf("Error obtaining pokemon data: %w", err))
		return nil
	}
	//if ok, then the following displays what happens when we attempt to catch the pokemon
	fmt.Printf("Throwing a Pokeball at %s...\n", pokemonName)
	catchThreshold := 30
	catchChance := rand.Int31n(int32(pokemon.BaseExperience))
	if catchChance <= int32(catchThreshold) {
		fmt.Println("Congratulations, you caught " + pokemonName + "!")
		saveToPokedex(cfg, pokemon)
	} else {
		fmt.Println(pokemonName + " escaped!")
	}
	return nil
}

func saveToPokedex(cfg *config, caughtPokemon *pokeapi.Pokemon) {
	_, ok := cfg.pokedex[caughtPokemon.Name]
	if !ok {
		cfg.pokedex[caughtPokemon.Name] = *caughtPokemon
	}
	fmt.Println("Pokemon you have caught in your Pokedex:")
	for _, pokemon := range cfg.pokedex {
		fmt.Println(pokemon.Name)
	}
}