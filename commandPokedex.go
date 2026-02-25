package main

import (
	"fmt"
)
//func(cfg *config, argument ...string) error
func commandPokedex(cfg *config, argument ...string) error {
	if len(cfg.pokedex) == 0 {
		fmt.Println("There are no pokemon in your pokedex!")
	}	
	fmt.Println("You have caught the following pokemon:")
	for _, pokemon := range cfg.pokedex {
		fmt.Println("\t-", pokemon.Name)
	}
	return nil
}
