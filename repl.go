package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

var registry = map[string]cliCommand{}

type config struct {
	Next     string
	Previous string
}

func Start() {
	//Initialize Config
	config := config{
		Next:     "",
		Previous: "",
	}

	// Create command Registry
	registry = map[string]cliCommand{
		"exit": {
			name:        "exit",
			description: "Exit the Pokedex",
			callback:    CommandExit,
		},
		"help": {
			name:        "help",
			description: "displays a help message",
			callback:    CommandHelp,
		},
		"map": {
			name:        "map",
			description: "get location areas of the pokemon world",
			callback:    CommandMap,
		},
		"mapb": {
			name:        "mapb",
			description: "get previous page's location area data",
			callback:    CommandMapB,
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
		err := ScanRegistry(words[0])
		if err != nil {
			fmt.Println(err)
			continue
		}

		switch words[0] {
		case "exit":
			registry["exit"].callback(&config)
		case "help":
			registry["help"].callback(&config)
		case "map":
			err := registry["map"].callback(&config)
			if err != nil {
				fmt.Println("error fetching results:", err)
			}
		case "mapb":
			err := registry["mapb"].callback(&config)
			if err != nil {
				fmt.Println("error fetching previous results:", err)
			}
		default:
			fmt.Println("Unknown command")
		}

	}
}

func cleanInput(text string) []string {
	words := strings.Fields(text)
	strs := []string{}
	for _, v := range words {
		strs = append(strs, strings.ToLower(v))
	}
	return strs
}
