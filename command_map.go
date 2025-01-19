package main

import (
	"fmt"
)

type Area struct {
	Name string `json:"name"`
	Url  string `json:"url"`
}

type Areas struct {
	Count    int    `json:"count"`
	Next     string `json:"next"`
	Previous string `json:"previous"`
	Results  []Area `json:"results"`
}

func commandMapf(config *Config) error {
	areas, err := config.apiClient.GetLocationsList(config.nextLocationsURL, *config.cache)

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

func commandMapb(config *Config) error {

	if config.prevLocationsURL == nil {
		fmt.Println("you're on the first page")
		return nil
	}
	areas, err := config.apiClient.GetLocationsList(config.prevLocationsURL, *config.cache)

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
