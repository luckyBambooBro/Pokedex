package main

import (
	"fmt"
	"github.com/luckybamboobro/pokedex/internal/pokeapi"
)

func commandInspect(cfg *config, argument ...string) error {
	//check if pokemon has been caught to proceed
	pokemon := argument[0]
	//if the pokemon has not been caught
	pokemonData, ok := cfg.pokedex[pokemon];
		if !ok {
			return fmt.Errorf("This pokemon is not in your pokedex! You must catch ", pokemon, " to view data")
		}
	//if the pokemon has been caught
	printPokemonData(&pokemonData)
	return nil	
}

func printPokemonData(pokemon *pokeapi.Pokemon) {
	fmt.Println("Name: ", pokemon.Name)
	fmt.Println("Height: ", pokemon.Height)
	fmt.Println("Weight: ", pokemon.Weight)
	fmt.Println("Stats")
	for _, stat := range pokemon.Stats {
		fmt.Printf("\t-%s: %d\n", stat.Stat.Name, stat.BaseStat)
	}
	fmt.Println("Types:")
	for _, poketype := range pokemon.Types {
		fmt.Println("\t-", poketype.Type.Name)
	}
}