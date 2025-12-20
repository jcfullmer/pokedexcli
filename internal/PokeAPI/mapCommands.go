package pokeapi

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/jcfullmer/pokedexcli/Types"
)

func CommandMap(conf *Types.Config) error {
	res, err := http.Get(conf.Next)
	if err != nil {
		return err
	}
	body, err := io.ReadAll(res.Body)
	defer res.Body.Close()
	if res.StatusCode > 299 {
		return fmt.Errorf("invalid status code")
	}
	if err != nil {
		return err
	}
	LocationRes := Types.JsonPokeAPI{}
	err = json.Unmarshal(body, &LocationRes)
	if err != nil {
		return nil
	}
	conf.Next = LocationRes.Next
	conf.Previous = LocationRes.Previous
	for _, location := range LocationRes.Results {
		fmt.Println(location.Name)
	}
	return nil
}

func CommandMapB(conf *Types.Config) error {
	if conf.Previous == "" {
		return fmt.Errorf("previous link is empty")
	}
	res, err := http.Get(conf.Previous)
	if err != nil {
		return err
	}
	body, err := io.ReadAll(res.Body)
	defer res.Body.Close()
	if res.StatusCode > 299 {
		return fmt.Errorf("invalid status code")
	}
	if err != nil {
		return err
	}
	LocationRes := Types.JsonPokeAPI{}
	err = json.Unmarshal(body, &LocationRes)
	if err != nil {
		return nil
	}
	conf.Next = LocationRes.Next
	conf.Previous = LocationRes.Previous
	for _, location := range LocationRes.Results {
		fmt.Println(location.Name)
	}
	return nil
}
