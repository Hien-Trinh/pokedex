package main

import (
	"crypto/rand"
	"fmt"
	"math/big"
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

	res, err := rand.Int(rand.Reader, big.NewInt(int64(resp.BaseExperience)))
	if err != nil {
		return err
	}
	resInt := res.Int64()
	threshold := int64(50)

	fmt.Printf("Throwing a Pokeball at %s...\n", name)

	if resInt > threshold {
		fmt.Printf("%s escaped\n", name)
		return nil
	}

	fmt.Printf("%s was caught\n", name)
	cfg.caughtPokemon[name] = resp
	return nil
}
