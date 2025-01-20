package pokeapi

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func (c *Client) GetPokemonList(areaName string) (ExploredLocation, error) {
	exploredLocation := ExploredLocation{}
	url := baseURL + "/location-area/" + areaName

	value, exists := c.cache.Get(url)
	if exists {
		err := json.Unmarshal(value, &exploredLocation)
		if err != nil {
			return ExploredLocation{}, fmt.Errorf("error unmarshaling cached data: %v", err)
		}
		return exploredLocation, nil
	}

	req, err := http.NewRequest("GET", url, nil)

	if err != nil {
		return ExploredLocation{}, fmt.Errorf("error creating request %v", err)

	}

	res, err := c.httpClient.Do(req)

	if err != nil {
		return ExploredLocation{}, fmt.Errorf("error making request %v", err)
	}

	data, err := io.ReadAll(res.Body)

	if err != nil {
		return ExploredLocation{}, fmt.Errorf("error reading request body %v", err)

	}
	err = json.Unmarshal(data, &exploredLocation)

	if err != nil {
		return ExploredLocation{}, fmt.Errorf("error unmarshalling data %v", err)
	}

	go c.cache.Add(url, data)

	return exploredLocation, nil
}
