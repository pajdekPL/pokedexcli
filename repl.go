package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/pajdekpl/pokedexcli/internal/pokeapi"
)

type Config struct {
	apiClient        pokeapi.Client
	pokemons         map[string]pokeapi.PokemonInfo
	prevLocationsURL *string
	nextLocationsURL *string
}

func startRepl(config Config) {
	scanner := bufio.NewScanner(os.Stdin)

	for {
		fmt.Print("Pokedex > ")
		if scanner.Scan() {
			if err := scanner.Err(); err != nil {
				fmt.Printf("Problem with input %s\n", err)
			}
			cmds := cleanInput(scanner.Text())
			if len(cmds) == 0 {
				fmt.Println("Please type command")
				continue
			}
			cmdRaw := cmds[0]
			cmd, existsCmd := getCommands()[cmdRaw]
			if existsCmd {
				err := cmd.callback(&config, cmds[1:]...)
				if err != nil {
					fmt.Println(err)
				}
				continue
			}
			fmt.Println("Unknown command")
		}
	}
}

// TODO add command arguments validator
type cliCommand struct {
	name        string
	description string
	callback    func(*Config, ...string) error
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
		"mapf": {
			name:        "mapf",
			description: "Displays next pokemon areas",
			callback:    commandMapf,
		},
		"mapb": {
			name:        "mapb",
			description: "Displays previous pokemon areas",
			callback:    commandMapb,
		},
		"explore": {
			name:        "explore",
			description: "Explore an area, usage 'explore <area_name>'",
			callback:    commandExplore,
		},
		"catch": {
			name:        "catch",
			description: "Try catching a given Pokemon 'explore <pokemon_name>'",
			callback:    commandCatch,
		},
		"inspect": {
			name:        "inspect",
			description: "Inspect a caught Pokemon 'inspect <pokemon_name>'",
			callback:    commandInspect,
		},
	}
}

func cleanInput(text string) []string {
	result := strings.Fields(strings.ToLower(text))

	return result
}
