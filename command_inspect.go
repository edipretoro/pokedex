package main

import (
	"fmt"
)

func commandInspect(cfg *config, name *string) error {
	if name == nil || *name == "" {
		return fmt.Errorf("please provide a Pokemon name to inspect")
	}

	pokemon, ok := cfg.pokedex[*name]
	if !ok {
		fmt.Println("you have not caught that pokemon")
	}
	fmt.Println("Name:", pokemon.Name)
	fmt.Println("Height:", pokemon.Height)
	fmt.Println("Weight:", pokemon.Weight)
	fmt.Println("Stats:")
	for s := range pokemon.Stats {
		fmt.Printf("  -%s: %d\n", pokemon.Stats[s].Stat.Name, pokemon.Stats[s].BaseStat)
	}
	fmt.Println("Types:")
	for t := range pokemon.Types {
		fmt.Println("  -", pokemon.Types[t].Type.Name)
	}
	return nil
}
