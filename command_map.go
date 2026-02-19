package main

import "fmt"


func commandMap(cfg *config) error {
	locationResp, err := cfg.pokeApiClient.ClientRequest(cfg.next)
	if err != nil {
		return err
	}

	for _, item := range locationResp.Results {
		fmt.Println(item.Name)
	}
	cfg.next = locationResp.Next
	cfg.previous = locationResp.Previous 
	return nil
}

func commandMapb(cfg *config) error {
	if cfg.previous == "" {
		fmt.Println("you're on the first page")
		return nil
	}
	locationResp, err := cfg.pokeApiClient.ClientRequest(cfg.previous)
	if err != nil {
		return err
	}

	for _, item := range locationResp.Results {
		fmt.Println(item.Name)
	}
	cfg.next = locationResp.Next
	cfg.previous = locationResp.Previous 
	return nil
}
