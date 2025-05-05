package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
)

var registry = map[string]cliCommand{}

func main() {

	// Create command Registry
	registry = map[string]cliCommand{
		"exit": {
			name:        "exit",
			description: "Exit the Pokedex",
			callback:    commandExit,
		},
		"help": {
			name:        "help",
			description: "displays a help message",
			callback:    commandHelp,
		},
	}
	fmt.Println("How can I help you?")
	// create scanner to read from command line
	scanner := bufio.NewScanner(os.Stdin)

	for {
		fmt.Printf("Pokedex > ")
		scanner.Scan()
		if scanner.Text() == "" {
			fmt.Println("Invalid input. Command cannot be blank")
			continue
		}

		words := cleanInput(scanner.Text())
		// scan the command registry if no exists throw error and continue loop
		err := scanRegistry(words[0])
		if err != nil {
			fmt.Println(err)
			continue
		}

		switch words[0] {
		case "exit":
			registry["exit"].callback()
		case "help":
			registry["help"].callback()
		default:
			println("Unknown command")
		}

	}

}

func commandExit() error {
	fmt.Println("Closing the Pokedex... Goodbye!")
	os.Exit(0)
	return nil
}

func commandHelp() error {
	fmt.Println("Welcome to the Pokedex!")
	fmt.Println("Usage:")
	fmt.Println("")
	if registry == nil {
		return errors.New("no commands found. registry empty")
	}
	for _, command := range registry {
		fmt.Printf("%s: %s\n", command.name, command.description)
	}

	return nil
}

func scanRegistry(command string) error {
	if command == "" {
		return errors.New("command cannot be empty")
	}
	for k, _ := range registry {
		if command == k {
			return nil
		}
	}
	return errors.New("Unknown command")
}
