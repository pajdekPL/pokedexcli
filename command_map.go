package main

import (
	"fmt"
)

func commandMapf(config *Config, args ...string) error {
	areas, err := config.apiClient.GetLocationsList(config.nextLocationsURL)

	if err != nil {
		return err
	}

	config.nextLocationsURL = areas.Next
	config.prevLocationsURL = areas.Previous

	for _, area := range areas.Results {
		fmt.Println(area.Name)
	}

	return nil
}

func commandMapb(config *Config, args ...string) error {

	if config.prevLocationsURL == nil {
		fmt.Println("you're on the first page")
		return nil
	}
	areas, err := config.apiClient.GetLocationsList(config.prevLocationsURL)

	if err != nil {
		return err
	}

	config.nextLocationsURL = areas.Next
	config.prevLocationsURL = areas.Previous

	for _, area := range areas.Results {
		fmt.Println(area.Name)
	}

	return nil
}
