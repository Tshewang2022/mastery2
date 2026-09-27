package main

import (
	"errors"
	"fmt"

	"github.com/pokemon/internal/pokeapi"
)

func commandMapf(cfg *config) error {
	locationsResp, err := cfg.pokeapiClient.ListLocations(cfg.nextLocationURL)
	if err != nil {
		return err
	}

	cfg.nextLocationURL = locationsResp.Next
	cfg.peviousLocationURL = locationsResp.Previous

	for _, loc := range locationsResp.Results {
		fmt.Println(loc.Name)
	}
	return nil
}

func commandMapb(cfg *config) error {
	if cfg.peviousLocationURL == nil {
		return errors.New("you're on the first page")
	}

	locationResp, err := cfg.pokeapiClient.ListLocations(cfg.peviousLocationURL)
	if err != nil {
		return err
	}

	cfg.nextLocationURL = locationResp.Next
	cfg.peviousLocationURL = locationResp.Previous

	for _, loc := range locationResp.Results {
		fmt.Println(loc.Name)
	}
	return nil
}

// it takes this parameters;
func commandExplore(cfg *config) error {
	if cfg.arg == "" {
		return fmt.Errorf("you must provide a location area name")
	}
	location := cfg.arg

	area, err := cfg.pokeapiClient.GetLocationArea(location)
	if err != nil {
		return fmt.Errorf("error exploring the location: %w", err)
	}

	fmt.Println("Found Pokemon:")
	for _, encounter := range area.PokemonEncounters {
		fmt.Printf(" - %s\n", encounter.Pokemon.Name)
	}
	return nil
}

func catchCommand(cfg *config) error {

	if cfg.arg == "" {
		return fmt.Errorf("you must provide a pokemon name")

	}
	name := cfg.arg
	fmt.Printf("Throwing a Pokeball at %s...\n", name)

	pokemon, caught, err := cfg.pokeapiClient.CatchPokemonByName(name)

	if err != nil {
		return fmt.Errorf("error catching pokemon: %w", err)
	}

	if !caught {
		fmt.Printf("%s escaped!\n", name)
		return nil
	}
	fmt.Printf("%s was caught!\n", name)

	if cfg.pokedex == nil {
		cfg.pokedex = make(map[string]pokeapi.Pokemon)
	}
	cfg.pokedex[name] = pokemon
	return nil
}

func inspectCommand(cfg *config) error {
	if cfg.arg == "" {
		return fmt.Errorf("you must provide a pokemon name")
	}
	name := cfg.arg

	pokemon, ok := cfg.pokedex[name]
	if !ok {
		fmt.Println("you have not caught that pokemon")
		return nil
	}
	fmt.Printf("Name: %s\n", pokemon.Name)
	fmt.Printf("Height: %d\n", pokemon.Height)
	fmt.Printf("Weight: %d\n", pokemon.Weight)
	fmt.Println("Stats:")
	for _, stat := range pokemon.Stats {
		fmt.Printf("  -%s: %d\n", stat.Stat.Name, stat.BaseStat)
	}
	fmt.Println("Types:")
	for _, t := range pokemon.Types {
		fmt.Printf("  - %s\n", t.Type.Name)
	}

	return nil
}

func pokedexCommand(cfg *config) error {

	if len(cfg.pokedex) == 0 {
		fmt.Println("No pokemon has been caught")
		return nil
	}
	fmt.Println("Your Pokedex:")
	for _, val := range cfg.pokedex {
		fmt.Printf("Name: %s\n", val.Name)
	}
	return nil
}
