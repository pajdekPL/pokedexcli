package pokeapi

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func (c *Client) GetLocationsList(pageUrl *string) (Areas, error) {
	areas := Areas{}
	url := baseURL + "/location-area"
	if pageUrl != nil {
		url = *pageUrl
	}

	value, exists := c.cache.Get(url)
	if exists {
		err := json.Unmarshal(value, &areas)
		if err != nil {
			return Areas{}, fmt.Errorf("error unmarshaling cached data: %v", err)
		}
		return areas, nil
	}

	req, err := http.NewRequest("GET", url, nil)

	if err != nil {
		return Areas{}, fmt.Errorf("error creating request: %v", err)
	}

	res, err := c.httpClient.Do(req)

	if err != nil {
		return Areas{}, fmt.Errorf("error making request: %v", err)
	}

	data, err := io.ReadAll(res.Body)
	if err != nil {
		return Areas{}, fmt.Errorf("error reading data from res.Body: %v", err)
	}
	go c.cache.Add(url, data)
	err = json.Unmarshal(data, &areas)

	if err != nil {
		return Areas{}, fmt.Errorf("error unMarshalling body: %v", err)
	}
	return areas, nil
}
