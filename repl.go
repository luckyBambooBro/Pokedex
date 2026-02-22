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

		//check for "explore" firstword
		if cleanedInput[0] == "explore" {
			explore(cfg, cleanedInput[1])
		} else {
			//if explore is not the first word, obtain first word only, then run one of the other commands
		commandName := cleanedInput[0]
		//if input is a legitimate command, call the callback function
		command, ok := getCommands()[commandName]
		if ok {
			err := command.callback(cfg, "")
			if err != nil {
				fmt.Println(err)
			}
		} else {
			fmt.Println("Unknown command")
			continue
		}
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
	callback    func(c *config, argument string) error
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
			callback:    explore,
		},
	}
}
