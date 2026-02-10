package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	//func NewScanner(r io.Reader) *Scanner
	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print("Pokedex > ")
		if scanner.Scan() {
			userInput := scanner.Text()
			userInput = cleanInput(userInput)[0]
			found := false
			for k, v := range getCommands() {
				if userInput == k {
					found = true
					v.callback()
					break
				}
			}
			if !found {
				fmt.Println("unknown command")
			}
		}
	}
}
