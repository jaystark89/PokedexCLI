package main

import (
	"fmt"
)

func commandPokedex(cfg *config, args ...string) error {
	fmt.Printf("Pokédex:\n")
	for _, caught := range cfg.userPokedex {
		fmt.Printf(" - %s\n", caught.Name)
	}
	return nil
}
