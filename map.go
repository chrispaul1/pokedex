package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"

	"github.com/chrispaul1/pokedexcli/internal/pokecache"
)

type Config struct {
	locationAreaSegment LocationAreaSegment
	Offset              int
	pokedex             map[string]Pokemon
}

type LocationAreaSegment struct {
	Count    int            `json:"count"`
	Next     *string        `json:"next"`
	Previous *string        `json:"previous"`
	Results  []LocationArea `json:"results"`
}

type LocationArea struct {
	Name string `json:"name"`
	Url  string `json:"url"`
}

func commandMap(c *Config, exploreCache *pokecache.Cache, areaName string, pokemonName string) error {
	url := fmt.Sprintf("https://pokeapi.co/api/v2/location-area/?limit=20&offset=%d", c.Offset)

	cachedData, found := pokecache.GetFromCache(url)
	if found {
		var cachedSegment LocationAreaSegment
		err := json.Unmarshal(cachedData, &cachedSegment)
		if err != nil {
			log.Fatal(err)
		}
		c.locationAreaSegment = cachedSegment
		c.Offset += 20
		for _, i := range cachedSegment.Results {
			fmt.Println(i.Name)
		}
		return nil
	}

	res, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
	}
	defer res.Body.Close()

	data, err := io.ReadAll(res.Body)
	if err != nil {
		log.Fatal(err)
	}

	var newSegment LocationAreaSegment
	err = json.Unmarshal(data, &newSegment)
	if err != nil {
		log.Fatal(err)
	}
	c.locationAreaSegment = newSegment
	c.Offset += 20
	for _, i := range newSegment.Results {
		fmt.Println(i.Name)
	}
	pokecache.AddToCache(url, data)
	return nil
}

func commandMapB(c *Config, exploreCache *pokecache.Cache, areaName string, pokemonName string) error {
	if c.Offset == 0 {
		fmt.Println("you're on the first page")
		return nil
	}
	c.Offset -= 20
	if c.Offset < 0 {
		c.Offset = 0
	}

	url := fmt.Sprintf("https://pokeapi.co/api/v2/location-area/?limit=20&offset=%d", c.Offset)
	res, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
	}
	data, err := io.ReadAll(res.Body)
	if err != nil {
		log.Fatal(err)
	}

	var newSegemnt LocationAreaSegment
	err = json.Unmarshal(data, &newSegemnt)
	if err != nil {
		log.Fatal(err)
	}
	c.locationAreaSegment = newSegemnt
	for _, i := range newSegemnt.Results {
		fmt.Println(i.Name)
	}
	return nil
}
