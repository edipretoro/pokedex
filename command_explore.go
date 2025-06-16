package main

import "fmt"

func commandExplore(cfg *config, name *string) error {
	if name == nil || *name == "" {
		return fmt.Errorf("please provide a location name to explore")
	}
	p, err := cfg.pokeapiClient.ListPokemon(name)
	if err != nil {
		return err
	}
	fmt.Printf("Exploring %s...\n", *name)
	for _, encounter := range p.PokemonEncounters {
		fmt.Println("-", encounter.Pokemon.Name)
	}
	return nil
}

