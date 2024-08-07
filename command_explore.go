package main

import (
	"fmt"
)

func commandExplore(cfg *config, args ...string) error {
	if len(args) != 1 {
		return fmt.Errorf("expected 1 argument, got %d", len(args))
	}

	name := args[0]
	resp, err := cfg.pokeapiClient.GetLocation(name)
	if err != nil {
		return err
	}

	fmt.Printf("Exploring %s...\n", name)
	fmt.Println("Found the following Pokemon:")
	for _, pokemon_encounters := range resp.PokemonEncounters {
		fmt.Printf("- %s\n", pokemon_encounters.Pokemon.Name)
	}

	return nil
}
