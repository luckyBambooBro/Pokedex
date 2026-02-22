package main

import "fmt"

func commandMap(cfg *config, argument string) error {
	//set default starting url if none provided
	if cfg.next == "" {
		cfg.next = baseURL
	}
	//send request
	locationResp, err := cfg.pokeApiClient.GeneralLocationAreaRequest(cfg.next)
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

func commandMapb(cfg *config, argument string) error {
	if cfg.previous == "" {
		fmt.Println("you're on the first page")
		return nil
	}
	locationResp, err := cfg.pokeApiClient.GeneralLocationAreaRequest(cfg.previous)
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
