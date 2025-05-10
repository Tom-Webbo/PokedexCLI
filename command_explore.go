package main

import "fmt"

func commandExplore(cfg *config, args ...string) error {
	if len(args) != 1 {
		return fmt.Errorf("please input name of location to explore")
	}

	location := args[0]
	fmt.Printf("Exploring %s", location)
	fmt.Println()
	locationsExplore, err := cfg.pokeapiClient.ExploreLocations(location)
	if err != nil {
		return err
	}
	for _, pokemonEncounter := range locationsExplore.PokemonEncounters {
		pokemeon := pokemonEncounter.Pokemon
		fmt.Println(pokemeon.Name)
	}
	return nil
}
