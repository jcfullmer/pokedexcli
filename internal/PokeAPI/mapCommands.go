package pokeapi

import (
	"fmt"

	"github.com/jcfullmer/pokedexcli/Types"
)

func CommandMap(conf *Types.Config) error {
	LocationRes, err := ReqToJsonStruct(conf.Next, conf)
	if err != nil {
		return err
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
	LocationRes, err := ReqToJsonStruct(conf.Previous, conf)
	if err != nil {
		return err
	}
	conf.Next = LocationRes.Next
	conf.Previous = LocationRes.Previous
	for _, location := range LocationRes.Results {
		fmt.Println(location.Name)
	}
	return nil
}
