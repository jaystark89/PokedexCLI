package main

import (
	"errors"
	"fmt"
)

func commandInspect(cfg *config, args ...string) error {
	if len(args) != 1 {
		return errors.New("Error: Can only inspect one (1) Pokemon at a time!")
	}

	pokemon, ok := cfg.userPokedex[args[0]]
	if !ok {
		return errors.New("You haven't caught that Pokemon yet!")
	}

	fmt.Printf("Name: %s\n", pokemon.Name)
	fmt.Printf("Height: %d\n", pokemon.Height)
	fmt.Printf("Weight: %d\n", pokemon.Weight)
	fmt.Printf("Stats:\n")
	fmt.Printf("  -hp: %d\n", pokemon.Stats[0].BaseStat)
	fmt.Printf("  -attack: %d\n", pokemon.Stats[1].BaseStat)
	fmt.Printf("  -defense: %d\n", pokemon.Stats[2].BaseStat)
	fmt.Printf("  -special-attack: %d\n", pokemon.Stats[3].BaseStat)
	fmt.Printf("  -special-defense: %d\n", pokemon.Stats[4].BaseStat)
	fmt.Printf("  -speed: %d\n", pokemon.Stats[5].BaseStat)
	fmt.Printf("Types:\n")
	for _, typeInfo := range pokemon.Types {
		fmt.Printf("  - %s\n", typeInfo.Type.Name)
	}

	return nil

}
