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

		err := commands.callback(cfg)
		if err != nil {
			fmt.Printf("Error: %v\n", err)
			continue
		}
	}

}

type cliCommand struct {
	name     string
	desc     string
	callback func(*config) error
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
			desc:     "List next 50 locations",
			callback: commandMap,
		},
		"mapb": {
			name:     "mapb",
			desc:     "List previous 50 locations",
			callback: commandMapb,
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
