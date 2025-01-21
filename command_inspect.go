package main

import (
	"fmt"
)

func commandInspect(c *Config, args ...string) error {
	if len(args) == 0 {
		return fmt.Errorf("please provide a pokemon name")
	}
	pokemonName := args[0]
	pokemonInfo, exists := c.pokemons[pokemonName]

	if !exists {
		return fmt.Errorf("you can't inspect Pokemon that you didn't catch, trying: %s", pokemonName)
	}

	fmt.Println("Name: ", pokemonInfo.Name)
	fmt.Println("Height: ", pokemonInfo.Height)
	fmt.Println("Weight", pokemonInfo.Weight)
	fmt.Println("Stats:")
	for _, stat := range pokemonInfo.Stats {
		fmt.Printf("  -%v: %v\n", stat.Stat.Name, stat.BaseStat)
	}
	fmt.Println("Types:")
	for _, typex := range pokemonInfo.Types {
		fmt.Printf("  - %s\n", typex.Type.Name)
	}
	return nil
}
