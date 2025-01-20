package pokeapi

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
)

func (c *Client) GetPokemonList(areaName string) (ExploredLocation, error) {
	exploredLocation := ExploredLocation{}
	endpoint, err := url.JoinPath(baseURL, "/location-area/", areaName)

	if err != nil {
		return ExploredLocation{}, fmt.Errorf("error creating endpoint url: %v", err)
	}

	value, exists := c.cache.Get(endpoint)
	if exists {
		err := json.Unmarshal(value, &exploredLocation)
		if err != nil {
			return ExploredLocation{}, fmt.Errorf("error unmarshaling cached data: %v", err)
		}
		return exploredLocation, nil
	}

	req, err := http.NewRequest("GET", endpoint, nil)

	if err != nil {
		return ExploredLocation{}, fmt.Errorf("error creating request: %v", err)

	}

	res, err := c.httpClient.Do(req)

	if err != nil {
		return ExploredLocation{}, fmt.Errorf("error making request: %v", err)
	}

	defer res.Body.Close()
	data, err := io.ReadAll(res.Body)

	if err != nil {
		return ExploredLocation{}, fmt.Errorf("error reading request body: %v", err)

	}
	err = json.Unmarshal(data, &exploredLocation)

	if err != nil {
		return ExploredLocation{}, fmt.Errorf("error unmarshalling data: %v", err)
	}

	go c.cache.Add(endpoint, data)

	return exploredLocation, nil
}
