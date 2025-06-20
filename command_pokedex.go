package main

import (
	"fmt"
)

func commandPokedex(cfg *config, name *string) error {
	if len(cfg.pokedex) > 0 {
		fmt.Println("Your Pokedex:")
	} else {
		fmt.Println("Your Pokedex is empty.")
		return nil
	}
	for pokemon := range cfg.pokedex {
		fmt.Println("  -", pokemon)
	}
	return nil
}
