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
	Name            string `json:"name"`
	Base_Experience int    `json:"base_experience"`
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

	if catchRate < 0.1 {
		catchRate = 0.1
	} else if catchRate > 0.8 {
		catchRate = 0.8
	}
	fmt.Printf("\nThrowing a Pokeball at %s...", pokemonName)

	if randomNum <= catchRate {
		fmt.Printf("\n%s was caught!\n", pokemonName)
		c.pokedex[pokemonName] = newPokemon
	} else {
		fmt.Printf("\n%s escaped!\n", pokemonName)
	}

	return nil
}
