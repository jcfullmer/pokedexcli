package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func cleanInput(text string) []string {
	output := strings.ToLower(text)
	words := strings.Fields(output)
	return words
}

func PokedexInput() {
	for true {
		fmt.Print("Pokedex > ")
		input := bufio.NewScanner(os.Stdin)
		input.Scan()
		words := cleanInput(input.Text())
		command := words[0]
		fmt.Printf("Your command was: %s\n", command)
	}
}
