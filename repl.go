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
		cleanedInput := cleanInput(userInput)
		if len(cleanedInput) == 0 {
			continue
		}
		
		//if explore is not the first word, obtain first word only, then run one of the other commands
		commandName := cleanedInput[0]
		args := []string{}
		if len(cleanedInput) > 1 {
			args = cleanedInput[1:]
		}
		//if input is a legitimate command, call the callback function
		command, ok := getCommands()[commandName]
		if ok {
			err := command.callback(cfg, args...)
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
	callback    func(cfg *config, argument ...string) error
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
			description: "View the previous list the location areas of the Pokeon World",
			callback:    commandMapb,
		},
		"explore": {
			name:        "explore",
			description: "View an explored location area via name/ID url",
			callback:    commandExplore,
		},
		"catch": {
			name:        "catch",
			description: "Attempt to catch [pokemon name]",
			callback:    commandCatch,
		},
		"inspect": {
			name:        "inspect",
			description: "View pokemon data from pokedex",
			callback:    commandInspect,
		},
	}
}
