package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func startRepl(cfg *config) {
	scanner := bufio.NewScanner(os.Stdin)

	for {
		fmt.Print("Pokedex > ")
		if !scanner.Scan() {
			continue
		}

		userInput := scanner.Text()
		//refine the input and capture the first word
		firstWord := cleanInput(userInput)
		if len(firstWord) == 0 {
			continue
		}
		//obtain first word only
		commandName := firstWord[0]
		//if input is a legitimate command, call the callback function
		command, ok := getCommands()[commandName]
		if ok {
			err := command.callback(cfg)
			if err != nil {
				fmt.Println(err)
			}
		} else {
			fmt.Println("Unknown command")
			continue
		}
	}

}

// lowercases input and strips whitespace
func cleanInput(text string) []string {
	lowerCase := strings.ToLower(text)
	words := strings.Fields(lowerCase)
	return words
}

type cliCommand struct {
	name        string
	description string
	callback    func(*config) error
}

func getCommands() map[string]cliCommand {
	return map[string]cliCommand{
		"help": {
			name:        "help",
			description: "displays help menu for user",
			callback:    commandHelp,
		},
		"exit": {
			name:        "exit",
			description: "Exit the Pokedex",
			callback:    commandExit,
		},
		"map": {
			name:        "map",
			description: "List the location areas of the Pokeon World",
			callback:    commandMap,
		},
		"mapb": {
			name:        "mapb",
			description: "View previous page of the location areas of the Pokeon World",
			callback:    commandMapB,
		},
	}
}
