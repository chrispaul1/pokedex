package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type cliCommand struct {
	name        string
	description string
	callback    func() error
}

var commands map[string]cliCommand

func startREPL() {
	commands = map[string]cliCommand{
		"help": {
			name:        "help",
			description: "Displays a help message",
			callback:    commandHelp,
		},
		"exit": {
			name:        "exit",
			description: "Exit the Pokedex",
			callback:    commandExit,
		},
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
			command, ok := commands[firstWord]
			if ok {
				if err := command.callback(); err != nil {
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

func commandExit() error {
	fmt.Print("Closing the Pokedex... Goodbye!\n")
	os.Exit(0)
	return nil
}

func commandHelp() error {
	keys := []string{"help", "exit"}
	fmt.Println("\nWelcome to the Pokedex!\nUsage:")
	fmt.Println()
	for _, i := range keys {
		fmt.Printf("%s: %s\n", commands[i].name, commands[i].description)
	}
	return nil
}
