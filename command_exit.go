package main

import (
	"fmt"
	"os"
)

func commandExit(c *Config, args ...string) error {
	defer os.Exit(0)
	fmt.Println("Closing the Pokedex... Goodbye!")
	return nil
}
