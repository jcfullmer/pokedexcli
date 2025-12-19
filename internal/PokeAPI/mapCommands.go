package pokeapi

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/jcfullmer/pokedexcli/Types"
)

const PokeApiLocationArea = "https://pokeapi.co/api/v2/location-area/"

func CommandMap(conf *Types.Config) error {
	res, err := http.Get(PokeApiLocationArea)
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
	err = json.Unmarshal(body, &conf)
	if err != nil {
		return nil
	}
	fmt.Println(conf.Next)

	return nil
}
