package pokeapi

type Area struct {
	Name string `json:"name"`
	Url  string `json:"url"`
}

type Areas struct {
	Count    int     `json:"count"`
	Next     *string `json:"next"`
	Previous *string `json:"previous"`
	Results  []Area  `json:"results"`
}
