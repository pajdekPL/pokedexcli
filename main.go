package main

import (
	"time"

	"github.com/pajdekpl/pokedexcli/internal/pokeapi"
)

func main() {
	pokeClient := pokeapi.NewClient(5*time.Second, 8*time.Second)
	cfg := &Config{
		apiClient: pokeClient,
		pokemons:  make(map[string]pokeapi.PokemonInfo),
	}
	startRepl(*cfg)
}
