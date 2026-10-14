package pokeapi

import (
	"encoding/json"
	"io"
	"net/http"
)

func (c *Client) LocationGet(name string) (RespDetailedLocation, error) {
	url := baseURL + "/location-area/" + name

	if val, ok := c.cache.Get(url); ok {
		locationGet := RespDetailedLocation{}
		err := json.Unmarshal(val, &locationGet)
		if err != nil {
			return RespDetailedLocation{}, err
		}
		return locationGet, nil
	}

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return RespDetailedLocation{}, err
	}

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return RespDetailedLocation{}, err
	}
	defer resp.Body.Close()

	dat, err := io.ReadAll(resp.Body)
	if err != nil {
		return RespDetailedLocation{}, err
	}

	locationResp := RespDetailedLocation{}
	err = json.Unmarshal(dat, &locationResp)
	if err != nil {
		return RespDetailedLocation{}, err
	}

	c.cache.Add(url, dat)
	return locationResp, nil
}
