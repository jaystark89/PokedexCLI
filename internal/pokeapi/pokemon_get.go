package pokeapi

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func (c *Client) PokemonGet(name string) (PokemonStats, error) {
	url := baseURL + "/pokemon/" + name

	if val, ok := c.cache.Get(url); ok {
		pokemonGet := PokemonStats{}
		err := json.Unmarshal(val, &pokemonGet)
		if err != nil {
			return PokemonStats{}, err
		}
		return pokemonGet, nil
	}

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return PokemonStats{}, err
	}

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return PokemonStats{}, err
	}
	defer resp.Body.Close()

	if resp.StatusCode > 299 {
		return PokemonStats{}, fmt.Errorf("Error: %d status code retrieved from PokeAPI", resp.StatusCode)
	}

	dat, err := io.ReadAll(resp.Body)
	if err != nil {
		return PokemonStats{}, err
	}

	pokemonResp := PokemonStats{}
	err = json.Unmarshal(dat, &pokemonResp)
	if err != nil {
		return PokemonStats{}, err
	}

	c.cache.Add(url, dat)
	return pokemonResp, nil
}
