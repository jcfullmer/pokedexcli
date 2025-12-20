package pokeapi

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/jcfullmer/pokedexcli/Types"
)

func ReqToJsonStruct(url string, c *Types.Config) (Types.JsonPokeAPI, error) {
	jsonRes := Types.JsonPokeAPI{}
	//Check Cache for result
	val, ok := c.Cache.Get(url)
	if ok {
		err := json.Unmarshal(val, &jsonRes)
		if err != nil {
			return jsonRes, err
		}
		return jsonRes, nil
	}
	res, err := http.Get(url)
	if err != nil {
		return jsonRes, err
	}
	body, err := io.ReadAll(res.Body)
	defer res.Body.Close()
	if res.StatusCode > 299 {
		return jsonRes, fmt.Errorf("invalid status code")
	}
	if err != nil {
		return jsonRes, err
	}

	err = json.Unmarshal(body, &jsonRes)
	if err != nil {
		return jsonRes, err
	}
	c.Cache.Add(url, body)
	return jsonRes, nil
}
