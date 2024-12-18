package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/chrispaul1/pokedexcli/internal/pokecache"
)

type cliCommand struct {
	name        string
	description string
	callback    func(*Config, *pokecache.Cache, string, string) error
}

var commands map[string]cliCommand

func startREPL() {
	pokecache.InitCache(5 * time.Second)
	exploreCache := pokecache.NewCache(5 * time.Second)
	commands = map[string]cliCommand{
		"help": {
			name:        "help",
			description: "Displays a help message",
			callback:    commandHelp,
		},
		"map": {
			name:        "map",
			description: "Displays the next 20 locations",
			callback:    commandMap,
		},
		"mapb": {
			name:        "mapb",
			description: "Displays the previous 20 locations",
			callback:    commandMapB,
		},
		"explore": {
			name:        "explore",
			description: "Displays the pokemon's in the area",
			callback:    commandExplore,
		},
		"catch": {
			name:        "catch",
			description: "Attempts to capture a pokemon",
			callback:    commandCatch,
		},
		"inspect": {
			name:        "inspect",
			description: "Lists the pokemon's detail",
			callback:    commandInspect,
		},
		"pokedex": {
			name:        "pokedex",
			description: "Lists all the pokemons that have been caught",
			callback:    commandPokedex,
		},
		"exit": {
			name:        "exit",
			description: "Exit the Pokedex",
			callback:    commandExit,
		},
	}

	config := &Config{
		pokedex: make(map[string]Pokemon),
	}
	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print("Pokedex > ")
		scanner.Scan()
		userText := scanner.Text()
		userText = strings.ToLower(userText)
		splitText := strings.Fields(userText)
		if len(splitText) > 0 {
			firstWord := splitText[0]
			areaName := ""
			pokemonName := ""
			command, ok := commands[firstWord]
			if ok {
				if len(splitText) > 1 {
					areaName = splitText[1]
					pokemonName = splitText[1]
				}
				if err := command.callback(config, exploreCache, areaName, pokemonName); err != nil {
					fmt.Println(err)
				}
			} else {
				fmt.Print("\nUnknown command\n")
			}
		}
	}
}

func cleanInput(text string) []string {
	text = strings.TrimSpace(text)
	text = strings.ToLower(text)
	str := strings.Fields(text)

	return str
}

func commandExit(c *Config, exploreCache *pokecache.Cache, areaName string, pokemonName string) error {
	fmt.Print("Closing the Pokedex... Goodbye!\n")
	os.Exit(0)
	return nil
}

func commandHelp(c *Config, exploreCache *pokecache.Cache, areaName string, pokemonName string) error {
	keys := []string{"help", "exit"}
	fmt.Println("\nWelcome to the Pokedex!\nUsage:")
	fmt.Println()
	for _, i := range keys {
		fmt.Printf("%s: %s\n", commands[i].name, commands[i].description)
	}
	return nil
}

func commandPokedex(c *Config, exploreCache *pokecache.Cache, areaName string, pokemonName string) error {
	fmt.Println("Your Pokedex:")
	for _, i := range c.pokedex {
		fmt.Printf(" - %s\n", i.Name)
	}
	return nil
}
