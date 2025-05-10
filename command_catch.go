package main

import "fmt"

func commandCatch(cfg *config, args ...string) error {
	if len(args) != 1 {
		return fmt.Errorf("please input the pokemons name")
	}
	name := args[0]
	pokemon, ok := cfg.pokeapiClient.GetPokemon(name)
	if ok != nil {
		return ok
	}
	fmt.Println(pokemon.Name)
	fmt.Println(pokemon.BaseExperience)
	cfg.pokedex[pokemon.Name] = pokemon
	fmt.Println(cfg.pokedex)
	return nil
}
