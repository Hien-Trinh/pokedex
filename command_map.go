package main

import (
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
	return nil
}

func commandMapb(cfg *config) error {
	resp, err := cfg.pokeapiClient.ListLocations(cfg.prevLocationURL)
	if err != nil {
		return err
	}

	fmt.Println("Locations:")
	for _, location := range resp.Results {
		fmt.Printf("- %s\n", location.Name)
	}
	return nil
}
