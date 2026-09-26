package main

import (
	"time"

	"github.com/pokemon/internal/pokeapi"
)

func main() {
	pokeAPI := pokeapi.NewClient(5 * time.Second)
	cfg := &config{
		commands:      getCommands(),
		pokeapiClient: pokeAPI,
	}
	startRepl(cfg)
}
