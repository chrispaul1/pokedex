package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"math/rand"
	"net/http"

	"github.com/chrispaul1/pokedexcli/internal/pokecache"
)

type Pokemon struct {
	Name            string  `json:"name"`
	Base_Experience int     `json:"base_experience"`
	Height          int     `json:"height"`
	Weight          int     `json:"weight"`
	PokemonStats    []Stats `json:"stats"`
	PokemonType     []Types `json:"types"`
}

type Stats struct {
	Base_Stat int  `json:"base_stat"`
	StatName  Stat `json:"stat"`
}

type Stat struct {
	Name string `json:"name"`
}

type Types struct {
	Type PokemonType `json:"type"`
}

type PokemonType struct {
	Name string `json:"name"`
}

func commandCatch(c *Config, exploreCache *pokecache.Cache, areaName string, pokemonName string) error {
	url := fmt.Sprintf("https://pokeapi.co/api/v2/pokemon/%s/", pokemonName)
	res, err := http.Get(url)
	if err != nil {
		fmt.Println(err)
		log.Fatal(err)
	}
	defer res.Body.Close()

	data, err := io.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err)
		log.Fatal(err)
	}

	var newPokemon Pokemon
	err = json.Unmarshal(data, &newPokemon)
	if err != nil {
		fmt.Println(err)
		log.Fatal(err)
	}

	randomNum := rand.Float64()

	catchRate := (float64)(50-((newPokemon.Base_Experience-100)/2)) / 100

	if catchRate < 0.3 {
		catchRate = 0.3
	} else if catchRate > 0.7 {
		catchRate = 0.7
	}
	fmt.Printf("\nThrowing a Pokeball at %s...", pokemonName)

	if randomNum <= catchRate {
		fmt.Printf("\n%s was caught!\n", pokemonName)
		fmt.Println("You may now inspect it with the inspect command")
		c.pokedex[pokemonName] = newPokemon
	} else {
		fmt.Printf("\n%s escaped!\n", pokemonName)
	}

	return nil
}
