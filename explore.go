package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"

	"github.com/chrispaul1/pokedexcli/internal/pokecache"
)

type AreaName struct {
	Name             string             `json:"name"`
	PokemonEncounter []PokemonEncounter `json:"pokemon_encounters"`
}

type PokemonEncounter struct {
	Pokemons PokemonName `json:"pokemon"`
}

type PokemonName struct {
	Name string `json:"name"`
}

func commandExplore(c *Config, exploreCache *pokecache.Cache, areaName string, pokemonName string) error {

	if len(areaName) == 0 {
		fmt.Println("Area name not given!")
		return nil
	}
	url := fmt.Sprintf("https://pokeapi.co/api/v2/location-area/%s", areaName)
	cachedData, found := exploreCache.Get(areaName)
	if found {
		var cachedArea AreaName
		err := json.Unmarshal(cachedData, &cachedArea)
		if err != nil {
			fmt.Println(err)
		}
		fmt.Printf("\nExploring %s...\nFound Pokemond:\n", areaName)

		for _, i := range cachedArea.PokemonEncounter {
			fmt.Println(" - ", i.Pokemons.Name)
		}
		return nil
	}

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

	var newArea AreaName
	err = json.Unmarshal(data, &newArea)
	if err != nil {
		fmt.Println(err)
		log.Fatal(err)
	}
	fmt.Println("\nFound Pokemon:")
	for _, i := range newArea.PokemonEncounter {
		fmt.Println(" - ", i.Pokemons.Name)
	}

	exploreCache.Add(areaName, data)
	return nil
}
