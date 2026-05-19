package pokeapi

import (
	"encoding/json"
	"io"
	"net/http"
)

func (c *Client) PropertiesRetrieval(pokemon string) (PokemonTraitList, error) {
	var p PokemonTraitList

	url := baseURL + "/pokemon/" + pokemon + "/"

	// cache logic
	if val, ok := c.userCache.Get(url); ok {
		if err := json.Unmarshal(val, &p); err != nil {
			return PokemonTraitList{}, err
		}
		return p, nil
	}

	// build the request
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return PokemonTraitList{}, err
	}

	// execute it
	res, err := c.cli.Do(req)
	if err != nil {
		return PokemonTraitList{}, err
	}
	defer res.Body.Close()

	// read response body
	body, err := io.ReadAll(res.Body)
	if err != nil {
		return PokemonTraitList{}, err
	}

	// store in cache
	c.userCache.Add(url, body)

	// unmarshal json into go structs
	err = json.Unmarshal(body, &p)
	if err != nil {
		return PokemonTraitList{}, err
	}

	return p, nil
}
