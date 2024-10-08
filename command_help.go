package main

import (
	"fmt"
)

func commandHelp(cfg *config, args ...string) error {
	fmt.Println()
	fmt.Println("Pokedex commands:")
	fmt.Println()
	for _, command := range getCommands() {
		fmt.Printf("%s : %s\n", command.name, command.desc)
	}
	fmt.Println()
	return nil
}
