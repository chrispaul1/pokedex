package main

import (
	"fmt"

	"github.com/chrispaul1/pokedexcli/internal/pokecache"
)

func commandInspect(c *Config, exploreCache *pokecache.Cache, areaName string, pokemonName string) error {

	pokemon, caught := c.pokedex[pokemonName]
	if len(pokemonName) == 0 {
		fmt.Println("Pokemon name not given!")
		return nil
	}
	if !caught {
		fmt.Printf("%s has not been caught!\n", pokemonName)
		return nil
	}

	fmt.Printf("Name: %s", pokemon.Name)
	fmt.Printf("\nHeight: %d", pokemon.Height)
	fmt.Printf("\nWeight: %d", pokemon.Weight)
	fmt.Println("\nStats:")
	for _, i := range pokemon.PokemonStats {
		fmt.Printf(" -%s: %d\n", i.StatName.Name, i.Base_Stat)
	}
	fmt.Println("Types:")
	for _, i := range pokemon.PokemonType {
		fmt.Printf(" - %s\n", i.Type.Name)
	}

	return nil
}
