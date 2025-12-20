package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/jcfullmer/pokedexcli/Types"
	pokeapi "github.com/jcfullmer/pokedexcli/internal/PokeAPI"
)

type cliCommand struct {
	name        string
	description string
	callback    func(*Types.Config) error
}

func GetCommands() map[string]cliCommand {
	return map[string]cliCommand{
		"exit": {
			name:        "exit",
			description: "Exit the Pokedex",
			callback:    commandExit,
		},
		"help": {
			name:        "help",
			description: "Display a help message",
			callback:    commandHelp,
		},
		"map": {
			name:        "map",
			description: "Get Location Areas",
			callback:    pokeapi.CommandMap,
		},
		"mapb": {
			name:        "mapb",
			description: "Get Previous Location Areas",
			callback:    pokeapi.CommandMapB,
		},
	}
}

func cleanInput(text string) []string {
	output := strings.ToLower(text)
	words := strings.Fields(output)
	return words
}

func PokedexInput() {
	input := bufio.NewScanner(os.Stdin)
	Conf := &Types.Config{
		Next:     Types.PokeApiLocationArea,
		Previous: "",
	}
	for {
		fmt.Print("Pokedex > ")
		input.Scan()
		words := cleanInput(input.Text())
		if len(words) == 0 {
			continue
		}
		command := words[0]
		val, ok := GetCommands()[command]
		if ok {
			err := val.callback(Conf)
			if err != nil {
				fmt.Println(err)
			}
		} else {
			fmt.Println("Unknown command")
		}
	}
}

func commandExit(conf *Types.Config) error {
	fmt.Println("Closing the Pokedex... Goodbye!")
	os.Exit(0)
	return nil
}

func commandHelp(conf *Types.Config) error {
	fmt.Println("Welcome to the Pokedex!\nUsage:")
	for _, command := range GetCommands() {
		fmt.Printf("%v: %v\n", command.name, command.description)
	}
	return nil
}
