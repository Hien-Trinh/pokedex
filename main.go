package main

import (
	"time"

	"github.com/Hien-Trinh/pokedex/internal/pokeapi"
)

func main() {
	pokeClient := pokeapi.NewClient(5 * time.Second)

	cfg := &config{
		pokeapiClient: pokeClient,
	}
	repl(cfg)
}
