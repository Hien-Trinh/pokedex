package main

import (
	"errors"
	"fmt"
)

func commandMap(cfg *config) error {
	resp, err := cfg.pokeapiClient.ListLocations(cfg.nextLocationURL)
	if err != nil {
		return err
	}

	fmt.Println("Locations:")
	for _, location := range resp.Results {
		fmt.Printf("- %s\n", location.Name)
	}

	cfg.nextLocationURL = resp.Next
	cfg.prevLocationURL = resp.Previous
	return nil
}

func commandMapb(cfg *config) error {
	if cfg.prevLocationURL == nil {
		return errors.New("You are already at the beginning of the list")
	}

	resp, err := cfg.pokeapiClient.ListLocations(nil)
	if err != nil {
		return err
	}

	fmt.Println("Locations:")
	for _, location := range resp.Results {
		fmt.Printf("- %s\n", location.Name)
	}

	cfg.nextLocationURL = resp.Next
	cfg.prevLocationURL = resp.Previous
	return nil
}
