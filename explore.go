package main

import "fmt"

func explore(cfg *config, url string) error {
	if url == "" {
		fmt.Println("Invalid url")
		return nil
	}
	url = baseURL + url
	locationResp, err := cfg.pokeApiClient.LocationAreaRequest(url)
	if err != nil {
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