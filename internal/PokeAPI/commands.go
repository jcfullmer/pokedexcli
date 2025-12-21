package pokeapi

import (
	"encoding/json"
	"fmt"
	"math/rand"

	"github.com/jcfullmer/pokedexcli/Types"
)

func CommandMap(conf *Types.Config, _ string) error {
	resBytes, err := ReqToJsonStruct(conf.Next, conf)
	if err != nil {
		return err
	}
	LocationRes := Types.JsonPokeAPI{}
	if err := json.Unmarshal(resBytes, &LocationRes); err != nil {
		return err
	}
	conf.Next = LocationRes.Next
	conf.Previous = LocationRes.Previous
	for _, location := range LocationRes.Results {
		fmt.Println(location.Name)
	}
	return nil
}

func CommandMapB(conf *Types.Config, _ string) error {
	if conf.Previous == "" {
		return fmt.Errorf("previous link is empty")
	}
	resBytes, err := ReqToJsonStruct(conf.Previous, conf)
	if err != nil {
		return err
	}
	LocationRes := Types.JsonPokeAPI{}
	if err := json.Unmarshal(resBytes, &LocationRes); err != nil {
		return err
	}
	conf.Next = LocationRes.Next
	conf.Previous = LocationRes.Previous
	for _, location := range LocationRes.Results {
		fmt.Println(location.Name)
	}
	return nil
}

func CommandExplore(conf *Types.Config, location string) error {
	fmt.Println("Exploring", location)
	areaURL := Types.PokeApiLocationArea + "location-area/" + location
	resBytes, err := ReqToJsonStruct(areaURL, conf)
	if err != nil {
		return err
	}
	Location := Types.Locations{}
	if err != json.Unmarshal(resBytes, &Location) {
		return err
	}
	for _, encounter := range Location.PokemonEncounters {
		fmt.Println("- " + encounter.Pokemon.Name)

	}
	return nil
}

func CommandCatch(conf *Types.Config, pokemon string) error {
	pokemonURL := Types.PokeApiLocationArea + "pokemon/" + pokemon
	fmt.Println(pokemonURL)
	resBytes, err := ReqToJsonStruct(pokemonURL, conf)
	if err != nil {
		return err
	}
	pokemonStruct := Types.Pokemon{}
	if err != json.Unmarshal(resBytes, &pokemonStruct) {
		return err
	}
	fmt.Printf("Throwing a Pokeball at %v...\n", pokemonStruct.Name)
	roll := rand.Intn(100)
	if roll > pokemonStruct.BaseExperience/3 {
		if _, ok := conf.CaughtPokemon[pokemonStruct.Name]; ok {
			return fmt.Errorf("you caught the pokemon, but you already had one so you had to release it.")
		} else {
			fmt.Printf("%v was caught!\n", pokemonStruct.Name)
			conf.CaughtPokemon[pokemonStruct.Name] = pokemonStruct
		}
	} else {
		fmt.Printf("%v escaped!\n", pokemonStruct.Name)
	}
	return nil
}

func CommandInspect(conf *Types.Config, pokemon string) error {
	poke, ok := conf.CaughtPokemon[pokemon]
	if !ok {
		return fmt.Errorf("you have not caught that pokemon")
	}
	fmt.Printf("Name: %v\n", poke.Name)
	fmt.Printf("Height: %v\n", poke.Height)
	fmt.Printf("Weight: %v\n", poke.Weight)
	fmt.Println("Stats:")
	for _, stat := range poke.Stats {
		fmt.Printf("	-%v: %v\n", stat.Stat.Name, stat.BaseStat)
	}
	fmt.Println("Types:")
	for _, t := range poke.Types {
		fmt.Printf("	-%v\n", t.Type.Name)
	}
	return nil
}
