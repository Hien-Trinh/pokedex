package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/Hien-Trinh/pokedex/internal/pokeapi"
)

type config struct {
	pokeapiClient   pokeapi.Client
	nextLocationURL *string
	prevLocationURL *string
	caughtPokemon   map[string]pokeapi.Pokemon
}

func repl(cfg *config) {
	reader := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print("Pokedex > ")
		reader.Scan()

		text := cleanInput(reader.Text())

		if len(text) == 0 {
			continue
		}

		commandName := text[0]

		commands, exist := getCommands()[commandName]
		if !exist {
			fmt.Printf("Pokedex: command not found: %s\n", commandName)
			continue
		}

		args := []string{}
		if len(text) > 1 {
			args = text[1:]
		}

		err := commands.callback(cfg, args...)
		if err != nil {
			fmt.Printf("Error: %v\n", err)
			continue
		}
	}

}

type cliCommand struct {
	name     string
	desc     string
	callback func(*config, ...string) error
}

func getCommands() map[string]cliCommand {
	return map[string]cliCommand{
		"help": {
			name:     "help",
			desc:     "Show available commands",
			callback: commandHelp,
		},
		"map": {
			name:     "map",
			desc:     "List next 20 locations",
			callback: commandMap,
		},
		"mapb": {
			name:     "mapb",
			desc:     "List previous 20 locations",
			callback: commandMapb,
		},
		"explore": {
			name:     "explore <location>",
			desc:     "Explore locations",
			callback: commandExplore,
		},
		"catch": {
			name:     "catch <pokemon>",
			desc:     "Catch a pokemon",
			callback: commandCatch,
		},
		"inspect": {
			name:     "inspect <pokemon>",
			desc:     "Inspect a pokemon",
			callback: commandInspect,
		},
		"pokedex": {
			name:     "pokedex",
			desc:     "Show list of caught pokemon",
			callback: commandPokedex,
		},
		"exit": {
			name:     "exit",
			desc:     "Exit pokedex",
			callback: commandExit,
		},
	}
}

func cleanInput(text string) []string {
	text = strings.ToLower(text)
	return strings.Fields(text)
}
