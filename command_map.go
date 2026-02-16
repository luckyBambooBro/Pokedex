package main

import "fmt"


func commandMap(c *config) error {
	locationResp, err := c.pokeApiClient.ClientRequest(c.next)
	if err != nil {
		return err
	}

	for _, item := range locationResp.Results {
		fmt.Println(item.Name)
	}
	c.next = locationResp.Next
	c.previous = locationResp.Previous 
	return nil
}

func commandMapb(c *config) error {
	if c.previous == "" {
		fmt.Println("you're on the first page")
		return nil
	}
	locationResp, err := c.pokeApiClient.ClientRequest(c.previous)
	if err != nil {
		return err
	}

	for _, item := range locationResp.Results {
		fmt.Println(item.Name)
	}
	c.next = locationResp.Next
	c.previous = locationResp.Previous 
	return nil
}
