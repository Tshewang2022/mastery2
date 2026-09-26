package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/pokemon/internal/pokeapi"
)

type cliCommand struct {
	config      *config
	name        string
	description string
	callback    func(cfg *config) error
}
type config struct {
	commands           map[string]cliCommand
	peviousLocationURL *string
	nextLocationURL    *string
	pokeapiClient      pokeapi.Client
}

func getCommands() map[string]cliCommand {
	return map[string]cliCommand{
		"help": {
			name:        "help",
			description: "Displays a help message",
			callback:    helpCommand,
		},
		"map": {
			name:        "map",
			description: "Get the next page of locations",
			callback:    commandMapf,
		},
		"mapb": {
			name:        "mapb",
			description: "Get the previous page of locations",
			callback:    commandMapb,
		},
		"exit": {
			name:        "exit",
			description: "Exit the Pokedex",
			callback:    exitCommand,
		},
		"explore": {
			name:        "explore",
			description: "List all the pokemen",
			callback:    findPoke,
		},
	}
}

func cleanInput(text string) []string {
	lowercase := strings.ToLower(text)
	return strings.Fields(lowercase)
}
func startRepl(cfg *config) {
	// what return type this returns
	scanner := bufio.NewScanner(os.Stdin)

	for {
		fmt.Print("Pokedex > ")
		if !scanner.Scan() {
			break
		}
		words := cleanInput(scanner.Text())

		if len(words) == 0 {
			continue
		}
		commandName := words[0]
		cmd, ok := cfg.commands[commandName]
		if !ok {
			fmt.Println("Unknown command")
			continue
		}

		if err := cmd.callback(cfg); err != nil {
			fmt.Println(err)
		}

	}
}
