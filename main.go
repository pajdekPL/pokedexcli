package main

import (
	"time"

	"github.com/pajdekpl/pokedexcli/internal/pokeapi"
	"github.com/pajdekpl/pokedexcli/internal/pokecache"
)

func main() {
	pokeClient := pokeapi.NewClient(5 * time.Second)
	cfg := &Config{
		apiClient: pokeClient,
		cache:     pokecache.NewCache(5 * time.Second),
	}
	startRepl(*cfg)
}
