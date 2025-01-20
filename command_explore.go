package main

import "fmt"

func commandExplore(c *Config, args ...string) error {
	if len(args) == 0 {
		return fmt.Errorf("you have to pass area_name to explore as a first parameter")
	}
	exploredArea, err := c.apiClient.GetPokemonList(args[0])

	if err != nil {
		return fmt.Errorf("problem exploring area: %v, err: %v", args[0], err)
	}

	for _, pokemon := range exploredArea.PokemonEncounters {
		fmt.Printf("- %s\n", pokemon.Pokemon.Name)
	}
	return nil
}
