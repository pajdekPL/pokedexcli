package main

import (
	"fmt"
)

func commandPokedex(c *Config, args ...string) error {
	fmt.Println("Your Pokedex:")
	for _, pokemon := range c.pokemons {
		fmt.Printf("  - %s\n", pokemon.Name)
	}
	return nil
}
