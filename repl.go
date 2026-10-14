package main

import (
	"fmt"
	"os"
	"strings"

	"github.com/jaystark89/pokedexcli/internal/pokeapi"
)

func cleanInput(text string) []string {

	text = strings.ToLower(text)
	words := strings.Fields(text)

	return words
}

type cliCommand struct {
	name        string
	description string
	callback    func(_ *config, input ...string) error
}

type config struct {
	commandMap    map[string]cliCommand
	NextMap       *string
	PreviousMap   *string
	pokeapiClient pokeapi.Client
	userPokedex   map[string]pokeapi.PokemonStats
}

type MapResponse struct {
	LocationNames []Names `json:"results"`
	NextUrl       *string `json:"next"`
	PreviousUrl   *string `json:"previous"`
}

type Names struct {
	ActualNames string `json:"name"`
}

var commandMap = map[string]cliCommand{
	"exit": {
		name:        "exit",
		description: "Exit the Pokedex",
		callback:    commandExit,
	},
	"help": {
		name:        "help",
		description: "Displays helpful information",
		callback:    commandHelp,
	},
	"mapf": {
		name:        "mapf",
		description: "Displays names of the next 20 locations in the Pokemon world",
		callback:    commandMapf,
	},
	"mapb": {
		name:        "mapb",
		description: "Displays the previous 20 locations in the Pokemon world",
		callback:    commandMapb,
	},
	"explore": {
		name:        "explore",
		description: "Explores an area in the Pokemon world",
		callback:    commandExplore,
	},
	"catch": {
		name:        "catch",
		description: "Attempts to catch a Pokemon",
		callback:    commandCatch,
	},
	"pokedex": {
		name:        "pokedex",
		description: "Check the list of all your caught Pokemon",
		callback:    commandPokedex,
	},
	"inspect": {
		name:        "inspect",
		description: "Inspect a Pokemon for insight on their specific stat info",
		callback:    commandInspect,
	},
}

func commandExit(_ *config, _ ...string) error {
	fmt.Println("Closing the Pokedex... Goodbye!")
	os.Exit(0)
	return nil
}

func commandHelp(_ *config, _ ...string) error {
	fmt.Println("Welcome to the Pokedex!\nUsage:\n\nhelp: Displays a help message\nexplore <location>: lists the Pokemon found at the location.\nexit: Exit the Pokedex")
	return nil
}
