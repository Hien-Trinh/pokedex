package main

import (
	"fmt"
	"math/rand"
)

func commandCatch(cfg *config, args ...string) error {
	if len(args) != 1 {
		return fmt.Errorf("expected 1 argument, got %d", len(args))
	}

	name := args[0]
	resp, err := cfg.pokeapiClient.GetPokemon(name)
	if err != nil {
		return err
	}

	res := rand.Intn(resp.BaseExperience)
	threshold := 50

	fmt.Printf("Throwing a Pokeball at %s...\n", name)

	if res > threshold {
		fmt.Printf("%s escaped\n", name)
		return nil
	}

	fmt.Printf("%s was caught\n", name)
	cfg.caughtPokemon[name] = resp
	return nil
}
