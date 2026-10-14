# PokedexCLI

My PokedexCLI program, built with guidance from [Boot.dev](https://www.boot.dev/)!

This program interfaces with the [PokeAPI](https://pokeapi.co/) to let you explore the different regions of the Pokémon games and discover which Pokémon inhabit them.

## Features

**Explore the world**
- `mapf` — Move forward through the locations.
- `mapb` — Move backward through the locations.
- `explore <location_name>` — Discover which Pokémon can be found in a location. Replace `<location_name>` with a location name shown by the map commands.

**Catch 'em all!**
- `catch <pokemon_name>` — Attempt to catch a Pokémon. Your chances of success depend on its base experience stat.
- `pokedex` — View the Pokémon you've caught so far.
- `inspect <pokemon_name>` — View detailed stats for a Pokémon you've caught. You have to catch it before you can inspect it!

## Getting Started

Want to try it out for yourself?

1. Clone this repository.
2. Navigate to the project root directory.
3. Build the program:

   ```bash
   go build ./...
   ```

4. Run the program:

   ```bash
   go run .
   ```

Thanks for checking out my GitHub! Stay tuned for more projects!
