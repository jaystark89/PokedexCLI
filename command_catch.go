package main

import (
	"errors"
	"fmt"
	"math/rand"
)

func commandCatch(cfg *config, args ...string) error {
	if len(args) != 1 {
		return errors.New("Error: Can only catch one (1) Pokemon at a time!")
	}

	fmt.Printf("Throwing a Pokeball at %s...\n", args[0])
	pokemonResp, err := cfg.pokeapiClient.PokemonGet(args[0])
	if err != nil {
		return err
	}

	catch_chance := (100 - (pokemonResp.BaseExperience / 5))
	if catch_chance < 10 {
		catch_chance = 10
	}
	catch_roll := rand.Intn(100) + 1

	if catch_roll > catch_chance {
		fmt.Printf("Ah shucks! Looks like %s got away!\n", args[0])
	} else {
		fmt.Printf("Success! The Pokemon %s was caught and added to the Pokedex!\n", args[0])
		cfg.userPokedex[args[0]] = pokemonResp
	}

	return nil
}
