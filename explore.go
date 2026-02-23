package main

import "fmt"

func explore(cfg *config, args ...string) error {
	if len(args) != 1 {
		fmt.Println("Invalid location area name")
		return nil
	}
	url := args[0]
	locationResp, err := cfg.pokeApiClient.LocationAreaRequest(url)
	if err != nil {
		fmt.Println(err)
		return err
	}
	if locationResp.PokemonEncounters == nil {
		fmt.Println("There are no pokemon in this location")
		return nil
	}
	for _, pokemon := range	locationResp.PokemonEncounters {
		fmt.Println(pokemon.Pokemon.Name)
	}
	return nil
}