package pokeapi

import (
	"fmt"
	"io"
	"net/http"

	"github.com/jcfullmer/pokedexcli/Types"
)

func ReqToJsonStruct(url string, c *Types.Config) ([]byte, error) {
	//Check Cache for result
	val, ok := c.Cache.Get(url)
	if ok {
		return val, nil
	}
	res, err := http.Get(url)
	if err != nil {
		return []byte{}, err
	}
	body, err := io.ReadAll(res.Body)
	defer res.Body.Close()
	if res.StatusCode > 299 {
		return []byte{}, fmt.Errorf("invalid status code")
	}
	if err != nil {
		return []byte{}, err
	}
	c.Cache.Add(url, body)
	return body, nil
}
