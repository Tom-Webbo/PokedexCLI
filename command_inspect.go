package main

import (
	"fmt"
)

func commandInspect(cfg *config, args ...string) error {
	if len(args) != 1 {
		return fmt.Errorf("please input pokemon name")
	}
	name := args[0]
	pokemon, ok := cfg.pokedex[name]
	if !ok {
		return fmt.Errorf("%s not found", name)
	}
	fmt.Printf(`
Name: %s
Height: %d
Weight: %d
`, pokemon.Name, pokemon.Height, pokemon.Weight)
	fmt.Println()
	fmt.Println("Stats:")
	for _, stat := range pokemon.Stats {
		fmt.Printf("-%s: %d", stat.Stat.Name, stat.BaseStat)
		fmt.Println()
	}
	fmt.Println("Types:")
	for _, ptype := range pokemon.Types {
		fmt.Printf(" - %s", ptype.Type.Name)
		fmt.Println()
	}

	return nil
}
