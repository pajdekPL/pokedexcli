package pokeapi

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
)

func (c *Client) GetPokemonInfo(pokemonName string) (PokemonInfo, error) {
	pokemonInfo := PokemonInfo{}
	endpoint, err := url.JoinPath(baseURL, "/pokemon/", pokemonName)

	if err != nil {
		return PokemonInfo{}, fmt.Errorf("error creating endpoint url: %v", err)
	}

	value, exists := c.cache.Get(endpoint)
	if exists {
		err := json.Unmarshal(value, &pokemonInfo)
		if err != nil {
			return PokemonInfo{}, fmt.Errorf("error unmarshaling cached data: %v", err)
		}
		return pokemonInfo, nil
	}

	req, err := http.NewRequest("GET", endpoint, nil)

	if err != nil {
		return PokemonInfo{}, fmt.Errorf("error creating request: %v", err)

	}

	res, err := c.httpClient.Do(req)

	if err != nil {
		return PokemonInfo{}, fmt.Errorf("error making request: %v", err)
	}

	defer res.Body.Close()
	data, err := io.ReadAll(res.Body)

	if err != nil {
		return PokemonInfo{}, fmt.Errorf("error reading request body: %v", err)

	}
	err = json.Unmarshal(data, &pokemonInfo)

	if err != nil {
		return PokemonInfo{}, fmt.Errorf("error unmarshalling data: %v", err)
	}

	go c.cache.Add(endpoint, data)

	return pokemonInfo, nil
}
