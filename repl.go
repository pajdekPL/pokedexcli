package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/pajdekpl/pokedexcli/internal/pokeapi"
	"github.com/pajdekpl/pokedexcli/internal/pokecache"
)

type Config struct {
	apiClient        pokeapi.Client
	prevLocationsURL *string
	nextLocationsURL *string
	cache            *pokecache.Cache
}

func startRepl(config Config) {
	scanner := bufio.NewScanner(os.Stdin)

	for {
		fmt.Print("Pokedex > ")
		if scanner.Scan() {
			cmds := cleanInput(scanner.Text())
			if len(cmds) == 0 {
				fmt.Println("Please type command")
				continue
			}
			cmdRaw := cmds[0]
			cmd, existsCmd := getCommands()[cmdRaw]
			if existsCmd {
				err := cmd.callback(&config)
				if err != nil {
					fmt.Println(err)
				}
				continue
			}
			fmt.Println("Unknown command")
		}

		if err := scanner.Err(); err != nil {
			fmt.Printf("Problem with input %s\n", err)
		}
	}
}

type cliCommand struct {
	name        string
	description string
	callback    func(*Config) error
}

func getCommands() map[string]cliCommand {
	return map[string]cliCommand{
		"exit": {
			name:        "exit",
			description: "Exit the Pokedex",
			callback:    commandExit,
		},
		"help": {
			name:        "help",
			description: "Displays a help message",
			callback:    commandHelp,
		},
		"map": {
			name:        "map",
			description: "Displays next pokemon areas",
			callback:    commandMapf,
		},
		"mapb": {
			name:        "map",
			description: "Displays previous pokemon areas",
			callback:    commandMapb,
		},
	}
}

func cleanInput(text string) []string {
	result := strings.Fields(strings.ToLower(text))

	return result
}
