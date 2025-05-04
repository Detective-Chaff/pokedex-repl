package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	fmt.Println("Welcome to Pokedex!")
	fmt.Println("How can I help you?")
	// create scanner to read from command line
	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Printf("Pokedex > ")
		scanner.Scan()
		if scanner.Text() == "" {
			fmt.Println("Invalid input. Command cannot be blank")
		}

		words := cleanInput(scanner.Text())
		fmt.Println("Your command was:", words[0])
	}

}
