package main

import (
	"fmt"
	"os"
	"strings"
)

func cleanInput(text string) []string {
	lowerCase := strings.ToLower(text)
	words := strings.Fields(lowerCase)
	return words
}

func exitCommand() error {
	fmt.Println("Closing the Pokedex... Goodbye!")
	os.Exit(0)
	return nil
}

func helpCommand() error {
	fmt.Printf(`Welcome to the Pokedex!
Usage: 

`)
	for _, v := range commands {
		fmt.Printf("%s: %s\n", v.name, v.description)
	}
	return nil
}

type cliCommand struct {
	name        string
	description string
	callback    func() error
}

var commands = map[string]cliCommand{
	"help": {
		name:        "help",
		description: "displays help menu for user",
		callback:    helpCommand,
	},
	"exit": {
		name:        "exit",
		description: "Exit the Pokedex",
		callback:    exitCommand,
	},
}
