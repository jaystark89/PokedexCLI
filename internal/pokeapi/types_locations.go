package pokeapi

type RespShallowLocations struct {
	Results  []Location `json:"results"`
	Next     *string    `json:"next"`
	Previous *string    `json:"previous"`
}

type Location struct {
	Name string `json:"name"`
}

type RespDetailedLocation struct {
	Name              string              `json:"name"`
	PokemonEncounters []PokemonEncounters `json:"pokemon_encounters"`
}

type PokemonEncounters struct {
	Pokemon PokemonInfo `json:"pokemon"`
}

type PokemonInfo struct {
	Name string `json:"name"`
}

type PokemonStats struct {
	Id             int    `json:"id"`
	Name           string `json:"name"`
	BaseExperience int    `json:"base_experience"`
	Height         int    `json:"height"`
	Weight         int    `json:"weight"`
	Stats          []struct {
		BaseStat int `json:"base_stat"`
		Effort   int `json:"effort"`
		Stat     struct {
			Name string `json:"name"`
			URL  string `json:"url"`
		} `json:"stat"`
	} `json:"stats"`
	Types []struct {
		Slot int `json:"slot"`
		Type struct {
			Name string `json:"name"`
			URL  string `json:"url"`
		} `json:"type"`
	} `json:"types"`
}
