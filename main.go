package main

import (
	"bufio"
	"fmt"
	"os"
	"time"

	"github.com/jaystark89/pokedexcli/internal/pokeapi"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	pokeConfig := config{
		commandMap:    commandMap,
		pokeapiClient: pokeapi.NewClient(time.Second*5, time.Second*5),
		userPokedex:   make(map[string]pokeapi.PokemonStats),
	}

	for {
		fmt.Print("Pokedex > ")

		if !scanner.Scan() {
			break
		}

		poke_input := scanner.Text()
		clean_input := cleanInput(poke_input)

		if len(clean_input) == 0 {
			continue
		}

		if command, ok := pokeConfig.commandMap[clean_input[0]]; ok {
			args := clean_input[1:]
			err := command.callback(&pokeConfig, args...)
			if err != nil {
				fmt.Println("Error:", err)
			}
		}

	}
}
