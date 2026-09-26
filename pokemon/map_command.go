package main

import (
	"errors"
	"fmt"
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
func findPoke(cfg *config) error {

	return nil
}
