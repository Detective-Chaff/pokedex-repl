package main

import (
	"errors"
	"fmt"
	"os"

	"github.com/detective-chaff/pokedex-repl/http"
)

type cliCommand struct {
	name        string
	description string
	callback    func(c *config) error
}

func CommandExit(c *config) error {
	fmt.Println("Closing the Pokedex... Goodbye!")
	os.Exit(0)
	return nil
}

func CommandHelp(c *config) error {
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

func ScanRegistry(command string) error {
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

func CommandMap(c *config) error {
	data, err := http.GetLocations(c.Next)
	if err != nil {
		return err
	}

	c.Next = data.Next
	if data.Previous != nil {
		fmt.Println(data.Previous.(string))
		c.Previous = data.Previous.(string)
	}

	for _, v := range data.Results {
		fmt.Println(v.Name)
	}
	return nil
}

func CommandMapB(c *config) error {
	data, err := http.GetLocations(c.Previous)
	if err != nil {
		return err
	}

	c.Next = data.Next
	if data.Previous != nil {
		// fmt.Println(data.Previous.(string))
		c.Previous = data.Previous.(string)
	}

	for _, v := range data.Results {
		fmt.Println(v.Name)
	}
	return nil
}
