package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func repl() {
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

		err := commands.callback()
		if err != nil {
			fmt.Printf("Error: %v\n", err)
			continue
		}
	}

}

type cliCommand struct {
	name     string
	desc     string
	callback func() error
}

func getCommands() map[string]cliCommand {
	return map[string]cliCommand{
		"help": cliCommand{
			name:     "help",
			desc:     "Show available commands",
			callback: commandHelp,
		},
		"exit": cliCommand{
			name:     "exit",
			desc:     "Exit pokedex",
			callback: commandExit,
		},
	}
}

func cleanInput(text string) []string {
	text = strings.ToLower(text)

	return strings.Split(text, " ")
}
