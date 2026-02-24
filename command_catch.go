package main

import "fmt"

func commandCatch(c *config, argument ...string) error {
	fmt.Println("lets catch")
	return nil
}