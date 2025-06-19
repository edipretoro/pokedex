package main

import (
	"fmt"
	"math/rand"
)

func commandCatch(cfg *config, name *string) error {
	const maxExperience int = 300
	if name == nil || *name == "" {
		return fmt.Errorf("please provide a Pokemon name and try and catch it")
	}
	pokemon, err := cfg.pokeapiClient.Pokemon(name)
	if err != nil {
		return err
	}
	chance := maxExperience - pokemon.BaseExperience
	_ = rand.Intn(chance)
	
	fmt.Printf("Throwing a Pokeball at %s...\n", *name)
	v := rand.Intn(maxExperience)
	if v < chance {
		fmt.Printf("%s was caught!\n", *name)
		cfg.pokedex[*name] = pokemon
	} else {
		fmt.Printf("%s escaped!\n", *name)
	}
	return nil
}
