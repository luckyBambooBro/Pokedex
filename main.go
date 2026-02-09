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
			fmt.Printf("Your command was: %s \n", userInput)
		}
	}
}



