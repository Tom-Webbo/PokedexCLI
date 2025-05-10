package main

import (
	"fmt"
	"math/rand/v2"
)

func commandCatch(cfg *config, args ...string) error {
	if len(args) != 1 {
		return fmt.Errorf("please input the pokemons name")
	}
	name := args[0]
	pokemon, ok := cfg.pokeapiClient.GetPokemon(name)
	if ok != nil {
		return ok
	}
	fmt.Printf("Throwing a Pokeball at %s...", pokemon.Name)
	fmt.Println()
	var caught bool
	switch {
	case pokemon.BaseExperience <= 100:
		r := rand.Int64N(2)
		if r == 1 {
			caught = true
		} else {
			caught = false
		}
	case pokemon.BaseExperience >= 101:
		r := rand.Int64N(4)
		if r == 1 {
			caught = true
		} else {
			caught = false
		}
	}

	if !caught {
		fmt.Printf("%s escaped!", pokemon.Name)
		fmt.Println()
	} else {
		fmt.Printf("%s was caught!", pokemon.Name)
		fmt.Println()
		cfg.pokedex[pokemon.Name] = pokemon
	}

	return nil
}
