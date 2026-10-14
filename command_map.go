package main

import (
	"errors"
	"fmt"
)

func commandMapf(cfg *config, _ ...string) error {
	locationsResp, err := cfg.pokeapiClient.ListLocations(cfg.NextMap)
	if err != nil {
		return err
	}

	cfg.NextMap = locationsResp.Next
	cfg.PreviousMap = locationsResp.Previous

	for _, loc := range locationsResp.Results {
		fmt.Println(loc.Name)
	}
	return nil
}

func commandMapb(cfg *config, _ ...string) error {
	if cfg.PreviousMap == nil {
		return errors.New("you're on the first page")
	}

	locationsResp, err := cfg.pokeapiClient.ListLocations(cfg.PreviousMap)
	if err != nil {
		return err
	}

	cfg.NextMap = locationsResp.Next
	cfg.PreviousMap = locationsResp.Previous

	for _, loc := range locationsResp.Results {
		fmt.Println(loc.Name)
	}
	return nil
}
