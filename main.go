package main

import (
	"fmt"
	"log"
	"time"

	"github.com/Hien-Trinh/pokedex/internal/pokeapi"
)

func main() {
	pokeClient := pokeapi.NewClient(5 * time.Second)

	resp, err := pokeClient.ListLocations()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(resp)
	// cfg := &config{
	// 	pokeapiClient: pokeClient,
	// }
	// repl(cfg)
}
