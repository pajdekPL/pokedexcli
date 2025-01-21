package main

import (
	"fmt"
	"math/rand"
)

func commandCatch(c *Config, args ...string) error {
	if len(args) == 0 {
		return fmt.Errorf("you have to pass pokemon name to catch as a first parameter")
	}
	pokemonName := args[0]
	fmt.Printf("Throwing a Pokeball at %s...\n", pokemonName)

	pokemonInfo, err := c.apiClient.GetPokemonInfo(pokemonName)
	if err != nil {
		return fmt.Errorf("problem getting pokemon info: %v, err: %v", pokemonName, err)
	}

	if rand.Intn(pokemonInfo.BaseExperience) < 50 {
		fmt.Printf("%s was caught!\n", pokemonName)
		c.pokemons[pokemonName] = pokemonInfo
		return nil
	}
	fmt.Printf("%s escaped!\n", pokemonName)
	return nil
}
